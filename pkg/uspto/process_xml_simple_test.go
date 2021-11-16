package uspto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessXMLSimpleVersion45(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/4-5-b2-patent-example.xml")
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

	// CPC
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("B", patDoc.Classifications[0].SubClass)
	ass.Equal("1", patDoc.Classifications[0].MainGroup)
	ass.Equal("24", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("US", patDoc.Classifications[0].GeneratingOffice)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)

	ass.Equal(patDoc.Classifications[1].Sequence, 2)

	ass.Equal("20130101", patDoc.Classifications[3].Version)
	ass.Equal("", patDoc.Classifications[3].ClassificationLevel)
	ass.Equal(1, patDoc.Classifications[3].Sequence)
	ass.Equal("A", patDoc.Classifications[3].Section)
	ass.Equal("01", patDoc.Classifications[3].Class)
	ass.Equal("B", patDoc.Classifications[3].SubClass)
	ass.Equal("1", patDoc.Classifications[3].MainGroup)
	ass.Equal("243", patDoc.Classifications[3].SubGroup)
	ass.Equal("F", patDoc.Classifications[3].FirstLater)
	ass.Equal("I", patDoc.Classifications[3].ClassificationValue)
	ass.Equal("US", patDoc.Classifications[3].GeneratingOffice)
	ass.Equal("B", patDoc.Classifications[3].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[3].Source)

}

func TestProcessXMLSimpleVersion25Design(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/2-5-s1-design-patent-example.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20020101", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("D0452599", patDoc.DocNumber)
	ass.Equal("S1", patDoc.Kind)
	ass.Equal("Build 20020101", patDoc.Status)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Confection item in the form of a tied braid", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	/*
		ass.NotEmpty(patDoc.Claims)
		ass.NotEmpty(patDoc.Description)
		ass.NotEmpty(patDoc.Abstract)
		// ass.NotEmpty(patDoc.Citations)

		// fmt.Println(patDoc.Abstract[0])
		// fmt.Println(patDoc.Claims[0])
		// fmt.Println(patDoc.Description[0])

		ass.Equal("20060101", patDoc.Classifications[3].Version)
		ass.Equal("A", patDoc.Classifications[3].ClassificationLevel)
		ass.Equal(1, patDoc.Classifications[3].Sequence)
		ass.Equal("A", patDoc.Classifications[3].Section)
		ass.Equal("01", patDoc.Classifications[3].Class)
		ass.Equal("B", patDoc.Classifications[3].SubClass)
		ass.Equal("1", patDoc.Classifications[3].MainGroup)
		ass.Equal("24", patDoc.Classifications[3].SubGroup)
		ass.Equal("F", patDoc.Classifications[3].FirstLater)
		ass.Equal("I", patDoc.Classifications[3].ClassificationValue)
		ass.Equal("US", patDoc.Classifications[3].GeneratingOffice)
		ass.Equal("B", patDoc.Classifications[3].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[3].Source)

		ass.Equal(patDoc.Classifications[1].Sequence, 2)

		ass.Equal("20130101", patDoc.Classifications[3].Version)
		ass.Equal("", patDoc.Classifications[3].ClassificationLevel)
		ass.Equal(1, patDoc.Classifications[3].Sequence)
		ass.Equal("A", patDoc.Classifications[3].Section)
		ass.Equal("01", patDoc.Classifications[3].Class)
		ass.Equal("B", patDoc.Classifications[3].SubClass)
		ass.Equal("1", patDoc.Classifications[3].MainGroup)
		ass.Equal("243", patDoc.Classifications[3].SubGroup)
		ass.Equal("F", patDoc.Classifications[3].FirstLater)
		ass.Equal("I", patDoc.Classifications[3].ClassificationValue)
		ass.Equal("US", patDoc.Classifications[3].GeneratingOffice)
		ass.Equal("B", patDoc.Classifications[3].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[3].Source)
	*/
}

func TestProcessXMLSimpleVersion5Patent(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/2-5-b1-patent.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20020101", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("06334226", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.Equal("Build 20020101", patDoc.Status)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Faucet support member", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	ass.NotEmpty(patDoc.Claims)
	ass.NotEmpty(patDoc.Description)
	ass.NotEmpty(patDoc.Abstract)
	ass.NotEmpty(patDoc.Citations)

	ass.Equal("3669141", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	/*
		//

		// fmt.Println(patDoc.Abstract[0])
		// fmt.Println(patDoc.Claims[0])
		// fmt.Println(patDoc.Description[0])

		ass.Equal("20060101", patDoc.Classifications[3].Version)
		ass.Equal("A", patDoc.Classifications[3].ClassificationLevel)
		ass.Equal(1, patDoc.Classifications[3].Sequence)
		ass.Equal("A", patDoc.Classifications[3].Section)
		ass.Equal("01", patDoc.Classifications[3].Class)
		ass.Equal("B", patDoc.Classifications[3].SubClass)
		ass.Equal("1", patDoc.Classifications[3].MainGroup)
		ass.Equal("24", patDoc.Classifications[3].SubGroup)
		ass.Equal("F", patDoc.Classifications[3].FirstLater)
		ass.Equal("I", patDoc.Classifications[3].ClassificationValue)
		ass.Equal("US", patDoc.Classifications[3].GeneratingOffice)
		ass.Equal("B", patDoc.Classifications[3].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[3].Source)

		ass.Equal(patDoc.Classifications[1].Sequence, 2)

		ass.Equal("20130101", patDoc.Classifications[3].Version)
		ass.Equal("", patDoc.Classifications[3].ClassificationLevel)
		ass.Equal(1, patDoc.Classifications[3].Sequence)
		ass.Equal("A", patDoc.Classifications[3].Section)
		ass.Equal("01", patDoc.Classifications[3].Class)
		ass.Equal("B", patDoc.Classifications[3].SubClass)
		ass.Equal("1", patDoc.Classifications[3].MainGroup)
		ass.Equal("243", patDoc.Classifications[3].SubGroup)
		ass.Equal("F", patDoc.Classifications[3].FirstLater)
		ass.Equal("I", patDoc.Classifications[3].ClassificationValue)
		ass.Equal("US", patDoc.Classifications[3].GeneratingOffice)
		ass.Equal("B", patDoc.Classifications[3].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[3].Source)
	*/
}
