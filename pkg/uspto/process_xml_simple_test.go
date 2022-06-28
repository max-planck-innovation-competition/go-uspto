package uspto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessXMLSimpleVersion45(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/grant/v4-5-B2.xml")
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
	// IPC
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

	// inventors
	ass.Equal("Yeow", patDoc.Inventors[0].FirsName)
	ass.Equal("Ng", patDoc.Inventors[0].LastName)
	ass.Equal("Andover", patDoc.Inventors[0].City)
	ass.Equal("KS", patDoc.Inventors[0].State)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)

	// inventors
	ass.Equal("Jack, Kenneth H.", patDoc.Representatives[0].Name)
}

func TestProcessXMLSimpleVersion25Design(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/v2-5-S1.xml")
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

func TestProcessXMLFileSimple(t *testing.T) {
	ass := assert.New(t)
	patDoc, err := ProcessXMLFileSimple("./test-data/grant/v2-5-B1_1.xml")
	ass.NoError(err)
	ass.NotEmpty(patDoc)
}

func TestProcessXMLSimpleVersion25Patent(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/grant/v2-5-B1_1.xml")
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

	// Classification
	ass.Equal(US, patDoc.Classifications[0].System)
	ass.Equal("", patDoc.Classifications[0].Version)
	ass.Equal("4675", patDoc.Classifications[0].Text)
	ass.Equal(1, patDoc.Classifications[0].Sequence)

	// FieldOfSearch
	ass.Equal(US, patDoc.FieldOfSearch[0].System)
	ass.Equal("", patDoc.FieldOfSearch[0].Version)
	ass.Equal("4675", patDoc.FieldOfSearch[0].Text)
	ass.Equal(1, patDoc.FieldOfSearch[0].Sequence)
	ass.Equal("", patDoc.FieldOfSearch[1].Version)
	ass.Equal("4676", patDoc.FieldOfSearch[1].Text)
	ass.Equal(2, patDoc.FieldOfSearch[1].Sequence)

	// Inventor
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Osamu", patDoc.Inventors[0].FirsName)
	ass.Equal("Tokunaga", patDoc.Inventors[0].LastName)
	ass.Equal("Fukuoka", patDoc.Inventors[0].City)

	// Owner
	ass.Equal(Country("JP"), patDoc.Owners[0].Country)
	ass.Equal("Toto Ltd.", patDoc.Owners[0].Name)
	ass.Equal("Fukuoka", patDoc.Owners[0].City)

	// Representative
	ass.Equal("Pillsbury Winthrop LLP", patDoc.Representatives[0].Name)
}
