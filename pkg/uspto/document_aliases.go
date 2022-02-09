package uspto

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var ErrEmptyDocumentNumber = errors.New("empty document number")
var ErrEmptyKind = errors.New("empty kind")
var ErrEmptyPublicationDate = errors.New("empty publication date")

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
func GenerateAliases(documentNumber, kind string, publicationDate time.Time) (aliases []string, err error) {
	// check input of document number
	if len(documentNumber) == 0 {
		err = ErrEmptyDocumentNumber
		return
	}
	// check input of document number
	if len(kind) == 0 {
		err = ErrEmptyKind
		return
	}
	// check date
	if publicationDate.IsZero() {
		err = ErrEmptyPublicationDate
		return
	}
	// init result
	aliases = []string{}
	// generate initial aliases
	aliases = append(aliases, documentNumber)           // adds [document number]
	aliases = append(aliases, "US"+documentNumber)      // adds [office + document number]
	aliases = append(aliases, "US"+documentNumber+kind) // adds [office + document number + kind]

	// add: year
	// check if document publication year is already part of the document number
	yearStr := strconv.Itoa(publicationDate.Year())
	if len(documentNumber) > 4 && documentNumber[:4] != yearStr {
		// if it's not part of the document add it
		aliases = append(aliases, "US"+yearStr+documentNumber)      // adds [office + document number + year]
		aliases = append(aliases, "US"+yearStr+documentNumber+kind) // adds [office + document number + year + kind]
	}

	// remove leading zero if doc number is 7 digits long
	// e.g. 0034018 --> 034018
	if len(documentNumber) == 7 && documentNumber[0] == '0' {
		aliases = append(aliases, "US"+documentNumber[1:])              // adds [office + document number without leading zero]
		aliases = append(aliases, "US"+documentNumber[1:]+kind)         // adds [office + document number without leading zero + kind]
		aliases = append(aliases, "US"+yearStr+documentNumber[1:])      // adds [office + year +document number without leading zero]
		aliases = append(aliases, "US"+yearStr+documentNumber[1:]+kind) // adds [office + document number without leading zero + kind]
	}
	return
}

func (p *UsptoPatentDocumentSimple) GenerateAliases() {
	res, err := GenerateAliases(p.DocNumber, p.Kind, p.DatePubl)
	if err != nil {
		log.Error(err)
		return
	}
	p.Aliases = res
}
