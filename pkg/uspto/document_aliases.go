package uspto

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
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
	if len(documentNumber) > 4 && documentNumber[:4] == yearStr {
		documentNumber = documentNumber[4:]
	}
	// if it's not part of the document add it
	// e.g. 034018 - MISSING [2013]
	aliases = append(aliases, "US"+yearStr+documentNumber)      // adds [office + year + document number]
	aliases = append(aliases, "US"+yearStr+documentNumber+kind) // adds [office + year + document number + kind]
	// remove preceding zeros
	aliases = append(aliases, "US"+yearStr+removeLeadingZero(documentNumber))       // adds [office + year + [-0{0,1}]document number]
	aliases = append(aliases, "US"+yearStr+removeLeadingZero(documentNumber)+kind)  // adds [office + year + [-0{0,1}]document number + kind]
	aliases = append(aliases, "US"+yearStr+removeLeadingZeros(documentNumber))      // adds [office + year + [-0{*}]document number]
	aliases = append(aliases, "US"+yearStr+removeLeadingZeros(documentNumber)+kind) // adds [office + year + [-0{*}]document number + kind]
	return
}

// removeLeadingZero removes ONE leading 0 from the document number
func removeLeadingZero(documentNumber string) string {
	documentNumber = strings.TrimPrefix(documentNumber, "0")
	return documentNumber
}

// regexpZeros matches all leading zeros
var regexpZeros = regexp.MustCompile(`^0+`)

// removeLeadingZeros removes ALL leading 0 from the document number
func removeLeadingZeros(documentNumber string) string {
	documentNumber = regexpZeros.ReplaceAllString(documentNumber, "")
	return documentNumber
}

func (p *UsptoPatentDocumentSimple) GenerateAliases() {
	res, err := GenerateAliases(p.DocNumber, p.Kind, p.DatePubl)
	if err != nil {
		log.Error(err)
		return
	}
	p.Aliases = res
}
