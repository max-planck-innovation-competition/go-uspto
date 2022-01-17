package uspto

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func ProcessApplicationXML15Simple(doc *goquery.Document) (patentDoc UsptoPatentDocumentSimple, err error) {
	root := doc.Find("patent-application-publication")

	// patentDoc.Lang, _ = root.Attr("lang")
	// patentDoc.File, _ = root.Attr("file")
	patentDoc.DtdVersion, _ = root.Attr("dtd")
	patentDoc.Status, _ = root.Attr("status")

	biblio := root.Find("subdoc-bibliographic-information").First()

	docId := biblio.Find("document-id").First()

	// Document number
	docNumber := strings.TrimSpace(docId.Find("doc-number").Text())
	if len(docNumber) == 0 {
		err = errors.New("no doc number present for building id")
		log.Error(err)
		return
	}
	patentDoc.DocNumber = docNumber
	// build patent doc id
	patentDoc.ID = docNumber
	patentDoc.Country = Country("US")
	// kind-code
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
	patentDoc.Kind = strings.TrimSpace(docId.Find("kind-code").Text())

	// document-date
	dateString := strings.TrimSpace(docId.Find("document-date").Text())
	parsedDate, errDate := time.Parse(layoutDatePubl, dateString)
	if errDate != nil {
		log.Warn("can not parse date", dateString)
	} else {
		patentDoc.DatePubl = parsedDate
	}

	// TODO countries

	techInfo := root.Find("technical-information")

	// title
	title := techInfo.Find("title-of-invention")
	patentDoc.Title = append(patentDoc.Title, Title{
		Language: "en",
		Text:     strings.TrimSpace(title.Text()),
	})

	// abstract
	abstract := strings.TrimSpace(root.Find("subdoc-abstract").Text())
	patentDoc.Abstract = append(patentDoc.Abstract, Abstract{
		Text:     strings.TrimSpace(abstract),
		Language: strings.TrimSpace("en"),
	})

	// todo inventors
	// todo classifications

	return
}
