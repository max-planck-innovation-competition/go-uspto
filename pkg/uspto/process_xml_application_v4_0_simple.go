package uspto

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func ProcessApplicationXML40Simple(doc *goquery.Document) (patentDoc UsptoPatentDocumentSimple, err error) {
	if err != nil {
		return
	}
	root := doc.Find("us-patent-application")
	if root == nil {
		return
	}

	patentDoc.Lang, _ = root.Attr("lang")
	patentDoc.DtdVersion, _ = root.Attr("dtd-version")
	patentDoc.File, _ = root.Attr("file")
	patentDoc.Status, _ = root.Attr("status")

	// parse the date form the string
	dateString, _ := root.Attr("date-publ")
	parsedDate, errDate := time.Parse(layoutDatePubl, dateString)
	if errDate != nil {
		log.Warn("can not parse date", dateString)
	} else {
		patentDoc.DatePubl = parsedDate
	}

	// upto bibliography
	biblio := root.Find("us-bibliographic-data-application").First()
	pubReferenceDoc := biblio.Find("publication-reference").First()
	docId := pubReferenceDoc.Find("document-id").First()
	patentDoc.DocNumber = docId.Find("doc-number").Text()
	patentDoc.Kind = docId.Find("kind").Text()
	// build doc id
	docNumber := docId.Find("doc-number").Text()
	if len(docNumber) == 0 {
		err = errors.New("no doc number present for building id")
		log.Error(err)
		return
	}

	country := docId.Find("country").Text()
	if len(country) == 0 {
		err = errors.New("no country present for building id")
		log.Error(err)
		return
	}
	patentDoc.Country = Country(strings.TrimSpace(country))

	// set id
	patentDoc.ID = docNumber

	/* v 4-0
	<invention-title id="d0e64">Novelty jeans</invention-title>
	*/
	title := biblio.Find("invention-title")
	patentDoc.Title = append(patentDoc.Title, Title{
		Language: strings.ToLower("en"),
		Text:     strings.TrimSpace(title.Text()),
	})

	// abstract
	abstract := root.Find("abstract")
	abstract.Each(func(i int, a *goquery.Selection) {
		patentDoc.Abstract = append(
			patentDoc.Abstract,
			Abstract{
				Text:     strings.TrimSpace(a.Text()),
				Language: strings.ToLower("en"),
			},
		)
	})

	// inventors
	/*
			 <parties>
		            <applicants>
		                <applicant sequence="00" app-type="applicant-inventor" designation="us-only">
		                    <addressbook>
		                        <last-name>Goldkind</last-name>
		                        <first-name>Tina</first-name>
		                        <address>
		                            <city>St. James</city>
		                            <state>NY</state>
		                            <country>US</country>
		                        </address>
		                    </addressbook>
	*/
	parties := biblio.Find("parties")
	inventors := parties.Find("applicants applicant")
	inventors.Each(func(i int, c *goquery.Selection) {
		patentDoc.Inventors = append(patentDoc.Inventors, Inventor{
			Country: Country(strings.TrimSpace(c.Find("addressbook address country").Text())),
			City:    strings.TrimSpace(c.Find("addressbook address city").Text()),
			Street:  strings.TrimSpace(c.Find("adr str").Text()),
			Name:    strings.TrimSpace(c.Find("addressbook last-name").Text() + ", " + c.Find("addressbook first-name").Text()),
			// USPTO
			State:    strings.TrimSpace(c.Find("addressbook address state").Text()),
			FirsName: strings.TrimSpace(c.Find("addressbook first-name").Text()),
			LastName: strings.TrimSpace(c.Find("addressbook last-name").Text()),
		})
	})
	// owners
	/*
		<B720>
			<B721>
				<snm>MARTIN, Brian Alexander</snm>
				<adr>
					<str>c/o Sony Europe IP Europe Jays Close, Viables</str>
					<city>Basingstoke, Hampshire RG22 4SB</city>
					<ctry>GB</ctry>
				</adr>
			</B721>
	*/
	applicants := parties.Find("applicant")
	owners := applicants.Find("us-applicant")
	owners.Each(func(i int, c *goquery.Selection) {
		patentDoc.Owners = append(patentDoc.Owners, Owner{
			Country: Country(strings.TrimSpace(c.Find("adr ctry").Text())),
			IID:     strings.TrimSpace(c.Find("iid").Text()),
			IRF:     strings.TrimSpace(c.Find("irf").Text()),
			City:    strings.TrimSpace(c.Find("adr city").Text()),
			Street:  strings.TrimSpace(c.Find("adr str").Text()),
			Name:    strings.TrimSpace(c.Find("snm").Text()),
		})
	})
	// representatives
	/*
			 <parties>
		            <applicants>
		                <applicant sequence="00" app-type="applicant-inventor" designation="us-only">
		                    <addressbook>
		                        <last-name>Goldkind</last-name>
		                        <first-name>Tina</first-name>
		                        <address>
		                            <city>St. James</city>
		                            <state>NY</state>
		                            <country>US</country>
		                        </address>
		                    </addressbook>
	*/
	representatives := root.Find("correspondence-address")
	representatives.Each(func(i int, c *goquery.Selection) {
		name := strings.TrimSpace(c.Find("addressbook orgname").Text())
		if len(name) == 0 {
			name = strings.TrimSpace(c.Find("addressbook name").Text())
		}
		if len(name) == 0 {
			name = strings.TrimSpace(c.Find("addressbook last-name").Text() + ", " + c.Find("addressbook first-name").Text())
		}

		patentDoc.Representatives = append(patentDoc.Representatives, Representative{
			Country: Country(strings.TrimSpace(c.Find("addressbook address country").Text())),
			City:    strings.TrimSpace(c.Find("addressbook address city").Text()),
			Street:  strings.TrimSpace(c.Find("addressbook address street").Text()),
			Name:    name,
		})
	})

	// Classifications IPCR
	/*
			<classification-ipc>
		            <edition>07</edition>
		            <main-classification>A41D001/06</main-classification>
		        </classification-ipc>
	*/
	ipcr := biblio.Find("classification-ipc")
	sequenceCounter := 1
	ipcr.Each(func(i int, c *goquery.Selection) {
		// do not use trim here
		item := ClassificationItem{
			System: IPC,
			Text:   strings.TrimSpace(c.Find("main-classification").Text()),
		}
		patentDoc.Classifications = append(patentDoc.Classifications, item)
		sequenceCounter++
	})

	// Classifications National
	/*
			<classification-national>
		            <country>US</country>
		            <main-classification>002227000</main-classification>
		        </classification-national>
	*/
	national := biblio.Find("classification-national")
	sequenceCounter = 1
	national.Each(func(i int, c *goquery.Selection) {
		// do not use trim here
		item := ClassificationItem{
			System: US,
			Text:   strings.TrimSpace(c.Find("main-classification").Text()),
		}
		patentDoc.Classifications = append(patentDoc.Classifications, item)
		sequenceCounter++
	})

	// generate aliases
	patentDoc.GenerateAliases()

	return
}
