package uspto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessXMLSimpleVersion25B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/grant/v2-5-B1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("06334226", patDoc.ID)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("06334226", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.Equal("Build 20020101", patDoc.Status)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20020101", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("2.5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Faucet support member", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(1255, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(31, len(patDoc.Claims))

	for i := 0; i <= 8; i++ {
		ass.Equal("en", patDoc.Claims[i].Language)
		ass.Equal(fmt.Sprintf("CLM-0000%v", i+1), patDoc.Claims[i].Id)
	}
	for i := 9; i <= 30; i++ {
		ass.Equal("en", patDoc.Claims[i].Language)
		ass.Equal(fmt.Sprintf("CLM-000%v", i+1), patDoc.Claims[i].Id)
	}

	ass.Equal(3060, len(patDoc.Claims[0].Text))
	ass.Equal(227, len(patDoc.Claims[1].Text))
	ass.Equal(233, len(patDoc.Claims[2].Text))
	ass.Equal(554, len(patDoc.Claims[3].Text))
	ass.Equal(571, len(patDoc.Claims[4].Text))
	ass.Equal(865, len(patDoc.Claims[5].Text))
	ass.Equal(402, len(patDoc.Claims[6].Text))
	ass.Equal(405, len(patDoc.Claims[7].Text))
	ass.Equal(453, len(patDoc.Claims[8].Text))
	ass.Equal(216, len(patDoc.Claims[9].Text))
	ass.Equal(584, len(patDoc.Claims[10].Text))
	ass.Equal(1201, len(patDoc.Claims[11].Text))
	ass.Equal(402, len(patDoc.Claims[12].Text))
	ass.Equal(236, len(patDoc.Claims[13].Text))
	ass.Equal(231, len(patDoc.Claims[14].Text))
	ass.Equal(206, len(patDoc.Claims[15].Text))
	ass.Equal(418, len(patDoc.Claims[16].Text))
	ass.Equal(289, len(patDoc.Claims[17].Text))
	ass.Equal(317, len(patDoc.Claims[18].Text))
	ass.Equal(402, len(patDoc.Claims[19].Text))
	ass.Equal(512, len(patDoc.Claims[20].Text))
	ass.Equal(237, len(patDoc.Claims[21].Text))
	ass.Equal(236, len(patDoc.Claims[22].Text))
	ass.Equal(399, len(patDoc.Claims[23].Text))
	ass.Equal(272, len(patDoc.Claims[24].Text))
	ass.Equal(313, len(patDoc.Claims[25].Text))
	ass.Equal(323, len(patDoc.Claims[26].Text))
	ass.Equal(279, len(patDoc.Claims[27].Text))
	ass.Equal(396, len(patDoc.Claims[28].Text))
	ass.Equal(342, len(patDoc.Claims[29].Text))
	ass.Equal(3353, len(patDoc.Claims[30].Text))

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(175171, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(14, len(patDoc.Citations))

	ass.Equal("3669141", patDoc.Citations[0].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Fukuoka", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[0].Street)
	ass.Equal("Osamu\n                            \n                            \n                                \n                                    Tokunaga", patDoc.Inventors[0].Name)
	ass.Equal("Osamu", patDoc.Inventors[0].FirstName)
	ass.Equal("Tokunaga", patDoc.Inventors[0].LastName)

	ass.NotEmpty(patDoc.Inventors[1])
	ass.Equal(Country("JP"), patDoc.Inventors[1].Country)
	ass.Equal("Fukuoka", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[1].Street)
	ass.Equal("Hideyuki\n                            \n                            \n                                \n                                    Hasebe", patDoc.Inventors[1].Name)
	ass.Equal("Hideyuki", patDoc.Inventors[1].FirstName)
	ass.Equal("Hasebe", patDoc.Inventors[1].LastName)

	ass.NotEmpty(patDoc.Inventors[2])
	ass.Equal(Country("JP"), patDoc.Inventors[2].Country)
	ass.Equal("Fukuoka", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[2].Street)
	ass.Equal("Sigemiti\n                            \n                            \n                                \n                                    Inada", patDoc.Inventors[2].Name)
	ass.Equal("Sigemiti", patDoc.Inventors[2].FirstName)
	ass.Equal("Inada", patDoc.Inventors[2].LastName)

	// owners
	ass.NotEmpty(patDoc.Owners[0])
	ass.Equal(Country("JP"), patDoc.Owners[0].Country)
	ass.Equal("Fukuoka", patDoc.Owners[0].City)
	ass.Equal("Toto Ltd.", patDoc.Owners[0].Name)

	// representatives
	ass.Equal("Pillsbury Winthrop LLP", patDoc.Representatives[0].Name)

	// classifications
	ass.Equal(US, patDoc.Classifications[0].System)
	ass.Equal("4675", patDoc.Classifications[0].Text)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("B521", patDoc.Classifications[0].Source)
	ass.Equal("US", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal(US, patDoc.Classifications[1].System)
	ass.Equal("4676", patDoc.Classifications[1].Text)
	ass.Equal(1, patDoc.Classifications[1].Sequence)
	ass.Equal("B522", patDoc.Classifications[1].Source)
	ass.Equal("US", patDoc.Classifications[1].GeneratingOffice)

	ass.Equal(US, patDoc.Classifications[2].System)
	ass.Equal("4678", patDoc.Classifications[2].Text)
	ass.Equal(2, patDoc.Classifications[2].Sequence)
	ass.Equal("B522", patDoc.Classifications[2].Source)
	ass.Equal("US", patDoc.Classifications[2].GeneratingOffice)

	ass.Equal(US, patDoc.Classifications[3].System)
	ass.Equal("4695", patDoc.Classifications[3].Text)
	ass.Equal(3, patDoc.Classifications[3].Sequence)
	ass.Equal("B522", patDoc.Classifications[3].Source)
	ass.Equal("US", patDoc.Classifications[3].GeneratingOffice)

	ass.Equal(US, patDoc.Classifications[4].System)
	ass.Equal("4696", patDoc.Classifications[4].Text)
	ass.Equal(4, patDoc.Classifications[4].Sequence)
	ass.Equal("B522", patDoc.Classifications[4].Source)
	ass.Equal("US", patDoc.Classifications[4].GeneratingOffice)

	// field of search
	for i := 0; i <= 15; i++ {
		ass.Equal(US, patDoc.FieldOfSearch[i].System)
		ass.Equal(i+1, patDoc.FieldOfSearch[i].Sequence)
		ass.Equal("B582", patDoc.FieldOfSearch[i].Source)
		ass.Equal("US", patDoc.FieldOfSearch[i].GeneratingOffice)
	}

	ass.Equal("4675", patDoc.FieldOfSearch[0].Text)
	ass.Equal("4676", patDoc.FieldOfSearch[1].Text)
	ass.Equal("4677", patDoc.FieldOfSearch[2].Text)
	ass.Equal("4678", patDoc.FieldOfSearch[3].Text)
	ass.Equal("4670", patDoc.FieldOfSearch[4].Text)
	ass.Equal("4695", patDoc.FieldOfSearch[5].Text)
	ass.Equal("4696", patDoc.FieldOfSearch[6].Text)
	ass.Equal("137359", patDoc.FieldOfSearch[7].Text)
	ass.Equal("137801", patDoc.FieldOfSearch[8].Text)
	ass.Equal("137360", patDoc.FieldOfSearch[9].Text)
	ass.Equal("137606", patDoc.FieldOfSearch[10].Text)
	ass.Equal("403373", patDoc.FieldOfSearch[11].Text)
	ass.Equal("403187", patDoc.FieldOfSearch[12].Text)
	ass.Equal("403199", patDoc.FieldOfSearch[13].Text)
	ass.Equal("4034081", patDoc.FieldOfSearch[14].Text)
	ass.Equal("4033742", patDoc.FieldOfSearch[15].Text)
	ass.Equal("4033745", patDoc.FieldOfSearch[16].Text)

}

func TestProcessXMLSimpleVersion25B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("./test-data/grant/v2-5-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("06530116", patDoc.ID)
	ass.Equal(Country("US"), patDoc.Country)
	ass.Equal("06530116", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.Equal("Build 20030116", patDoc.Status)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20030311", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("2.5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Vacuum cleaner with muffled detachable blower exhaust", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[0].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(860, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(20, len(patDoc.Claims))

	for i := 0; i <= 8; i++ {
		ass.Equal("en", patDoc.Claims[i].Language)
		ass.Equal(fmt.Sprintf("CLM-0000%v", i+1), patDoc.Claims[i].Id)
	}
	for i := 9; i <= 19; i++ {
		ass.Equal("en", patDoc.Claims[i].Language)
		ass.Equal(fmt.Sprintf("CLM-000%v", i+1), patDoc.Claims[i].Id)
	}

	ass.Equal(850, len(patDoc.Claims[0].Text))
	ass.Equal(209, len(patDoc.Claims[1].Text))
	ass.Equal(169, len(patDoc.Claims[2].Text))
	ass.Equal(397, len(patDoc.Claims[3].Text))
	ass.Equal(335, len(patDoc.Claims[4].Text))
	ass.Equal(287, len(patDoc.Claims[5].Text))
	ass.Equal(348, len(patDoc.Claims[6].Text))
	ass.Equal(389, len(patDoc.Claims[7].Text))
	ass.Equal(847, len(patDoc.Claims[8].Text))
	ass.Equal(211, len(patDoc.Claims[9].Text))
	ass.Equal(170, len(patDoc.Claims[10].Text))
	ass.Equal(399, len(patDoc.Claims[11].Text))
	ass.Equal(337, len(patDoc.Claims[12].Text))
	ass.Equal(289, len(patDoc.Claims[13].Text))
	ass.Equal(1143, len(patDoc.Claims[14].Text))
	ass.Equal(874, len(patDoc.Claims[15].Text))
	ass.Equal(171, len(patDoc.Claims[16].Text))
	ass.Equal(399, len(patDoc.Claims[17].Text))
	ass.Equal(196, len(patDoc.Claims[18].Text))
	ass.Equal(212, len(patDoc.Claims[19].Text))

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(15073, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(33, len(patDoc.Citations))

	ass.Equal("2276844", patDoc.Citations[0].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// inventors
	ass.NotEmpty(patDoc.Inventors[0])
	ass.Equal("Jersey Shore", patDoc.Inventors[0].City)
	ass.Empty(patDoc.Inventors[0].Street)
	ass.Equal("Robert C.Berfield", patDoc.Inventors[0].Name)
	ass.Equal("Robert C.", patDoc.Inventors[0].FirstName)
	ass.Equal("Berfield", patDoc.Inventors[0].LastName)

	ass.NotEmpty(patDoc.Inventors[1])
	ass.Equal("Williamsport", patDoc.Inventors[1].City)
	ass.Empty(patDoc.Inventors[1].Street)
	ass.Equal("RonaldGriffin", patDoc.Inventors[1].Name)
	ass.Equal("Ronald", patDoc.Inventors[1].FirstName)
	ass.Equal("Griffin", patDoc.Inventors[1].LastName)

	ass.NotEmpty(patDoc.Inventors[2])
	ass.Equal("Avis", patDoc.Inventors[2].City)
	ass.Empty(patDoc.Inventors[2].Street)
	ass.Equal("CraigSeasholtz", patDoc.Inventors[2].Name)
	ass.Equal("Craig", patDoc.Inventors[2].FirstName)
	ass.Equal("Seasholtz", patDoc.Inventors[2].LastName)

	// owners
	ass.NotEmpty(patDoc.Owners[0])
	ass.Equal("Williamsport", patDoc.Owners[0].City)
	ass.Equal("Shop Vac Corporation", patDoc.Owners[0].Name)

	// classifications
	ass.Equal(US, patDoc.Classifications[0].System)
	ass.Equal("15328", patDoc.Classifications[0].Text)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("B521", patDoc.Classifications[0].Source)
	ass.Equal("US", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal(US, patDoc.Classifications[1].System)
	ass.Equal("153276", patDoc.Classifications[1].Text)
	ass.Equal(1, patDoc.Classifications[1].Sequence)
	ass.Equal("B522", patDoc.Classifications[1].Source)
	ass.Equal("US", patDoc.Classifications[1].GeneratingOffice)

	ass.Equal(US, patDoc.Classifications[2].System)
	ass.Equal("15326", patDoc.Classifications[2].Text)
	ass.Equal(2, patDoc.Classifications[2].Sequence)
	ass.Equal("B522", patDoc.Classifications[2].Source)
	ass.Equal("US", patDoc.Classifications[2].GeneratingOffice)

	ass.Equal(US, patDoc.Classifications[3].System)
	ass.Equal("15412", patDoc.Classifications[3].Text)
	ass.Equal(3, patDoc.Classifications[3].Sequence)
	ass.Equal("B522", patDoc.Classifications[3].Source)
	ass.Equal("US", patDoc.Classifications[3].GeneratingOffice)

	// field of search
	for i := 0; i <= 4; i++ {
		ass.Equal(US, patDoc.FieldOfSearch[i].System)
		ass.Equal(i+1, patDoc.FieldOfSearch[i].Sequence)
		ass.Equal("B582", patDoc.FieldOfSearch[i].Source)
		ass.Equal("US", patDoc.FieldOfSearch[i].GeneratingOffice)
	}

	ass.Equal("15328", patDoc.FieldOfSearch[0].Text)
	ass.Equal("153276", patDoc.FieldOfSearch[1].Text)
	ass.Equal("15353", patDoc.FieldOfSearch[2].Text)
	ass.Equal("15326", patDoc.FieldOfSearch[3].Text)
	ass.Equal("15412", patDoc.FieldOfSearch[4].Text)

}
