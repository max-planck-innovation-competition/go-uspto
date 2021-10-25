package uspto

import (
	"archive/zip"
	"bufio"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strings"
	"sync"
)

var regexFile = regexp.MustCompile(`file="([A-Z0-9-.]+)"`)

func ReadPatentXMLZip(sourceFile, destinationFolder string) (err error) {
	log.Info("reading file", sourceFile)
	// read zip file
	readCloser, err := zip.OpenReader(sourceFile)
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	log.Info("file read")
	// close file after read
	defer func() {
		errClose := readCloser.Close()
		if errClose != nil {
			log.Fatalf("Failed to close file: %s", errClose)
		}
	}()
	log.Info("iterate over files")
	// iterate over all files in the zip directory
	for _, file := range readCloser.File {
		log.WithField("filename", file.Name).Info("found")
		if errFile := processZippedFiles(file, destinationFolder); errFile != nil {
			return
		}
	}
	log.Info("done")
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
	var chContent chan string
	var chFilename chan string
	var chFileEnd chan bool
	// start 2nd process
	go fileWriter(ctx, destinationFolder, &wg, chContent, chFilename, chFileEnd)
	// scan file
	scanner := bufio.NewScanner(fc)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case `<?xml version="1.0" encoding="UTF-8"?>`:
			logger.Trace("xml header")
			wg.Add(1)
			chContent <- line
			// new file
			break
		case `</us-patent-grant>`:
			logger.Trace("us-patent-grant xml end")
			// end of the file
			wg.Add(1)
			chContent <- line
			wg.Add(1)
			chFileEnd <- true
			break
		default:
			if strings.Contains(line, `<us-patent-grant `) {
				logger.Trace("us-patent-grant title")
				res := regexFile.FindAllStringSubmatch(line, -1)
				if len(res) != 1 && len(res[0]) != 1 {
					msg := "failed extract filename"
					err = fmt.Errorf(msg)
					logger.Error(err)
					return
				}
				filename := res[0][0]
				// send the filename
				wg.Add(1)
				chFilename <- filename
				// send the content
				wg.Add(1)
				chContent <- line
			}
			chContent <- line
		}
		logger.Trace("wait")
		wg.Wait()
	}
	// scanner error
	if err = scanner.Err(); err != nil {
		ctx.Done()
		msg := "failed to scan file %s line by line: %s"
		err = fmt.Errorf(msg, file.Name, err)
		logger.Error(err)
		return
	}
	// close everything
	ctx.Done()
	close(chFilename)
	close(chContent)
	close(chFileEnd)
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
		case <-chFileEnd:
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
		case content := <-chContent:
			logger.Trace("content", content)
			_, errWrite := buf.WriteString(content)
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
		wg.Done()
	}

}
