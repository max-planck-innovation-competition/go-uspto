package uspto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessXMLSimple(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/test-single.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20211019", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("B2", patDoc.Kind)
	ass.Equal("PRODUCTION", patDoc.Status)
	ass.Equal(Country("US"), patDoc.Country)
	ass.NotEmpty(patDoc.Title)
	ass.Equal(patDoc.Title[0].Text, "Wheeled hand truck")
	ass.Equal(patDoc.Title[0].Language, "en")
	ass.NotEmpty(patDoc.Claims)
	ass.NotEmpty(patDoc.Description)
	ass.NotEmpty(patDoc.Abstract)
	// ass.NotEmpty(patDoc.Citations)

	fmt.Println(patDoc.Abstract[0])
	fmt.Println(patDoc.Claims[0])
	fmt.Println(patDoc.Description[0])

}
