package uspto

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Product string

const (
	XmlPatents Product = "PTGRXML"
)

type BulkFileResponse struct {
	ProductLinkPath   string      `json:"productLinkPath"`
	ProductIdentifier int         `json:"productIdentifier"`
	ProductShortName  string      `json:"productShortName"`
	ProductDesc       string      `json:"productDesc"`
	ProductTitle      string      `json:"productTitle"`
	ProductFrequency  string      `json:"productFrequency"`
	ProductLevel      string      `json:"productLevel"`
	ProductFromDate   string      `json:"productFromDate"`
	ProductToDate     string      `json:"productToDate"`
	NumberOfFiles     int         `json:"numberOfFiles"`
	ParentProduct     interface{} `json:"parentProduct"`
	ProductFiles      []struct {
		FileLinkPath    string `json:"fileLinkPath"`
		FileIdentifier  int    `json:"fileIdentifier"`
		FileName        string `json:"fileName"`
		FileSize        int    `json:"fileSize"`
		FileDownloadURL string `json:"fileDownloadUrl"`
		FileFromTime    string `json:"fileFromTime"`
		FileToTime      string `json:"fileToTime"`
		FileType        string `json:"fileType"`
		FileReleaseDate string `json:"fileReleaseDate"`
	} `json:"productFiles"`
}

// GetPatentXmlBulkFileList returns the download links to the zipped archives between two dates
func GetPatentXmlBulkFileList(start time.Time, end time.Time) (downloadLinks []string, err error) {
	res, err := GetBulkFileList(XmlPatents, start, end)
	if err != nil {
		log.Error(err)
		return
	}
	// iterate over res
	for _, r := range res.ProductFiles {
		downloadLinks = append(downloadLinks, r.FileDownloadURL)
	}
	return
}

// GetBulkFileList returns a list of available files of a specific USPTO product
func GetBulkFileList(product Product, start time.Time, end time.Time) (res BulkFileResponse, err error) {

	// params
	m := map[string]interface{}{
		"name":     string(product),
		"fromDate": start.Format("2006-01"),
		"toDate":   end.Format("2006-01"),
	}
	jsonStr, _ := json.Marshal(m)
	params := url.Values{"data": {string(jsonStr[:])}}

	// init http client
	client := NewHttpClient()
	// build req
	// https://developer.uspto.gov/products/bdss/get/ajax?data=
	// {"name":"PTGRXML","fromDate":"2020-01","toDate":"2021-10"}

	reqUrl := ENDPOINT_HOST + "/products/bdss/get/ajax?" + params.Encode()
	log.Debug("GET: ", reqUrl)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Error(err)
		return
	}
	// add header
	req.Header.Add("user-agent", "raw")
	// make request
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("No 200 status code: " + strconv.Itoa(resp.StatusCode))
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return
	}
	// parse data
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		log.Error(err)
		return
	}
	// close body
	err = resp.Body.Close()
	if err != nil {
		log.Error(err)
		return
	}
	return
}
