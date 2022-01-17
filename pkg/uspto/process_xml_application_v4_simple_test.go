package uspto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestApplicationProcessXMLSimpleVersion45(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/4-4-a1-patent-application-example.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210902", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("A1", patDoc.Kind)
	ass.Equal("PRODUCTION", patDoc.Status)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal("METHODS AND APPARATUS FOR YARD AND GARDEN DEBRIS COLLECTION", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)
	ass.NotEmpty(patDoc.Claims)
	ass.NotEmpty(patDoc.Description)
	ass.NotEmpty(patDoc.Abstract)
	// ass.NotEmpty(patDoc.Citations)

	// fmt.Println(patDoc.Abstract[0])
	// fmt.Println(patDoc.Claims[0])
	// fmt.Println(patDoc.Description[0])

	// IPC
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("B", patDoc.Classifications[0].SubClass)
	ass.Equal("1", patDoc.Classifications[0].MainGroup)
	ass.Equal("04", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("US", patDoc.Classifications[0].GeneratingOffice)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)

	ass.Equal(patDoc.Classifications[1].Sequence, 2)
	// IPC
	ass.Equal("20130101", patDoc.Classifications[5].Version)
	ass.Equal("", patDoc.Classifications[5].ClassificationLevel)
	ass.Equal(2, patDoc.Classifications[5].Sequence)
	ass.Equal("A", patDoc.Classifications[5].Section)
	ass.Equal("01", patDoc.Classifications[5].Class)
	ass.Equal("D", patDoc.Classifications[5].SubClass)
	ass.Equal("11", patDoc.Classifications[5].MainGroup)
	ass.Equal("06", patDoc.Classifications[5].SubGroup)
	ass.Equal("L", patDoc.Classifications[5].FirstLater)
	ass.Equal("I", patDoc.Classifications[5].ClassificationValue)
	ass.Equal("US", patDoc.Classifications[5].GeneratingOffice)
	ass.Equal("B", patDoc.Classifications[5].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[5].Source)

	// inventors
	ass.Equal("Daniel", patDoc.Inventors[0].FirsName)
	ass.Equal("O'Keefe", patDoc.Inventors[0].LastName)
	ass.Equal("Jefferson Valley", patDoc.Inventors[0].City)
	ass.Equal("NY", patDoc.Inventors[0].State)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
}
