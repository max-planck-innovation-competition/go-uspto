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

func TestReadPatentXMLZip25(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	// 2.5
	err := ProcessBulkFile("./test-data/pg020101.zip", "./test-data/pg020101/xml")
	ass.NoError(err)
}

// version 4-0
func TestReadPatentApplicationsXMLZip(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ProcessBulkFile("./test-data/ipa210902.zip", "./test-data/xml-application")
	ass.NoError(err)
}

// version 1-5
func TestReadPatentApplicationsXMLZip15(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ProcessBulkFile("./test-data/application/bulk/1-5-pab20010315_wk11.zip", "./test-data/xml-application")
	ass.NoError(err)
}

// 2022
func TestReadPatentApplicationsXMLZip2022(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ProcessBulkFile("./test-data/ipa220203.zip", "./test-data/ipa220203")
	ass.NoError(err)
}
