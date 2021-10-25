package uspto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadPatentXMLZip(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	err := ReadPatentXMLZip("./test-data/ipg211019.zip", "./test-data/xml")
	ass.NoError(err)
}
