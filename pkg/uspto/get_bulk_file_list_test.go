package uspto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetBulkFileList(t *testing.T) {
	ass := assert.New(t)
	loc, err := time.LoadLocation("Europe/Berlin")
	ass.NoError(err)
	start := time.Date(2021, 9, 1, 0, 0, 0, 0, loc)
	end := time.Date(2021, 10, 1, 0, 0, 0, 0, loc)
	res, err := GetBulkFileList(XmlPatentGrants, start, end)
	ass.NoError(err)
	ass.NotNil(res)
	fmt.Println(res)
	ass.NotEmpty(res.ProductFiles)
}

func TestGetPatentGrantXmlBulkFileList(t *testing.T) {
	ass := assert.New(t)
	loc, err := time.LoadLocation("Europe/Berlin")
	ass.NoError(err)
	start := time.Date(2021, 9, 1, 0, 0, 0, 0, loc)
	end := time.Date(2021, 10, 1, 0, 0, 0, 0, loc)
	res, err := GetPatentGrantXmlBulkFileList(start, end)
	ass.NoError(err)
	ass.NotNil(res)
	fmt.Println(res)
	ass.NotEmpty(res)
}

func TestGetPatentGrantXmlBulkFileListEarly(t *testing.T) {
	ass := assert.New(t)
	loc, err := time.LoadLocation("Europe/Berlin")
	ass.NoError(err)
	start := time.Date(2002, 9, 1, 0, 0, 0, 0, loc)
	end := time.Date(2002, 10, 1, 0, 0, 0, 0, loc)
	res, err := GetPatentGrantXmlBulkFileList(start, end)
	ass.NoError(err)
	ass.NotNil(res)
	fmt.Println(res)
	ass.NotEmpty(res)
}

func TestGetPatentApplicationXmlBulkFileList(t *testing.T) {
	ass := assert.New(t)
	loc, err := time.LoadLocation("Europe/Berlin")
	ass.NoError(err)
	start := time.Date(2021, 9, 1, 0, 0, 0, 0, loc)
	end := time.Date(2021, 10, 1, 0, 0, 0, 0, loc)
	res, err := GetPatentApplicationXmlBulkFileList(start, end)
	ass.NoError(err)
	ass.NotNil(res)
	fmt.Println(res)
	ass.NotEmpty(res)
}
