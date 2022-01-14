package uspto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadPatentXMLZip(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ProcessBulkFile("./test-data/ipg211019.zip", "./test-data/xml")
	ass.NoError(err)
}

func TestReadPatentApplicationsXMLZip(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ProcessBulkFile("./test-data/ipa210902.zip", "./test-data/xml-application")
	ass.NoError(err)
}

func TestReadPatentXMLZip25(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	// 2.5
	err := ProcessBulkFile("./test-data/pg020101.zip", "./test-data/pg020101/xml")
	ass.NoError(err)
}
