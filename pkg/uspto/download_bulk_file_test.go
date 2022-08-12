package uspto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloadBulkFile(t *testing.T) {
	skipTest(t)
	ass := assert.New(t)
	filePath, err := DownloadBulkFile("https://bulkdata.uspto.gov/data/patent/grant/redbook/fulltext/2021/ipg210907.zip", "./test-data")
	ass.NoError(err)
	ass.NotEmpty(filePath)
}

func TestDownloadApplicationBulkFile(t *testing.T) {
	skipTest(t)
	ass := assert.New(t)
	filePath, err := DownloadBulkFile("https://bulkdata.uspto.gov/data/patent/application/redbook/fulltext/2021/ipa210902.zip", "./test-data")
	ass.NoError(err)
	ass.NotEmpty(filePath)
}

func TestDownloadApplicationBulkFile2(t *testing.T) {
	skipTest(t)
	ass := assert.New(t)
	filePath, err := DownloadBulkFile("https://bulkdata.uspto.gov/data/patent/application/redbook/fulltext/2022/ipa220203.zip", "./test-data")
	ass.NoError(err)
	ass.NotEmpty(filePath)
}

func TestDownloadApplicationBulkFile3(t *testing.T) {
	skipTest(t)
	ass := assert.New(t)
	filePath, err := DownloadBulkFile("https://bulkdata.uspto.gov/data/patent/application/redbook/fulltext/2022/ipa220526.zip", "./test-data")
	ass.NoError(err)
	ass.NotEmpty(filePath)
}
