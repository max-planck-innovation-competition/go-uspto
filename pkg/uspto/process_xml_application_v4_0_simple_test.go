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

	ass.Equal("20050000001", patDoc.ID)
	ass.Equal("US20050000001A1-20050106.XML", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("20050000001", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20050106", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("PARALLEL-RUN", patDoc.Status)
	ass.Equal("v4.0 2004-12-02", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Novelty jeans", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(504, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal("Goldkind, Tina", patDoc.Inventors[0].Name)
	ass.Equal("Tina", patDoc.Inventors[0].FirstName)
	ass.Equal("Goldkind", patDoc.Inventors[0].LastName)
	ass.Equal("St. James", patDoc.Inventors[0].City)
	ass.Equal("NY", patDoc.Inventors[0].State)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Empty(patDoc.Inventors[0].Street)

	// representatives
	ass.Equal(Country("US"), patDoc.Representatives[0].Country)
	ass.Empty(patDoc.Representatives[0].IID)
	ass.Equal("NEW YORK", patDoc.Representatives[0].City)
	ass.Empty(patDoc.Representatives[0].Street)
	ass.Equal("SCHWEITZER CORNMAN GROSS & BONDELL LLP", patDoc.Representatives[0].Name)

	// classifications
	// IPC
	ass.Equal(ClassificationSystem("IPC"), patDoc.Classifications[0].System)
	ass.Equal("A41D001/06", patDoc.Classifications[0].Text)
	ass.Equal(0, patDoc.Classifications[0].Sequence)
	// national
	ass.Equal(ClassificationSystem("US"), patDoc.Classifications[1].System)
	ass.Equal("002227000", patDoc.Classifications[1].Text)
	ass.Equal(0, patDoc.Classifications[1].Sequence)

}

func TestApplicationProcessXMLSimpleVersion40A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/v4-0-A2_1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("20050004655", patDoc.ID)
	ass.Equal("US20050004655A2-20050106.XML", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("20050004655", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20050106", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("PARALLEL-RUN", patDoc.Status)
	ass.Equal("v4.0 2004-12-02", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("METHODS AND APPARATUS FOR A STENT HAVING AN EXPANDABLE WEB STRUCTURE", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(552, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal("Von Oepen, Randolf", patDoc.Inventors[0].Name)
	ass.Equal("Randolf", patDoc.Inventors[0].FirstName)
	ass.Equal("Von Oepen", patDoc.Inventors[0].LastName)
	ass.Equal("Los Altos Hills", patDoc.Inventors[0].City)
	ass.Equal("CA", patDoc.Inventors[0].State)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Empty(patDoc.Inventors[0].Street)

	ass.NotEmpty(patDoc.Inventors[1])
	ass.Equal("Seibold, Gerd", patDoc.Inventors[1].Name)
	ass.Equal("Gerd", patDoc.Inventors[1].FirstName)
	ass.Equal("Seibold", patDoc.Inventors[1].LastName)
	ass.Equal("Ammerbuch", patDoc.Inventors[1].City)
	ass.Equal(Country("DE"), patDoc.Inventors[1].Country)
	ass.Empty(patDoc.Inventors[1].Street)
	ass.Empty(patDoc.Inventors[1].State)

	// representatives
	ass.Equal(Country("US"), patDoc.Representatives[0].Country)
	ass.Empty(patDoc.Representatives[0].IID)
	ass.Equal("SAN DIEGO", patDoc.Representatives[0].City)
	ass.Empty(patDoc.Representatives[0].Street)
	ass.Equal("LUCE, FORWARD, HAMILTON & SCRIPPS LLP", patDoc.Representatives[0].Name)

	// classifications
	// IPC
	ass.Equal(ClassificationSystem("IPC"), patDoc.Classifications[0].System)
	ass.Equal("A61F002/06", patDoc.Classifications[0].Text)
	ass.Equal(0, patDoc.Classifications[0].Sequence)
	// national
	ass.Equal(ClassificationSystem("US"), patDoc.Classifications[1].System)
	ass.Equal("623001160", patDoc.Classifications[1].Text)
	ass.Equal(0, patDoc.Classifications[1].Sequence)

}
