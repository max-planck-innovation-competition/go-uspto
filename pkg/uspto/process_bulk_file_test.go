package uspto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadPatentXMLZip(t *testing.T) {
	//skipTest(t)
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ProcessBulkFile(
		"./test-data/test.zip", "./test-data/xml")
	ass.NoError(err)
}

func TestReadPatentXMLZip25(t *testing.T) {
	skipTest(t)
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	// 2.5
	err := ProcessBulkFile("./test-data/pg020101.zip", "./test-data/pg020101/xml")
	ass.NoError(err)
}

// version 4-0
func TestReadPatentApplicationsXMLZip(t *testing.T) {
	skipTest(t)
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ProcessBulkFile("./test-data/ipa210902.zip", "./test-data/xml-application")
	ass.NoError(err)
}

// version 1-5
func TestReadPatentApplicationsXMLZip15(t *testing.T) {
	skipTest(t)
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ProcessBulkFile("./test-data/application/grant-bulk/1-5-pab20010315_wk11.zip", "./test-data/xml-application")
	ass.NoError(err)
}

// missing grant
func TestReadPatentGrantXMLZip2022(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ProcessBulkFile("./test-data/grant-bulk/ipg210803.zip", "./test-data/grant-bulk")
	ass.NoError(err)
}
