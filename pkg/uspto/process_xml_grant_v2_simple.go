package uspto

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func ProcessXML2Simple(doc *goquery.Document) (patentDoc UsptoPatentDocumentSimple, err error) {
	root := doc.Find("PATDOC")

	// patentDoc.Lang, _ = root.Attr("lang")
	// patentDoc.File, _ = root.Attr("file")
	patentDoc.DtdVersion, _ = root.Attr("dtd")
	patentDoc.Status, _ = root.Attr("status")

	b100 := root.Find("B100")

	// B110 Document number
	docNumber := strings.TrimSpace(b100.Find("B110").Text())
	if len(docNumber) == 0 {
		err = errors.New("no doc number present for building id")
		log.Error(err)
		return
	}
	patentDoc.DocNumber = docNumber

	// B130 Kind
	/*
		Kind codes changed effective 2001-01-02 to accommodate pre-grant publication status.
		A1 - Utility Patent issued prior to January 2, 2001.
		A1 - Utility Patent Application published on or after January 2, 2001.
		A2 - Second or subsequent publication of a Utility Patent Application.
		A9 - Corrected published Utility Patent Application.
		Bn - Reexamination Certificate issued prior to January 2, 2001. NOTE: "n" represents a value 1 through 9.
		B1 - Utility Patent (no pre-grant publication) issued on or after January 2, 2001.
		B2 - Utility Patent (with pre-grant publication) issued on or after January 2, 2001.
		Cn - Reexamination Certificate issued on or after January 2, 2001. NOTE: "n" represents a value 1 through 9 denoting the publication level.
		E1 - Reissue Patent.
		Fn - Reexamination Certificate of a Reissue Patent NOTE: "n" represents a value 1 through 9 denoting the publication level.
		H1 - Statutory Invention Registration (SIR) Patent Documents. SIR documents began with the December 3, 1985 issue.
		I1 - "X" Patents issued from July 31, 1790 to July 13, 1836.
		I2 - "X" Reissue Patents issued from July 31, 1790 to July 13, 1836.
		I3 - Additional Improvements - Patents issued between 1838 and 1861.
		I4 - Defensive Publication - Documents issued from November 5, 1968 through May 5, 1987.
		I5 - Trial Voluntary Protest Program (TVPP) Patent Documents.
		NP - Non-Patent Literature.
		P1 - Plant Patent issued prior to January 2, 2001.
		P1 - Plant Patent Application published on or after January 2, 2001.
		P2 - Plant Patent (no pre-grant publication) issued on or after January 2, 2001.
		P3 - Plant Patent (with pre-grant publication) issued on or after January 2, 2001.
		P4 - Second or subsequent publication of a Plant Patent Application.
		P9 - Correction publication of a Plant Patent Application.
		S1 - Design Patent.
	*/
	patentDoc.Kind = strings.TrimSpace(b100.Find("B130").Text())

	// B140 - Document date (publication or issue)
	dateString := strings.TrimSpace(b100.Find("B140").Text())
	parsedDate, errDate := time.Parse(layoutDatePubl, dateString)
	if errDate != nil {
		log.Warn("can not parse date", dateString)
	} else {
		patentDoc.DatePubl = parsedDate
	}

	// B190 - Publishing country or organization code from WIPO Standard ST.3.
	country := strings.TrimSpace(b100.Find("B190").Text())

	if len(country) == 0 {
		err = errors.New("no country present for building id")
		log.Error(err)
		return
	}
	patentDoc.Country = Country(country)

	// build patent doc id
	patentDoc.ID = docNumber

	// title
	// B540 Title of invention
	title := root.Find("B540")
	patentDoc.Title = append(patentDoc.Title, Title{
		Language: "en",
		Text:     strings.TrimSpace(title.Text()),
	})

	// abstract
	// SDOAB Subdocument Abstract. All US patent types have an abstract;
	SDOABs := root.Find("SDOAB")
	SDOABs.Each(func(i int, SDOAB *goquery.Selection) {
		lang, _ := SDOAB.Attr("la")
		if len(lang) == 0 {
			lang = "en"
		}
		patentDoc.Abstract = append(patentDoc.Abstract, Abstract{
			Text:     strings.TrimSpace(SDOAB.Text()),
			Language: strings.TrimSpace(strings.ToLower(lang)),
		})
	})

	// description
	// SDODE Subdocument: Description of the invention.
	SDODEs := root.Find("SDODE")
	SDODEs.Each(func(i int, SDODE *goquery.Selection) {
		lang, _ := SDODE.Attr("la")
		if len(lang) == 0 {
			lang = "en"
		}
		patentDoc.Description = append(patentDoc.Description, Description{
			Text:     strings.TrimSpace(SDODE.Text()),
			Language: strings.TrimSpace(strings.ToLower(lang)),
		})
	})

	// claims
	// SDOCL Subdocument: Claims
	/*
		<SDOCL>
		        <H LVL="1">
		            <STEXT>
		                <PDAT>What is claimed is:</PDAT>
		            </STEXT>
		        </H>
		        <CL>
		            <CLM ID="CLM-00001">
		                <PARA ID="P-00157" LVL="0">
		                    <PTEXT>
		                        <PDAT>1. A faucet support member configured to releasably fix a faucet in a mounting hole that
		                            is formed in a base plate, the faucet having at least one of a base end portion and a handy
		                            spray, the faucet support member comprising:
		                        </PDAT>
		                    </PTEXT>
		                </PARA>
	*/
	SDOCLs := root.Find("SDOCL")
	SDOCLs.Each(func(i int, SDOCL *goquery.Selection) {
		lang, _ := SDOCL.Attr("la")
		if len(lang) == 0 {
			lang = "en"
		}
		clms := root.Find("CL CLM")
		clms.Each(func(i int, c *goquery.Selection) {
			id, _ := c.Attr("id")
			patentDoc.Claims = append(patentDoc.Claims, Claim{
				Text:     strings.TrimSpace(c.Text()),
				Language: strings.TrimSpace(strings.ToLower(lang)),
				Id:       id,
			})
		})
	})

	// citations
	/*
		<B560>
		                <B561>
		                    <PCIT>
		                        <DOC>
		                            <DNUM>
		                                <PDAT>3669141</PDAT>
		                            </DNUM>
		                            <DATE>
		                                <PDAT>19720600</PDAT>
		                            </DATE>
		                            <KIND>
		                                <PDAT>A</PDAT>
		                            </KIND>
		                        </DOC>
		                        <PARTY-US>
		                            <NAM>
		                                <SNM>
		                                    <STEXT>
		                                        <PDAT>Schmitt</PDAT>
		                                    </STEXT>
		                                </SNM>
		                            </NAM>
		                        </PARTY-US>
		                        <PNC>
		                            <PDAT>285208</PDAT>
		                        </PNC>
		                    </PCIT>
		                    <CITED-BY-EXAMINER/>
		                </B561>
	*/
	b560 := root.Find("B560")
	b561s := b560.Find("B561")
	b561s.Each(func(i int, b561 *goquery.Selection) {
		country := strings.TrimSpace(b561.Find("PCIT CTRY").Text())
		if len(country) == 0 {
			country = "US"
		}
		patentDoc.Citations = append(patentDoc.Citations, Citation{
			DocNumber: strings.TrimSpace(b561.Find("PCIT DNUM").Text()),
			Kind:      strings.TrimSpace(b561.Find("PCIT KIND").Text()),
			Country:   Country(country),
		})
	})

	// Classification
	sequenceCounter := 1
	B521s := root.Find("B520 B521")
	// B521: US: Original classification
	B521s.Each(func(i int, c *goquery.Selection) {
		// do not use trim here
		item := ClassificationItem{
			System:           US,
			Text:             strings.TrimSpace(c.Text()),
			Sequence:         sequenceCounter,
			Source:           "B521",
			GeneratingOffice: "US",
		}
		patentDoc.Classifications = append(patentDoc.Classifications, item)
		sequenceCounter++
	})
	sequenceCounter = 1
	B522s := root.Find("B520 B522")
	// B522: US: Cross-reference classification (official, or XR)
	B522s.Each(func(i int, c *goquery.Selection) {
		// do not use trim here
		item := ClassificationItem{
			System:           US,
			Text:             strings.TrimSpace(c.Text()),
			Source:           "B522",
			GeneratingOffice: "US",
			Sequence:         sequenceCounter,
		}
		patentDoc.Classifications = append(patentDoc.Classifications, item)
		sequenceCounter++
	})
	// Field of search
	sequenceCounter = 1
	usClasses := root.Find("B580 B582")
	usClasses.Each(func(i int, c *goquery.Selection) {
		// do not use trim here
		item := ClassificationItem{
			System:           US,
			Text:             strings.TrimSpace(c.Text()),
			Sequence:         sequenceCounter,
			Source:           "B582",
			GeneratingOffice: "US",
		}
		patentDoc.FieldOfSearch = append(patentDoc.FieldOfSearch, item)
		sequenceCounter++
	})

	// B720: Inventor information
	B721s := root.Find("B700 B720 B721")
	B721s.Each(func(i int, c *goquery.Selection) {
		// do not use trim here
		inventor := Inventor{
			Country:   Country(strings.TrimSpace(strings.ToUpper(c.Find("PARTY-US ADR CTRY").Text()))),
			City:      strings.TrimSpace(c.Find("PARTY-US ADR CITY").Text()),
			Street:    strings.TrimSpace(c.Find("PARTY-US ADR STR").Text()),
			Name:      strings.TrimSpace(c.Find("PARTY-US NAM").Text()),
			FirstName: strings.TrimSpace(c.Find("PARTY-US NAM FNM").Text()),
			LastName:  strings.TrimSpace(c.Find("PARTY-US NAM SNM").Text()),
			State:     "",
		}
		patentDoc.Inventors = append(patentDoc.Inventors, inventor)
	})

	// B730: Assignee
	B731s := root.Find("B700 B730 B731")
	B731s.Each(func(i int, c *goquery.Selection) {
		name := strings.TrimSpace(c.Find("PARTY-US NAM ONM").Text())
		if len(name) == 0 {
			name = strings.TrimSpace(c.Find("PARTY-US NAM SNM").Text() + ", " + c.Find("PARTY-US NAM FNM").Text())
		}
		// do not use trim here
		owner := Owner{
			Name:    name,
			IID:     "",
			IRF:     "",
			Country: Country(strings.TrimSpace(strings.ToUpper(c.Find("PARTY-US ADR CTRY").Text()))),
			City:    strings.TrimSpace(c.Find("PARTY-US ADR CITY").Text()),
			Street:  strings.TrimSpace(c.Find("PARTY-US ADR STR").Text()),
		}
		patentDoc.Owners = append(patentDoc.Owners, owner)
	})

	// B740: Attorney, agent, or representative.
	B741s := root.Find("B700 B740 B741")
	B741s.Each(func(i int, c *goquery.Selection) {
		name := strings.TrimSpace(c.Find("PARTY-US NAM ONM").Text())
		if len(name) == 0 {
			name = strings.TrimSpace(c.Find("PARTY-US NAM SNM").Text() + ", " + c.Find("PARTY-US NAM FNM").Text())
		}
		// do not use trim here
		rep := Representative{
			Name:    name,
			IID:     "",
			Country: Country(strings.TrimSpace(strings.ToUpper(c.Find("PARTY-US ADR CTRY").Text()))),
			City:    strings.TrimSpace(c.Find("PARTY-US ADR CITY").Text()),
			Street:  strings.TrimSpace(c.Find("PARTY-US ADR STR").Text()),
		}
		patentDoc.Representatives = append(patentDoc.Representatives, rep)
	})

	// generate aliases
	patentDoc.GenerateAliases()

	return
}
