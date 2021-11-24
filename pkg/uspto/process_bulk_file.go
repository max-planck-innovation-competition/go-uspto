package uspto

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var regexFile = regexp.MustCompile(`file="([A-Z0-9-.]+)"`)

func ProcessBulkFile(sourceFile, destinationFolder string) (err error) {
	logger := log.WithField("zipFile", sourceFile)
	logger.Info("reading file")
	// read zip file
	readCloser, err := zip.OpenReader(sourceFile)
	if err != nil {
		msg := "Failed to open: %s"
		logger.Fatalf(msg, err)
	}
	logger.Info("file read")
	// close file after read
	defer func() {
		errClose := readCloser.Close()
		if errClose != nil {
			logger.Fatalf("Failed to close file: %s", errClose)
		}
	}()
	logger.Info("iterate over files")
	// iterate over all files in the zip directory
	for _, file := range readCloser.File {
		logger.WithField("filename", file.Name).Info("found")
		extension := filepath.Ext(file.Name)
		if extension == ".xml" {
			if errFile := processZippedFiles(file, destinationFolder); errFile != nil {
				return
			}
		} else {
			logger.WithField("filename", file.Name).Info("skipping file")
		}
	}
	logger.Info("successfully done")
	return
}

func processZippedFiles(file *zip.File, destinationFolder string) (err error) {
	logger := log.WithField("filename", file.Name).WithField("routine", "main")
	logger.Info("process file")
	ctx := context.TODO()
	fc, err := file.Open()
	if err != nil {
		msg := "failed to open zip %s for reading: %s"
		err = fmt.Errorf(msg, file.Name, err)
		logger.Error(err)
		return
	}
	defer func() {
		errClose := fc.Close()
		if errClose != nil {
			ctx.Done()
			logger.Fatalf("Failed to close file: %s", errClose)
		}
	}()
	// init channels and sync
	var wg sync.WaitGroup
	chContent := make(chan string)
	chFilename := make(chan string)
	chFileEnd := make(chan bool)
	// start 2nd process
	go fileWriter(ctx, destinationFolder, &wg, chContent, chFilename, chFileEnd)
	// scan file
	scanner := bufio.NewScanner(fc)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	counter := 1
	for scanner.Scan() {
		line := scanner.Text()
		switch strings.TrimSpace(line) {
		case `</us-patent-grant>`, `</PATDOC>`:
			// identify the end of the document
			logger.Trace("us-patent-grant xml end")
			logger.WithField("count", counter).Info("done with file")
			// end of the file
			wg.Add(1)
			chContent <- line
			wg.Add(1)
			chFileEnd <- true
			counter++
			break
		default:
			// extract filename from xml file
			if strings.Contains(line, `<us-patent-grant `) {
				// version 4.0
				logger.Trace("us-patent-grant filename")
				// if there is a line which contains the beginning of the document xml tag
				// extract the filename
				res := regexFile.FindAllStringSubmatch(line, -1)
				if len(res) != 1 && len(res[0]) != 1 {
					msg := "failed extract filename"
					err = fmt.Errorf(msg)
					logger.Error(err)
					return
				}
				filename := res[0][1]
				logger.WithField("xmlFile", filename).Info("start")
				// send the filename
				wg.Add(1)
				chFilename <- filename
			} else if strings.Contains(line, `<B110>`) {
				// version 2.5
				logger.Trace("b110 number")
				xmlElement, _ := goquery.NewDocumentFromReader(bytes.NewReader([]byte(line)))
				filename := strings.TrimSpace(xmlElement.Find("B110").Text())
				if len(filename) > 0 {
					logger.WithField("xmlFile", filename).Info("start")
					// append ending
					filename += ".XML"
					// send the filename
					wg.Add(1)
					chFilename <- filename
				} else {
					msg := "failed extract filename"
					err = fmt.Errorf(msg)
					logger.Error(err)
				}
			}
			// send the content
			wg.Add(1)
			chContent <- line
		}
		logger.Trace("wait")
		wg.Wait()
	}
	// scanner error
	if err = scanner.Err(); err != nil {
		msg := "failed to scan file %s line by line: %s"
		err = fmt.Errorf(msg, file.Name, err)
		logger.Error(err)
		ctx.Done()
		return
	}
	return
}

func fileWriter(
	ctx context.Context,
	destinationFolder string,
	wg *sync.WaitGroup,
	chContent <-chan string,
	chFilename <-chan string,
	chFileEnd <-chan bool,
) {
	logger := log.WithField("routine", "writer")
	logger.Trace("started")
	filename := ""
	var buf strings.Builder

	for {
		select {
		case <-ctx.Done():
			logger.Info("received context done")
			return
		case end := <-chFileEnd:
			if end {
				// if the last string was transmitted
				logger.Trace("last stuff was transmitted")
				// create a new file based on the filename
				if len(filename) == 0 {
					msg := "failed to extract filename: %s"
					logger.Fatalf(msg, filename)
					return
				}
				file, err := os.Create(destinationFolder + "/" + filename)
				if err != nil {
					logger.Error(err)
					return
				}
				// write the data to the file
				_, errWrite := file.WriteString(buf.String())
				if errWrite != nil {
					ctx.Done()
					msg := "failed to write to buffer: %s"
					logger.Fatalf(msg, errWrite)
					return
				}
				// close the file
				errClose := file.Close()
				if errClose != nil {
					ctx.Done()
					msg := "failed to write close file: %s"
					logger.Fatalf(msg, errClose)
					return
				}
				// clear the string builder
				buf.Reset()
				// clear the filename
				filename = ""
				break
			}
		case content := <-chContent:
			logger.
				// WithField("content", content).
				Trace("received data")
			// skip the empty line if there is nothing in the buffer
			if buf.Len() == 0 && len(content) == 0 {
				break
			}
			// if there is content write it to the buffer
			_, errWrite := buf.WriteString(content + "\n")
			if errWrite != nil {
				ctx.Done()
				msg := "failed to write to buffer: %s"
				logger.Fatalf(msg, errWrite)
				return
			}
			break
		case filename = <-chFilename:
			logger.Debug("Set filename", filename)
			break
		}
		log.Trace("done")
		wg.Done()
	}

}
