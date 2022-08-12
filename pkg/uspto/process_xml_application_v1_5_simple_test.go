package uspto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessApplicationXML15Simple1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/v1-5-A1_1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20010315", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("A1", patDoc.Kind)
	ass.Equal("20010000001", patDoc.DocNumber)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Solvent mixture for use in a vapor degreaser and method of cleaning an article in a vapor degreaser utilizing said solvent", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal("The invention provides a solvent mixture comprising n-propyl bromide, a mixture of low boiling solvents and, preferably, a defluxing and/or ionics removing additive and/or at least one saturated terpene. The invention also provides a method of cleaning an article (e.g., an electrical, plastic, or metal part) in a vapor degreaser using the solvent mixture. The solvent mixture of the invention is non-flammable, non-corrosive, and non-hazardous. In addition, it has a high solvency and a very low ozone depletion potential. Thus, using the solvent mixture of the invention, oil, grease, rosin flux, and other organic material can be readily removed from the article of interest in an environmentally safe manner.", patDoc.Abstract[0].Text)
	ass.Equal("en", patDoc.Abstract[0].Language)
	ass.Equal("20010000001", patDoc.ID)
	return
}
func TestProcessApplicationXML15Simple2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/application/v1-5-A1_2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20010315", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("A1", patDoc.Kind)
	ass.Equal("20010000002", patDoc.DocNumber)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Process for detecting and correcting a fiber orientation cross direction profile change", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal("Process and a device for detecting a change in the fiber orientation cross-direction profile of a paper or cardboard web in the manufacturing process on a paper or cardboard machine. It is recognized that a change in the fiber orientation cross-direction profile in the web is inferred or determined by way of a characteristic change in the basis weight cross-direction profile or at least one measurement quantity that correlates to it, and/or by way of a characteristic change in the basis weight in the travel direction of the web or at least one measurement quantity that correlates to it. The device includes an arrangement for detecting a change in the basis weight cross-direction profile and/or a characteristic chronological course of the change in the basis weight, which is characteristic for a change in the fiber orientation.", patDoc.Abstract[0].Text)
	ass.Equal("en", patDoc.Abstract[0].Language)
	ass.Equal("20010000002", patDoc.ID)
	return
}
