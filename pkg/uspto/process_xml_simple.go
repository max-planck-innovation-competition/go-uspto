package uspto

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"strings"
)

const (
	layoutDatePubl = "20060102"
)

var ErrCanNotFindParser = errors.New("can not find parser for this document")

func ProcessXMLFileSimple(filePath string) (patentDoc UsptoPatentDocumentSimple, err error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	return ProcessXMLSimple(data)
}

// ProcessXMLSimple transforms the raw response of the xml data into a simple patent
func ProcessXMLSimple(raw []byte) (patentDoc UsptoPatentDocumentSimple, err error) {
	// parse doc
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(raw))

	// ##################################################################
	// Grants
	// version 2.4 and 2.5
	root := doc.Find("PATDOC")
	if len(root.Nodes) > 0 {
		return ProcessXML2Simple(doc)
	}
	// version 4 and above
	root = doc.Find("us-patent-grant")
	if len(root.Nodes) > 0 {
		return ProcessGrantXML4Simple(doc)
	}

	// ##################################################################
	// Applications
	// version 4 and above
	root = doc.Find("us-patent-application")
	if len(root.Nodes) > 0 {
		// get dtd version
		dtdVersion, attPresent := root.Attr("dtd-version")
		// clean
		dtdVersion = strings.TrimSpace(dtdVersion)
		if attPresent {
			switch dtdVersion {
			case "v4.0 2004-12-02":
				return ProcessApplicationXML40Simple(doc)
			}
		}
		return ProcessApplicationXML4Simple(doc)
	}
	// version 1-5
	root = doc.Find("patent-application-publication")
	if len(root.Nodes) > 0 {
		return ProcessApplicationXML15Simple(doc)
	}

	// If no parser is found, return error
	err = ErrCanNotFindParser

	return
}
