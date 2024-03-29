package uspto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestApplicationProcessXMLSimpleVersion44A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/v4-4-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("20210267109", patDoc.ID)
	ass.Equal("US20210267109A1-20210902.XML", patDoc.File)
	ass.Equal("EN", patDoc.Lang)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("20210267109", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210902", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("PRODUCTION", patDoc.Status)
	ass.Equal("v4.4 2014-04-03", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("METHODS AND APPARATUS FOR YARD AND GARDEN DEBRIS COLLECTION", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(353, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(7764, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(53204, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal("Jefferson Valley", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[0].Street)
	ass.Equal("O'Keefe, Daniel", patDoc.Inventors[0].Name)
	ass.Equal("Daniel", patDoc.Inventors[0].FirstName)
	ass.Equal("O'Keefe", patDoc.Inventors[0].LastName)
	ass.Equal("NY", patDoc.Inventors[0].State)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)

	// Classifications
	// IPC
	for i := 0; i <= 3; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[0].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20210902", patDoc.Classifications[0].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("US", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("B", patDoc.Classifications[0].SubClass)
	ass.Equal("1", patDoc.Classifications[0].MainGroup)
	ass.Equal("04", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("B", patDoc.Classifications[1].Section)
	ass.Equal("08", patDoc.Classifications[1].Class)
	ass.Equal("B", patDoc.Classifications[1].SubClass)
	ass.Equal("1", patDoc.Classifications[1].MainGroup)
	ass.Equal("00", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("B", patDoc.Classifications[2].Section)
	ass.Equal("08", patDoc.Classifications[2].Class)
	ass.Equal("B", patDoc.Classifications[2].SubClass)
	ass.Equal("13", patDoc.Classifications[2].MainGroup)
	ass.Equal("00", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("A", patDoc.Classifications[3].Section)
	ass.Equal("01", patDoc.Classifications[3].Class)
	ass.Equal("D", patDoc.Classifications[3].SubClass)
	ass.Equal("11", patDoc.Classifications[3].MainGroup)
	ass.Equal("06", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	// CPC
	for i := 4; i <= 7; i++ {
		ass.Equal(CPC, patDoc.Classifications[i].System)
		ass.Equal(i-3, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[0].Version)
		ass.Equal("", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20210902", patDoc.Classifications[0].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("US", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("A", patDoc.Classifications[4].Section)
	ass.Equal("01", patDoc.Classifications[4].Class)
	ass.Equal("B", patDoc.Classifications[4].SubClass)
	ass.Equal("1", patDoc.Classifications[4].MainGroup)
	ass.Equal("04", patDoc.Classifications[4].SubGroup)
	ass.Equal("F", patDoc.Classifications[4].FirstLater)

	ass.Equal("A", patDoc.Classifications[5].Section)
	ass.Equal("01", patDoc.Classifications[5].Class)
	ass.Equal("D", patDoc.Classifications[5].SubClass)
	ass.Equal("11", patDoc.Classifications[5].MainGroup)
	ass.Equal("06", patDoc.Classifications[5].SubGroup)
	ass.Equal("L", patDoc.Classifications[5].FirstLater)

	ass.Equal("B", patDoc.Classifications[6].Section)
	ass.Equal("08", patDoc.Classifications[6].Class)
	ass.Equal("B", patDoc.Classifications[6].SubClass)
	ass.Equal("13", patDoc.Classifications[6].MainGroup)
	ass.Equal("00", patDoc.Classifications[6].SubGroup)
	ass.Equal("L", patDoc.Classifications[6].FirstLater)

	ass.Equal("B", patDoc.Classifications[7].Section)
	ass.Equal("08", patDoc.Classifications[7].Class)
	ass.Equal("B", patDoc.Classifications[7].SubClass)
	ass.Equal("1", patDoc.Classifications[7].MainGroup)
	ass.Equal("005", patDoc.Classifications[7].SubGroup)
	ass.Equal("L", patDoc.Classifications[7].FirstLater)

}

func TestApplicationProcessXMLSimpleVersion44A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/v4-4-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("20150075015", patDoc.ID)
	ass.Equal("US20150075015A2-20150319.XML", patDoc.File)
	ass.Equal("EN", patDoc.Lang)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("20150075015", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20150319", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("PRODUCTION", patDoc.Status)
	ass.Equal("v4.4 2014-04-03", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("RAZOR WITH CUTTING BLADE ROTATABLE ABOUT MULTIPLE AXES", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(800, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(5031, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(16684, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// ass.NotEmpty(patDoc.Citations)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal("Setauket", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[0].Street)
	ass.Equal("Bucco, Thomas J.", patDoc.Inventors[0].Name)
	ass.Equal("Thomas J.", patDoc.Inventors[0].FirstName)
	ass.Equal("Bucco", patDoc.Inventors[0].LastName)
	ass.Equal("NY", patDoc.Inventors[0].State)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)

	// Classifications
	// IPC
	ass.Equal(ClassificationSystem("IPC"), patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("B", patDoc.Classifications[0].Section)
	ass.Equal("26", patDoc.Classifications[0].Class)
	ass.Equal("B", patDoc.Classifications[0].SubClass)
	ass.Equal("21", patDoc.Classifications[0].MainGroup)
	ass.Equal("52", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20150319", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("US", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal(ClassificationSystem("CPC"), patDoc.Classifications[1].System)
	ass.Equal(1, patDoc.Classifications[1].Sequence)
	ass.Equal("B", patDoc.Classifications[1].Section)
	ass.Equal("26", patDoc.Classifications[1].Class)
	ass.Equal("B", patDoc.Classifications[1].SubClass)
	ass.Equal("21", patDoc.Classifications[1].MainGroup)
	ass.Equal("521", patDoc.Classifications[1].SubGroup)
	ass.Equal("20130101", patDoc.Classifications[1].Version)
	ass.Equal("", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20150319", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("US", patDoc.Classifications[1].GeneratingOffice)

}

func TestApplicationProcessXMLSimpleVersion45A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/v4-5-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("20220159901", patDoc.ID)
	ass.Equal("US20220159901A1-20220526.XML", patDoc.File)
	ass.Equal("EN", patDoc.Lang)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("20220159901", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20220526", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("PRODUCTION", patDoc.Status)
	ass.Equal("v4.5 2021-08-30", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("SEED DISTRIBUTOR USED FOR AGRICULTURAL MACHINES", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(1009, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(4846, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(25753, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// ass.NotEmpty(patDoc.Citations)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal(Country("AR"), patDoc.Inventors[0].Country)
	ass.Equal("Provincia Santa Fe", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[0].Street)
	ass.Equal("Spinsanti, Ariel Alberto Antonio", patDoc.Inventors[0].Name)
	ass.Equal("Ariel Alberto Antonio", patDoc.Inventors[0].FirstName)
	ass.Equal("Spinsanti", patDoc.Inventors[0].LastName)
	ass.Equal("", patDoc.Inventors[0].State)

	// Classifications
	// IPC
	ass.Equal(ClassificationSystem("IPC"), patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("C", patDoc.Classifications[0].SubClass)
	ass.Equal("7", patDoc.Classifications[0].MainGroup)
	ass.Equal("04", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20220526", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("US", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal(ClassificationSystem("CPC"), patDoc.Classifications[1].System)
	ass.Equal(1, patDoc.Classifications[1].Sequence)
	ass.Equal("A", patDoc.Classifications[1].Section)
	ass.Equal("01", patDoc.Classifications[1].Class)
	ass.Equal("C", patDoc.Classifications[1].SubClass)
	ass.Equal("7", patDoc.Classifications[1].MainGroup)
	ass.Equal("04", patDoc.Classifications[1].SubGroup)
	ass.Equal("20130101", patDoc.Classifications[1].Version)
	ass.Equal("", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20220526", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("US", patDoc.Classifications[1].GeneratingOffice)

}

func TestApplicationProcessXMLSimpleVersion45A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/v4-5-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("20220160114", patDoc.ID)
	ass.Equal("US20220160114A2-20220526.XML", patDoc.File)
	ass.Equal("EN", patDoc.Lang)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("20220160114", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20220526", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("PRODUCTION", patDoc.Status)
	ass.Equal("v4.5 2021-08-30", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("CARRYING CASE WITH PIVOTING HOUSING", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(604, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1737, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(12219, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// ass.NotEmpty(patDoc.Citations)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal(Country("CA"), patDoc.Inventors[0].Country)
	ass.Equal("Vancouver", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[0].Street)
	ass.Equal("Bevis, Matthew", patDoc.Inventors[0].Name)
	ass.Equal("Matthew", patDoc.Inventors[0].FirstName)
	ass.Equal("Bevis", patDoc.Inventors[0].LastName)
	ass.Empty(patDoc.Inventors[0].State)

	// Classifications
	// IPC
	ass.Equal(ClassificationSystem("IPC"), patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("45", patDoc.Classifications[0].Class)
	ass.Equal("F", patDoc.Classifications[0].SubClass)
	ass.Equal("3", patDoc.Classifications[0].MainGroup)
	ass.Equal("04", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20220526", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("US", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal(ClassificationSystem("IPC"), patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("A", patDoc.Classifications[1].Section)
	ass.Equal("45", patDoc.Classifications[1].Class)
	ass.Equal("F", patDoc.Classifications[1].SubClass)
	ass.Equal("3", patDoc.Classifications[1].MainGroup)
	ass.Equal("02", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20220526", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("US", patDoc.Classifications[1].GeneratingOffice)
	//CPC
	ass.Equal(ClassificationSystem("CPC"), patDoc.Classifications[2].System)
	ass.Equal(1, patDoc.Classifications[2].Sequence)
	ass.Equal("A", patDoc.Classifications[2].Section)
	ass.Equal("45", patDoc.Classifications[2].Class)
	ass.Equal("F", patDoc.Classifications[2].SubClass)
	ass.Equal("3", patDoc.Classifications[2].MainGroup)
	ass.Equal("047", patDoc.Classifications[2].SubGroup)
	ass.Equal("20130101", patDoc.Classifications[2].Version)
	ass.Equal("", patDoc.Classifications[2].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[2].FirstLater)
	ass.Equal("I", patDoc.Classifications[2].ClassificationValue)
	ass.Equal("20220526", patDoc.Classifications[2].ActionDate)
	ass.Equal("B", patDoc.Classifications[2].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[2].Source)
	ass.Equal("US", patDoc.Classifications[2].GeneratingOffice)

	ass.Equal(ClassificationSystem("CPC"), patDoc.Classifications[3].System)
	ass.Equal(2, patDoc.Classifications[3].Sequence)
	ass.Equal("A", patDoc.Classifications[3].Section)
	ass.Equal("45", patDoc.Classifications[3].Class)
	ass.Equal("F", patDoc.Classifications[3].SubClass)
	ass.Equal("3", patDoc.Classifications[3].MainGroup)
	ass.Equal("02", patDoc.Classifications[3].SubGroup)
	ass.Equal("20130101", patDoc.Classifications[3].Version)
	ass.Equal("", patDoc.Classifications[3].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)
	ass.Equal("I", patDoc.Classifications[3].ClassificationValue)
	ass.Equal("20220526", patDoc.Classifications[3].ActionDate)
	ass.Equal("B", patDoc.Classifications[3].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[3].Source)
	ass.Equal("US", patDoc.Classifications[3].GeneratingOffice)

}
