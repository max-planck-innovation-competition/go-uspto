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
	aliasesMap := map[string]struct{}{}
	// generate initial aliases
	aliasesMap[documentNumber] = struct{}{}           // adds [document number]
	aliasesMap["US"+documentNumber] = struct{}{}      // adds [office + document number]
	aliasesMap["US"+documentNumber+kind] = struct{}{} // adds [office + document number + kind]

	// add: year
	// check if document publication year is already part of the document number
	yearStr := strconv.Itoa(publicationDate.Year())
	if len(documentNumber) > 4 && documentNumber[:4] == yearStr {
		documentNumber = documentNumber[4:]
	}
	// if it's not part of the document add it
	// e.g. 034018 - MISSING [2013]
	aliasesMap["US"+yearStr+documentNumber] = struct{}{}      // adds [office + year + document number]
	aliasesMap["US"+yearStr+documentNumber+kind] = struct{}{} // adds [office + year + document number + kind]
	// remove preceding zeros
	aliasesMap["US"+yearStr+removeLeadingZero(documentNumber)] = struct{}{}       // adds [office + year + [-0{0,1}]document number]
	aliasesMap["US"+yearStr+removeLeadingZero(documentNumber)+kind] = struct{}{}  // adds [office + year + [-0{0,1}]document number + kind]
	aliasesMap["US"+yearStr+removeLeadingZeros(documentNumber)] = struct{}{}      // adds [office + year + [-0{*}]document number]
	aliasesMap["US"+yearStr+removeLeadingZeros(documentNumber)+kind] = struct{}{} // adds [office + year + [-0{*}]document number + kind]

	// convert alias map to string array
	for k, _ := range aliasesMap {
		aliases = append(aliases, k)
	}
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
