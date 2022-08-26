package uspto

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func ProcessGrantXML4Simple(doc *goquery.Document) (patentDoc UsptoPatentDocumentSimple, err error) {
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
			State:     strings.TrimSpace(c.Find("addressbook address state").Text()),
			FirstName: strings.TrimSpace(c.Find("addressbook first-name").Text()),
			LastName:  strings.TrimSpace(c.Find("addressbook last-name").Text()),
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
	representatives := root.Find("agents agent")
	representatives.Each(func(i int, c *goquery.Selection) {

		name := strings.TrimSpace(c.Find("addressbook orgname").Text())
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
	// Classifications IPCR
	/*
		<classifications-ipcr>
			<classification-ipcr>
				<ipc-version-indicator>
					<date>20060101</date>
				</ipc-version-indicator>
				<classification-level>A</classification-level>
				<section>A</section>
				<class>01</class>
				<subclass>B</subclass>
				<main-group>1</main-group>
				<subgroup>24</subgroup>
				<symbol-position>F</symbol-position>
				<classification-value>I</classification-value>
				<action-date>
					<date>20211019</date>
				</action-date>
				<generating-office>
					<country>US</country>
				</generating-office>
				<classification-status>B</classification-status>
				<classification-data-source>H</classification-data-source>
			</classification-ipcr>
		</classifications-ipcr>
	*/
	ipcr := biblio.Find("classifications-ipcr classification-ipcr")
	sequenceCounter := 1
	ipcr.Each(func(i int, c *goquery.Selection) {
		// do not use trim here
		item := ClassificationItem{
			System:                 IPC,
			Text:                   "",
			Sequence:               sequenceCounter,
			Section:                strings.TrimSpace(c.Find("section").Text()),
			Class:                  strings.TrimSpace(c.Find("class").Text()),
			SubClass:               strings.TrimSpace(c.Find("subclass").Text()),
			MainGroup:              strings.TrimSpace(c.Find("main-group").Text()),
			SubGroup:               strings.TrimSpace(c.Find("subgroup").Text()),
			Version:                strings.TrimSpace(c.Find("ipc-version-indicator date").Text()),
			ClassificationLevel:    strings.TrimSpace(c.Find("classification-level").Text()),
			FirstLater:             strings.TrimSpace(c.Find("symbol-position").Text()),
			ClassificationValue:    strings.TrimSpace(c.Find("classification-value").Text()),
			ActionDate:             strings.TrimSpace(c.Find("action-date date").Text()),
			OriginalOrReclassified: strings.TrimSpace(c.Find("classification-status").Text()), // todo: not sure about that
			Source:                 strings.TrimSpace(c.Find("classification-data-source").Text()),
			GeneratingOffice:       strings.TrimSpace(c.Find("generating-office").Text()),
		}
		patentDoc.Classifications = append(patentDoc.Classifications, item)
		sequenceCounter++
	})

	// Classifications CPC
	/*
		<classifications-cpc>
		            <main-cpc>
		                <classification-cpc>
		                    <cpc-version-indicator>
		                        <date>20130101</date>
		                    </cpc-version-indicator>
		                    <section>A</section>
		                    <class>01</class>
		                    <subclass>B</subclass>
		                    <main-group>1</main-group>
		                    <subgroup>243</subgroup>
		                    <symbol-position>F</symbol-position>
		                    <classification-value>I</classification-value>
		                    <action-date>
		                        <date>20211019</date>
		                    </action-date>
		                    <generating-office>
		                        <country>US</country>
		                    </generating-office>
		                    <classification-status>B</classification-status>
		                    <classification-data-source>H</classification-data-source>
		                    <scheme-origination-code>C</scheme-origination-code>
		                </classification-cpc>
		            </main-cpc>
		            <further-cpc>
		                <classification-cpc>
	*/
	sequenceCounter = 1
	cpc := biblio.Find("classifications-cpc classification-cpc")
	cpc.Each(func(i int, c *goquery.Selection) {
		// do not use trim here
		item := ClassificationItem{
			System:                 CPC,
			Text:                   "",
			Sequence:               sequenceCounter,
			Section:                strings.TrimSpace(c.Find("section").Text()),
			Class:                  strings.TrimSpace(c.Find("class").Text()),
			SubClass:               strings.TrimSpace(c.Find("subclass").Text()),
			MainGroup:              strings.TrimSpace(c.Find("main-group").Text()),
			SubGroup:               strings.TrimSpace(c.Find("subgroup").Text()),
			Version:                strings.TrimSpace(c.Find("cpc-version-indicator date").Text()),
			ClassificationLevel:    strings.TrimSpace(c.Find("classification-level").Text()),
			FirstLater:             strings.TrimSpace(c.Find("symbol-position").Text()),
			ClassificationValue:    strings.TrimSpace(c.Find("classification-value").Text()),
			ActionDate:             strings.TrimSpace(c.Find("action-date date").Text()),
			OriginalOrReclassified: strings.TrimSpace(c.Find("classification-status").Text()), // todo: not sure about that
			Source:                 strings.TrimSpace(c.Find("classification-data-source").Text()),
			GeneratingOffice:       strings.TrimSpace(c.Find("generating-office").Text()),
		}
		patentDoc.Classifications = append(patentDoc.Classifications, item)
		sequenceCounter++
	})

	/*
		classes := biblio.Find("us-field-of-classification-search classification-cpc-text")
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
			item := NewClassificationItemFromString(c.Find("text").Text(), seqInt)
			patentDoc.Classifications = append(patentDoc.Classifications, item)
		})
	*/
	return
}
