package uspto

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var ErrNoFilenameInUrl = errors.New("no filename in url")
var ErrFileIsNoZip = errors.New("file is not zip")

func DownloadBulkFile(url string, exportDirectory string) (filePath string, err error) {
	logger := log.WithFields(log.Fields{"url": url, "exportDirectory": exportDirectory})
	// get the filename from the url
	urlElements := strings.Split(url, "/")
	if len(urlElements) == 0 {
		err = ErrNoFilenameInUrl
		logger.Error(err)
		return
	}
	// check if the file is a zip file
	fileName := urlElements[len(urlElements)-1]
	if !strings.Contains(fileName, ".zip") {
		err = ErrFileIsNoZip
		logger.Error(err)
		return
	}
	// create the output file
	filePath = filepath.Join(exportDirectory, fileName)
	out, err := os.Create(filePath)
	if err != nil {
		logger.Error(err)
		return
	}
	defer out.Close()

	// init the http client
	httpClient := NewHttpClient()
	// get the data
	log.WithField("url", url).Info("start download")
	resp, err := httpClient.Get(url)
	if err != nil {
		logger.Error(err)
		return
	}
	defer resp.Body.Close()

	// check server response
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", resp.Status)
		logger.Error(err)
		return
	}

	// write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}
	logger.Info("download complete")
	logger.WithField("path", filePath).Info("file saved")
	return
}
