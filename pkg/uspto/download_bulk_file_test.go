package uspto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloadBulkFile(t *testing.T) {
	ass := assert.New(t)
	err := DownloadBulkFile("https://bulkdata.uspto.gov/data/patent/grant/redbook/fulltext/2021/ipg210907.zip", "./test-data")
	ass.NoError(err)
}
