package uspto

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
)

const (
	layoutDatePubl = "20060102"
)

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

	// Applications
	// version 4 and above
	root = doc.Find("us-patent-application")
	if len(root.Nodes) > 0 {
		return ProcessApplicationXML4Simple(doc)
	}

	return
}
