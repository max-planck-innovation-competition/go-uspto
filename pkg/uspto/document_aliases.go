package uspto

import (
	"strconv"
)

// GenerateAliases generates alternative ids base on the information in the patent
// usually it is the document number as well as the kind and the year
// sometimes leading zeros are removed
// #############################################################################
// Applications
// 0061044 ->
// US 2009 0061044 A1: with the year and the kind
// US 3904033 A: without year
// #############################################################################
// Grants
// 9265015 ->
// US 9265015 B2: without year
// US 9265015: without kind
// #############################################################################
// Examples for real document number values
// 06335824 | B1 US
// 20010000059 | A1 US
// 20140001198 | A1 US
// 20140001198 | A1 US
// 20130334279 | A1 US
// #############################################################################
// Examples for identifiers that must match
// US2013034018A1
func (p *UsptoPatentDocumentSimple) GenerateAliases() {
	p.Aliases = []string{}
	// generate initial aliases
	p.Aliases = append(p.Aliases, p.DocNumber)             // adds [document number]
	p.Aliases = append(p.Aliases, "US"+p.DocNumber)        // adds [office + document number]
	p.Aliases = append(p.Aliases, "US"+p.DocNumber+p.Kind) // adds [office + document number + kind]

	// add: year
	// check if document publication year is already part of the document number
	yearStr := strconv.Itoa(p.DatePubl.Year())
	if len(p.DocNumber) > 4 && p.DocNumber[:4] != yearStr {
		// if it's not part of the document add it
		p.Aliases = append(p.Aliases, "US"+yearStr+p.DocNumber)        // adds [office + document number + year]
		p.Aliases = append(p.Aliases, "US"+yearStr+p.DocNumber+p.Kind) // adds [office + document number + year + kind]
	}

	// remove leading zero if doc number is 7 digits long
	// e.g. 0034018 --> 034018
	if len(p.DocNumber) == 7 && p.DocNumber[0] == '0' {
		p.Aliases = append(p.Aliases, "US"+p.DocNumber[1:])                // adds [office + document number without leading zero]
		p.Aliases = append(p.Aliases, "US"+p.DocNumber[1:]+p.Kind)         // adds [office + document number without leading zero + kind]
		p.Aliases = append(p.Aliases, "US"+yearStr+p.DocNumber[1:])        // adds [office + year +document number without leading zero]
		p.Aliases = append(p.Aliases, "US"+yearStr+p.DocNumber[1:]+p.Kind) // adds [office + document number without leading zero + kind]
	}
}
