package uspto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloadBulkFile(t *testing.T) {
	ass := assert.New(t)
	filePath, err := DownloadBulkFile("https://bulkdata.uspto.gov/data/patent/grant/redbook/fulltext/2021/ipg210907.zip", "./test-data")
	ass.NoError(err)
	ass.NotEmpty(filePath)
}

func TestDownloadGrantBulkFile(t *testing.T) {
	ass := assert.New(t)
	filePath, err := DownloadBulkFile("https://bulkdata.uspto.gov/data/patent/application/redbook/fulltext/2021/ipa210902.zip", "./test-data")
	ass.NoError(err)
	ass.NotEmpty(filePath)
}
