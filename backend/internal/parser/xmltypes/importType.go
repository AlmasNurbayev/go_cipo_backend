package xmltypes

// КоммерческаяИнформация was generated 2024-03-01 19:49:36 by https://xml-to-go.github.io/ in Ukraine.
type ImportType struct {
	КоммерческаяИнформация struct {
		XMLName          string `xml:"КоммерческаяИнформация"`
		Text             string `xml:",chardata"`
		Xmlns            string `xml:"xmlns,attr"`
		Xs               string `xml:"xs,attr"`
		Xsi              string `xml:"xsi,attr"`
		ВерсияСхемы      string `xml:"ВерсияСхемы,attr"`
		ДатаФормирования string `xml:"ДатаФормирования,attr"`
		Классификатор    struct {
			Text         string `xml:",chardata"`
			Ид           string `xml:"Ид"`
			Наименование string `xml:"Наименование"`
			Владелец     struct {
				Text         string `xml:",chardata"`
				Ид           string `xml:"Ид"`
				Наименование string `xml:"Наименование"`
				ИНН          string `xml:"ИНН"`
			} `xml:"Владелец"`
			Группы struct {
				Text   string `xml:",chardata"`
				Группа []struct {
					Text         string `xml:",chardata"`
					Ид           string `xml:"Ид"`
					Наименование string `xml:"Наименование"`
				} `xml:"Группа"`
			} `xml:"Группы"`
			Свойства struct {
				Text     string `xml:",chardata"`
				Свойство []struct {
					Text             string `xml:",chardata"`
					Ид               string `xml:"Ид"`
					Наименование     string `xml:"Наименование"`
					ТипЗначений      string `xml:"ТипЗначений"`
					ВариантыЗначений struct {
						Text       string `xml:",chardata"`
						Справочник []struct {
							Text       string `xml:",chardata"`
							ИдЗначения string `xml:"ИдЗначения"`
							Значение   string `xml:"Значение"`
						} `xml:"Справочник"`
					} `xml:"ВариантыЗначений"`
				} `xml:"Свойство"`
			} `xml:"Свойства"`
		} `xml:"Классификатор"`
		Каталог struct {
			Text                    string `xml:",chardata"`
			СодержитТолькоИзменения string `xml:"СодержитТолькоИзменения,attr"`
			Ид                      string `xml:"Ид"`
			ИдКлассификатора        string `xml:"ИдКлассификатора"`
			Наименование            string `xml:"Наименование"`
			Владелец                struct {
				Text         string `xml:",chardata"`
				Ид           string `xml:"Ид"`
				Наименование string `xml:"Наименование"`
				ИНН          string `xml:"ИНН"`
			} `xml:"Владелец"`
			Товары ТоварыType `xml:"Товары"`
		} `xml:"Каталог"`
	}
}

type ТоварыType struct {
	Text  string `xml:",chardata"`
	Товар []struct {
		Text           string   `xml:",chardata"`
		Ид             string   `xml:"Ид"`
		Артикул        string   `xml:"Артикул"`
		Наименование   string   `xml:"Наименование"`
		Картинка       []string `xml:"Картинка"`
		БазоваяЕдиница struct {
			Text                    string `xml:",chardata"`
			Код                     string `xml:"Код,attr"`
			НаименованиеПолное      string `xml:"НаименованиеПолное,attr"`
			МеждународноеСокращение string `xml:"МеждународноеСокращение,attr"`
			Пересчет                struct {
				Text        string `xml:",chardata"`
				Единица     string `xml:"Единица"`
				Коэффициент string `xml:"Коэффициент"`
			} `xml:"Пересчет"`
		} `xml:"БазоваяЕдиница"`
		Группы struct {
			Text string `xml:",chardata"`
			Ид   string `xml:"Ид"`
		} `xml:"Группы"`
		Описание        string `xml:"Описание"`
		ЗначенияСвойств struct {
			Text             string `xml:",chardata"`
			ЗначенияСвойства []struct {
				Text     string `xml:",chardata"`
				Ид       string `xml:"Ид"`
				Значение string `xml:"Значение"`
			} `xml:"ЗначенияСвойства"`
		} `xml:"ЗначенияСвойств"`
		ЗначенияРеквизитов struct {
			Text              string `xml:",chardata"`
			ЗначениеРеквизита []struct {
				Text         string `xml:",chardata"`
				Наименование string `xml:"Наименование"`
				Значение     string `xml:"Значение"`
			} `xml:"ЗначениеРеквизита"`
		} `xml:"ЗначенияРеквизитов"`
	} `xml:"Товар"`
}
