package uspto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRemoveZero(t *testing.T) {
	ass := assert.New(t)

	r := removeLeadingZero("0")
	ass.Equal("", r)

	r = removeLeadingZero("00")
	ass.Equal("0", r)

	r = removeLeadingZero("10")
	ass.Equal("10", r)
}

func TestRemoveZeros(t *testing.T) {
	ass := assert.New(t)

	r := removeLeadingZeros("0")
	ass.Equal("", r)

	r = removeLeadingZeros("00")
	ass.Equal("", r)

	r = removeLeadingZeros("A00")
	ass.Equal("A00", r)

	r = removeLeadingZeros("900")
	ass.Equal("900", r)

	r = removeLeadingZeros("009")
	ass.Equal("9", r)
}

func TestUsptoPatentDocumentSimple_GenerateAliases(t *testing.T) {
	ass := assert.New(t)
	location, err := time.LoadLocation("America/New_York")
	ass.NoError(err)
	p := UsptoPatentDocumentSimple{
		DocNumber: "001008",
		Kind:      "A1",
		Country:   Country("US"),
		DatePubl:  time.Date(2001, 4, 5, 10, 0, 0, 0, location),
	}
	p.GenerateAliases()
	ass.Contains(p.Aliases, "US2001001008A1")
}

func TestUsptoPatentDocumentSimple_GenerateAliases2(t *testing.T) {
	ass := assert.New(t)
	location, err := time.LoadLocation("America/New_York")
	ass.NoError(err)
	p := UsptoPatentDocumentSimple{
		DocNumber: "0247212",
		Kind:      "A1",
		Country:   Country("US"),
		DatePubl:  time.Date(2009, 4, 5, 10, 0, 0, 0, location),
	}
	p.GenerateAliases()
	ass.Contains(p.Aliases, "US2009247212A1")
}

func TestUsptoPatentDocumentSimple_GenerateAliases3(t *testing.T) {
	ass := assert.New(t)
	location, err := time.LoadLocation("America/New_York")
	ass.NoError(err)
	p := UsptoPatentDocumentSimple{
		DocNumber: "0034018",
		Kind:      "A1",
		Country:   Country("US"),
		DatePubl:  time.Date(2013, 4, 5, 10, 0, 0, 0, location),
	}
	p.GenerateAliases()
	ass.Contains(p.Aliases, "US2013034018A1")
}

func TestUsptoPatentDocumentSimple_GenerateAliases4(t *testing.T) {
	ass := assert.New(t)
	location, err := time.LoadLocation("America/New_York")
	ass.NoError(err)
	p := UsptoPatentDocumentSimple{
		DocNumber: "20100159919",
		Kind:      "A1",
		Country:   Country("US"),
		DatePubl:  time.Date(2010, 4, 5, 10, 0, 0, 0, location),
	}
	p.GenerateAliases()
	ass.Contains(p.Aliases, "US2010159919A1")
}

func TestUsptoPatentDocumentSimple_GenerateAliases5(t *testing.T) {
	ass := assert.New(t)
	location, err := time.LoadLocation("America/New_York")
	ass.NoError(err)
	p := UsptoPatentDocumentSimple{
		DocNumber: "20080063115",
		Kind:      "A1",
		Country:   Country("US"),
		DatePubl:  time.Date(2008, 4, 5, 10, 0, 0, 0, location),
	}
	p.GenerateAliases()
	ass.Contains(p.Aliases, "US20080063115A1")
	ass.Contains(p.Aliases, "US200863115A1")
}
