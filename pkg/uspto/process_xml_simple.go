package uspto

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const (
	layoutDatePubl = "20060102"
)

// ProcessXMLSimple transforms the raw response of the xml data into a simple patent
func ProcessXMLSimple(raw []byte) (patentDoc UsptoPatentDocumentSimple, err error) {
	// parse doc
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(raw))
	if err != nil {
		return
	}
	root := doc.Find("us-patent-grant")
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
	biblio := root.Find("us-bibliographic-data-grant").First()
	pubReferenceDoc := biblio.Find("publication-reference").First()
	docId := pubReferenceDoc.Find("document-id").First()
	patentDoc.DocNumber = docId.Find("doc-number").Text()
	patentDoc.Kind = docId.Find("kind").Text()
	// build doc id
	docNumber := docId.Find("doc-number").Text()
	country := docId.Find("country").Text()
	if len(country) == 0 {
		err = errors.New("no country present for building id")
		log.Error(err)
		return
	}
	if len(docNumber) == 0 {
		err = errors.New("no doc number present for building id")
		log.Error(err)
		return
	}
	patentDoc.ID = country + "-" + docNumber
	patentDoc.Country = Country(strings.TrimSpace(country))

	// title
	/*
		<invention-title id="d2e43">Wheeled hand truck</invention-title>
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
	// description
	description := root.Find("description")
	description.Each(func(i int, d *goquery.Selection) {
		patentDoc.Description = append(
			patentDoc.Description,
			Description{
				Text:     strings.TrimSpace(d.Text()),
				Language: strings.ToLower("en"),
			})
	})
	// claims
	/*
		<claims id="claims">
			<claim id="CLM-00002" num="00002">
				<claim-text>2. The wheeled hand truck of <claim-ref idref="CLM-00001">claim 1</claim-ref> further comprising
					a &#x201c;U&#x201d; handle having left and right arms, the left and right pins respectively comprising
					lower ends of said arms.
				</claim-text>
			</claim>
		...
		</claims>
	*/
	claims := root.Find("claims")
	claims.Each(func(i int, c *goquery.Selection) {
		lang := "en"
		id, _ := c.Attr("id")
		patentDoc.Claims = append(patentDoc.Claims, Claim{
			Text:     strings.TrimSpace(c.Text()),
			Language: strings.TrimSpace(lang),
			Id:       id,
		})
	})
	// citations
	/*
		<patcit id="ref-pcit0001" dnum="US20120281566A">
			<document-id>
				<country>US</country>
				<doc-number>20120281566</doc-number>
				<kind>A</kind>
			</document-id>
		</patcit>
		<crossref idref="pcit0001">[0006]</crossref>
	*/
	citations := root.Find("patcit")
	citations.Each(func(i int, c *goquery.Selection) {
		docId := c.Find("document-id")
		if docId.Size() > 0 {
			patentDoc.Citations = append(patentDoc.Citations, Citation{
				Country:   Country(strings.TrimSpace(docId.Find("country").Text())),
				DocNumber: strings.TrimSpace(docId.Find("doc-number").Text()),
				Kind:      strings.TrimSpace(docId.Find("kind").Text()),
			})
		}
	})
	// inventors
	/*
		<inventors>
			<inventor sequence="001" designation="us-only">
				<addressbook>
					<last-name>Ng</last-name>
					<first-name>Yeow</first-name>
					<address>
						<city>Andover</city>
						<state>KS</state>
						<country>US</country>
					</address>
				</addressbook>
			</inventor>
		</inventors>
	*/
	parties := biblio.Find("us-parties")
	inventors := parties.Find("inventors inventor")
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
	applicants := parties.Find("us-applicant")
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
		<B740>
			<B741>
				<snm>D Young & Co LLP</snm>
				<iid>101533551</iid>
				<adr>
					<str>120 Holborn</str>
					<city>London EC1N 2DY</city>
					<ctry>GB</ctry>
				</adr>
			</B741>
		</B740>
	*/
	representatives := root.Find("B741")
	representatives.Each(func(i int, c *goquery.Selection) {
		patentDoc.Representatives = append(patentDoc.Representatives, Representative{
			Country: Country(strings.TrimSpace(c.Find("adr ctry").Text())),
			IID:     strings.TrimSpace(c.Find("iid").Text()),
			City:    strings.TrimSpace(c.Find("adr city").Text()),
			Street:  strings.TrimSpace(c.Find("adr str").Text()),
			Name:    strings.TrimSpace(c.Find("snm").Text()),
		})
	})
	// ContractingStates
	/*
		<B800>
			<B840>
				<ctry>AL</ctry>
				<ctry>AT</ctry>

	*/
	countries := root.Find("B840 ctry")
	countries.Each(func(i int, c *goquery.Selection) {
		patentDoc.ContractingStates = append(patentDoc.ContractingStates, Country(strings.TrimSpace(c.Text())))
	})
	// Classifications
	/*
		<B510EP>
			<classification-ipcr sequence="1">
				<text>B60T 17/22 20060101AFI20200403BHEP</text>
			</classification-ipcr>
		</B510EP>
	*/
	classes := root.Find("B510EP classification-ipcr")
	classes.Each(func(i int, c *goquery.Selection) {
		seq, ex := c.Attr("sequence")
		if !ex {
			log.Warn("classification ipcr: seq does not exist")
		}
		seqInt, warn := strconv.Atoi(seq)
		if warn != nil {
			log.Warn("classification ipcr: can not parse seq string", warn)
		}
		// do not use trim here
		item := NewClassificationItem(c.Find("text").Text(), seqInt)
		patentDoc.Classifications = append(patentDoc.Classifications, item)
	})
	return
}
