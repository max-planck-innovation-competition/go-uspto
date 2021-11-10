package uspto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessXMLSimple(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/test-single.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20211019", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("B2", patDoc.Kind)
	ass.Equal("PRODUCTION", patDoc.Status)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Wheeled hand truck", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)
	ass.NotEmpty(patDoc.Claims)
	ass.NotEmpty(patDoc.Description)
	ass.NotEmpty(patDoc.Abstract)
	// ass.NotEmpty(patDoc.Citations)

	// fmt.Println(patDoc.Abstract[0])
	// fmt.Println(patDoc.Claims[0])
	// fmt.Println(patDoc.Description[0])

	ass.Equal("20060101", patDoc.IpcClassifications[0].Version)
	ass.Equal("A", patDoc.IpcClassifications[0].ClassificationLevel)
	ass.Equal(1, patDoc.IpcClassifications[0].Sequence)
	ass.Equal("A", patDoc.IpcClassifications[0].Section)
	ass.Equal("01", patDoc.IpcClassifications[0].Class)
	ass.Equal("B", patDoc.IpcClassifications[0].SubClass)
	ass.Equal("1", patDoc.IpcClassifications[0].MainGroup)
	ass.Equal("24", patDoc.IpcClassifications[0].SubGroup)
	ass.Equal("F", patDoc.IpcClassifications[0].FirstLater)
	ass.Equal("I", patDoc.IpcClassifications[0].ClassificationValue)
	ass.Equal("US", patDoc.IpcClassifications[0].GeneratingOffice)
	ass.Equal("B", patDoc.IpcClassifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.IpcClassifications[0].Source)

	ass.Equal(patDoc.IpcClassifications[1].Sequence, 2)

	ass.Equal("20130101", patDoc.CpcClassifications[0].Version)
	ass.Equal("", patDoc.CpcClassifications[0].ClassificationLevel)
	ass.Equal(1, patDoc.CpcClassifications[0].Sequence)
	ass.Equal("A", patDoc.CpcClassifications[0].Section)
	ass.Equal("01", patDoc.CpcClassifications[0].Class)
	ass.Equal("B", patDoc.CpcClassifications[0].SubClass)
	ass.Equal("1", patDoc.CpcClassifications[0].MainGroup)
	ass.Equal("243", patDoc.CpcClassifications[0].SubGroup)
	ass.Equal("F", patDoc.CpcClassifications[0].FirstLater)
	ass.Equal("I", patDoc.CpcClassifications[0].ClassificationValue)
	ass.Equal("US", patDoc.CpcClassifications[0].GeneratingOffice)
	ass.Equal("B", patDoc.CpcClassifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.CpcClassifications[0].Source)

}
