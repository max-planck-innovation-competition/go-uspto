package uspto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestApplicationProcessXMLSimpleVersion40A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/v4-0-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20050106", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("A1", patDoc.Kind)
	ass.Equal("v4.0 2004-12-02", patDoc.DtdVersion)
	ass.Equal("PARALLEL-RUN", patDoc.Status)
	ass.Equal("20050000001", patDoc.DocNumber)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Novelty jeans", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)
	ass.NotEmpty(patDoc.Abstract)

	// IPC
	ass.Equal("A41D001/06", patDoc.Classifications[0].Text)

	// NATIONAL
	ass.Equal("002227000", patDoc.Classifications[1].Text)

	// inventors
	ass.Equal("Tina", patDoc.Inventors[0].FirsName)
	ass.Equal("Goldkind", patDoc.Inventors[0].LastName)
	ass.Equal("St. James", patDoc.Inventors[0].City)
	ass.Equal("NY", patDoc.Inventors[0].State)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
}

func TestApplicationProcessXMLSimpleVersion40A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/v4-0-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20050106", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("A2", patDoc.Kind)
	ass.Equal("v4.0 2004-12-02", patDoc.DtdVersion)
	ass.Equal("PARALLEL-RUN", patDoc.Status)
	ass.Equal("US20050004655A2-20050106.XML", patDoc.File)
	ass.Equal("20050004655", patDoc.DocNumber)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal("METHODS AND APPARATUS FOR A STENT HAVING AN EXPANDABLE WEB STRUCTURE", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)
	ass.NotEmpty(patDoc.Abstract)

	// IPC
	ass.Equal("A61F002/06", patDoc.Classifications[0].Text)

	// NATIONAL
	ass.Equal("623001160", patDoc.Classifications[1].Text)

	// inventors
	ass.Equal("Randolf", patDoc.Inventors[0].FirsName)
	ass.Equal("Von Oepen", patDoc.Inventors[0].LastName)
	ass.Equal("Los Altos Hills", patDoc.Inventors[0].City)
	ass.Equal("CA", patDoc.Inventors[0].State)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
}
