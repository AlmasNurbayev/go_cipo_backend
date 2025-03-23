package xmltypes

type ClassificatorType struct {
	Text         string `xml:",chardata"`
	Ид           string `xml:"Ид"`
	Наименование string `xml:"Наименование"`
	Владелец     struct {
		Text         string `xml:",chardata"`
		Ид           string `xml:"Ид"`
		Наименование string `xml:"Наименование"`
		ИНН          string `xml:"ИНН"`
	} `xml:"Владелец"`
}
