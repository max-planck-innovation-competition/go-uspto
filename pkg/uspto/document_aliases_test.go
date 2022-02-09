package uspto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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
