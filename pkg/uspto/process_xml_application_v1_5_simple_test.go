package uspto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessApplicationXML15Simple(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/1-5-a1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20010315", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("A1", patDoc.Kind)
	ass.Equal("20010000001", patDoc.DocNumber)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Solvent mixture for use in a vapor degreaser and method of cleaning an article in a vapor degreaser utilizing said solvent", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)
	ass.NotEmpty(patDoc.Abstract)
	return
}
