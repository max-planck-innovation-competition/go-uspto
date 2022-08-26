package uspto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessXMLSimpleVersion40B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/grant/v4-0-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("06857133", patDoc.ID)
	ass.Equal("US06857133-20050222.XML", patDoc.File)
	ass.Equal("EN", patDoc.Lang)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("06857133", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.Equal("PARALLEL-RUN", patDoc.Status)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20050222", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("v40 2004-12-02", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Method and apparatus for manufacturing and installing water resistant cover on a limb", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(151, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(4292, len(patDoc.Claims[0].Text))

	// description
	ass.Equal(30345, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(8, len(patDoc.Citations))

	ass.Equal("3820200", patDoc.Citations[0].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// representatives
	ass.Equal("Nissle, P.C., Tod R.", patDoc.Representatives[0].Name)

}

func TestProcessXMLSimpleVersion45B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/grant/v4-5-B1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("09358892", patDoc.ID)
	ass.Equal("US09358892-20160607.XML", patDoc.File)
	ass.Equal("EN", patDoc.Lang)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("09358892", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.Equal("PRODUCTION", patDoc.Status)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20160607", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("v4.5 2014-04-03", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("System and method for pre-charging a hybrid vehicle for improving reverse driving performance", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(877, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(10376, len(patDoc.Claims[0].Text))

	// description
	ass.Equal(41301, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(24, len(patDoc.Citations))

	ass.Equal("5856709", patDoc.Citations[0].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Nagoya", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[0].Street)
	ass.Equal("Gaither, Geoffrey D.", patDoc.Inventors[0].Name)
	ass.Equal("Gaither", patDoc.Inventors[0].LastName)
	ass.Equal("Geoffrey D.", patDoc.Inventors[0].FirstName)

	// representatives
	ass.Equal("Snell & Wilmer LLP", patDoc.Representatives[0].Name)

	// classifications
	// IPC
	for i := 0; i <= 5; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("20160607", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("US", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("B", patDoc.Classifications[0].Section)
	ass.Equal("60", patDoc.Classifications[0].Class)
	ass.Equal("L", patDoc.Classifications[0].SubClass)
	ass.Equal("9", patDoc.Classifications[0].MainGroup)
	ass.Equal("00", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[1].Section)
	ass.Equal("60", patDoc.Classifications[1].Class)
	ass.Equal("L", patDoc.Classifications[1].SubClass)
	ass.Equal("11", patDoc.Classifications[1].MainGroup)
	ass.Equal("00", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[2].Section)
	ass.Equal("60", patDoc.Classifications[2].Class)
	ass.Equal("L", patDoc.Classifications[2].SubClass)
	ass.Equal("11", patDoc.Classifications[2].MainGroup)
	ass.Equal("18", patDoc.Classifications[2].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[2].Version)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)
	ass.Equal("I", patDoc.Classifications[2].ClassificationValue)

	ass.Equal("G", patDoc.Classifications[3].Section)
	ass.Equal("01", patDoc.Classifications[3].Class)
	ass.Equal("C", patDoc.Classifications[3].SubClass)
	ass.Equal("21", patDoc.Classifications[3].MainGroup)
	ass.Equal("26", patDoc.Classifications[3].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[3].Version)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)
	ass.Equal("I", patDoc.Classifications[3].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[4].Section)
	ass.Equal("60", patDoc.Classifications[4].Class)
	ass.Equal("K", patDoc.Classifications[4].SubClass)
	ass.Equal("35", patDoc.Classifications[4].MainGroup)
	ass.Equal("00", patDoc.Classifications[4].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[4].Version)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)
	ass.Equal("I", patDoc.Classifications[4].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[5].Section)
	ass.Equal("60", patDoc.Classifications[5].Class)
	ass.Equal("W", patDoc.Classifications[5].SubClass)
	ass.Equal("20", patDoc.Classifications[5].MainGroup)
	ass.Equal("00", patDoc.Classifications[5].SubGroup)
	ass.Equal("20160101", patDoc.Classifications[5].Version)
	ass.Equal("L", patDoc.Classifications[5].FirstLater)
	ass.Equal("N", patDoc.Classifications[5].ClassificationValue)

	// CPC
	for i := 6; i <= 11; i++ {
		ass.Equal(CPC, patDoc.Classifications[i].System)
		ass.Equal(i-5, patDoc.Classifications[i].Sequence)

		ass.Equal("20130101", patDoc.Classifications[i].Version)
		ass.Equal("", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("20160607", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("US", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("B", patDoc.Classifications[6].Section)
	ass.Equal("60", patDoc.Classifications[6].Class)
	ass.Equal("L", patDoc.Classifications[6].SubClass)
	ass.Equal("11", patDoc.Classifications[6].MainGroup)
	ass.Equal("1809", patDoc.Classifications[6].SubGroup)
	ass.Equal("F", patDoc.Classifications[6].FirstLater)
	ass.Equal("I", patDoc.Classifications[6].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[7].Section)
	ass.Equal("60", patDoc.Classifications[7].Class)
	ass.Equal("K", patDoc.Classifications[7].SubClass)
	ass.Equal("35", patDoc.Classifications[7].MainGroup)
	ass.Equal("00", patDoc.Classifications[7].SubGroup)
	ass.Equal("L", patDoc.Classifications[7].FirstLater)
	ass.Equal("I", patDoc.Classifications[7].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[8].Section)
	ass.Equal("60", patDoc.Classifications[8].Class)
	ass.Equal("L", patDoc.Classifications[8].SubClass)
	ass.Equal("11", patDoc.Classifications[8].MainGroup)
	ass.Equal("1861", patDoc.Classifications[8].SubGroup)
	ass.Equal("L", patDoc.Classifications[8].FirstLater)
	ass.Equal("I", patDoc.Classifications[8].ClassificationValue)

	ass.Equal("G", patDoc.Classifications[9].Section)
	ass.Equal("01", patDoc.Classifications[9].Class)
	ass.Equal("C", patDoc.Classifications[9].SubClass)
	ass.Equal("21", patDoc.Classifications[9].MainGroup)
	ass.Equal("26", patDoc.Classifications[9].SubGroup)
	ass.Equal("L", patDoc.Classifications[9].FirstLater)
	ass.Equal("I", patDoc.Classifications[9].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[10].Section)
	ass.Equal("60", patDoc.Classifications[10].Class)
	ass.Equal("K", patDoc.Classifications[10].SubClass)
	ass.Equal("2350", patDoc.Classifications[10].MainGroup)
	ass.Equal("1076", patDoc.Classifications[10].SubGroup)
	ass.Equal("L", patDoc.Classifications[10].FirstLater)
	ass.Equal("A", patDoc.Classifications[10].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[11].Section)
	ass.Equal("60", patDoc.Classifications[11].Class)
	ass.Equal("W", patDoc.Classifications[11].SubClass)
	ass.Equal("20", patDoc.Classifications[11].MainGroup)
	ass.Equal("00", patDoc.Classifications[11].SubGroup)
	ass.Equal("L", patDoc.Classifications[11].FirstLater)
	ass.Equal("A", patDoc.Classifications[11].ClassificationValue)
}

func TestProcessXMLSimpleVersion45B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/grant/v4-5-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("11147202", patDoc.ID)
	ass.Equal("US11147202-20211019.XML", patDoc.File)
	ass.Equal("EN", patDoc.Lang)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("11147202", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.Equal("PRODUCTION", patDoc.Status)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20211019", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("v4.5 2014-04-03", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Wheeled hand truck", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(847, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(3974, len(patDoc.Claims[0].Text))

	// description
	ass.Equal(20201, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(41, len(patDoc.Citations))

	ass.Equal("554139", patDoc.Citations[0].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Andover", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[0].Street)
	ass.Equal("Ng, Yeow", patDoc.Inventors[0].Name)
	ass.Equal("Ng", patDoc.Inventors[0].LastName)
	ass.Equal("Yeow", patDoc.Inventors[0].FirstName)
	ass.Equal("KS", patDoc.Inventors[0].State)

	// representatives
	ass.Equal("Jack, Kenneth H.", patDoc.Representatives[0].Name)
	ass.Equal("Davis & Jack, LLC", patDoc.Representatives[1].Name)

	// classifications
	// IPC
	for i := 0; i <= 2; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("B", patDoc.Classifications[i].SubClass)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20211019", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("US", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("1", patDoc.Classifications[0].MainGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("B", patDoc.Classifications[1].Section)
	ass.Equal("62", patDoc.Classifications[1].Class)
	ass.Equal("1", patDoc.Classifications[1].MainGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("B", patDoc.Classifications[2].Section)
	ass.Equal("62", patDoc.Classifications[2].Class)
	ass.Equal("5", patDoc.Classifications[2].MainGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	// CPC
	for i := 3; i <= 7; i++ {
		ass.Equal(CPC, patDoc.Classifications[i].System)
		ass.Equal(i-2, patDoc.Classifications[i].Sequence)
		ass.Equal("B", patDoc.Classifications[i].SubClass)
		ass.Equal("20130101", patDoc.Classifications[i].Version)
		ass.Equal("", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("20211019", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("US", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("A", patDoc.Classifications[3].Section)
	ass.Equal("01", patDoc.Classifications[3].Class)
	ass.Equal("1", patDoc.Classifications[3].MainGroup)
	ass.Equal("243", patDoc.Classifications[3].SubGroup)
	ass.Equal("F", patDoc.Classifications[3].FirstLater)
	ass.Equal("I", patDoc.Classifications[3].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[4].Section)
	ass.Equal("62", patDoc.Classifications[4].Class)
	ass.Equal("1", patDoc.Classifications[4].MainGroup)
	ass.Equal("12", patDoc.Classifications[4].SubGroup)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)
	ass.Equal("I", patDoc.Classifications[4].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[5].Section)
	ass.Equal("62", patDoc.Classifications[5].Class)
	ass.Equal("5", patDoc.Classifications[5].MainGroup)
	ass.Equal("06", patDoc.Classifications[5].SubGroup)
	ass.Equal("L", patDoc.Classifications[5].FirstLater)
	ass.Equal("I", patDoc.Classifications[5].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[6].Section)
	ass.Equal("62", patDoc.Classifications[6].Class)
	ass.Equal("2202", patDoc.Classifications[6].MainGroup)
	ass.Equal("70", patDoc.Classifications[6].SubGroup)
	ass.Equal("L", patDoc.Classifications[6].FirstLater)
	ass.Equal("A", patDoc.Classifications[6].ClassificationValue)

	ass.Equal("B", patDoc.Classifications[7].Section)
	ass.Equal("62", patDoc.Classifications[7].Class)
	ass.Equal("2206", patDoc.Classifications[7].MainGroup)
	ass.Equal("006", patDoc.Classifications[7].SubGroup)
	ass.Equal("L", patDoc.Classifications[7].FirstLater)
	ass.Equal("A", patDoc.Classifications[7].ClassificationValue)

}
