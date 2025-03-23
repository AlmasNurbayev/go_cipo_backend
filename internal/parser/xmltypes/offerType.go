package xmltypes

// КоммерческаяИнформация was generated 2024-03-01 19:46:09 by https://xml-to-go.github.io/ in Ukraine.
type OfferType struct {
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
			Свойства struct {
				Text     string `xml:",chardata"`
				Свойство struct {
					Text             string `xml:",chardata"`
					ДляПредложений   string `xml:"ДляПредложений"`
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
		ПакетПредложений struct {
			Text                    string `xml:",chardata"`
			СодержитТолькоИзменения string `xml:"СодержитТолькоИзменения,attr"`
			Ид                      string `xml:"Ид"`
			Наименование            string `xml:"Наименование"`
			ИдКаталога              string `xml:"ИдКаталога"`
			ИдКлассификатора        string `xml:"ИдКлассификатора"`
			Владелец                struct {
				Text         string `xml:",chardata"`
				Ид           string `xml:"Ид"`
				Наименование string `xml:"Наименование"`
				ИНН          string `xml:"ИНН"`
			} `xml:"Владелец"`
			ТипыЦен struct {
				Text    string `xml:",chardata"`
				ТипЦены []struct {
					Text         string `xml:",chardata"`
					Ид           string `xml:"Ид"`
					Наименование string `xml:"Наименование"`
					Валюта       string `xml:"Валюта"`
					Налог        struct {
						Text         string `xml:",chardata"`
						Наименование string `xml:"Наименование"`
						УчтеноВСумме string `xml:"УчтеноВСумме"`
						Акциз        string `xml:"Акциз"`
					} `xml:"Налог"`
				} `xml:"ТипЦены"`
			} `xml:"ТипыЦен"`
			Склады struct {
				Text  string `xml:",chardata"`
				Склад []struct {
					Text         string `xml:",chardata"`
					Ид           string `xml:"Ид"`
					Наименование string `xml:"Наименование"`
					Контакты     struct {
						Text    string `xml:",chardata"`
						Контакт struct {
							Text     string `xml:",chardata"`
							Тип      string `xml:"Тип"`
							Значение string `xml:"Значение"`
						} `xml:"Контакт"`
					} `xml:"Контакты"`
				} `xml:"Склад"`
			} `xml:"Склады"`
			Предложения ПредложенияType `xml:"Предложения"`
		} `xml:"ПакетПредложений"`
	}
}

type ПредложенияType struct {
	Text        string `xml:",chardata"`
	Предложение []struct {
		Text           string `xml:",chardata"`
		Ид             string `xml:"Ид"`
		Артикул        string `xml:"Артикул"`
		Наименование   string `xml:"Наименование"`
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
		ЗначенияСвойств struct {
			Text             string `xml:",chardata"`
			ЗначенияСвойства []struct {
				Text         string `xml:",chardata"`
				Ид           string `xml:"Ид"`
				Наименование string `xml:"Наименование"`
				Значение     string `xml:"Значение"`
			} `xml:"ЗначенияСвойства"`
		} `xml:"ЗначенияСвойств"`
		ХарактеристикиТовара struct {
			Text                 string `xml:",chardata"`
			ХарактеристикаТовара struct {
				Text         string `xml:",chardata"`
				Ид           string `xml:"Ид"`
				Наименование string `xml:"Наименование"`
				Значение     string `xml:"Значение"`
			} `xml:"ХарактеристикаТовара"`
		} `xml:"ХарактеристикиТовара"`
		Цены struct {
			Text string `xml:",chardata"`
			Цена []struct {
				Text          string  `xml:",chardata"`
				Представление string  `xml:"Представление"`
				ИдТипаЦены    string  `xml:"ИдТипаЦены"`
				ЦенаЗаЕдиницу float32 `xml:"ЦенаЗаЕдиницу"`
				Валюта        string  `xml:"Валюта"`
				Коэффициент   string  `xml:"Коэффициент"`
			} `xml:"Цена"`
		} `xml:"Цены"`
		Количество float32 `xml:"Количество"`
		Склад      []struct {
			Text               string  `xml:",chardata"`
			ИдСклада           string  `xml:"ИдСклада,attr"`
			КоличествоНаСкладе float32 `xml:"КоличествоНаСкладе,attr"`
		} `xml:"Склад"`
	} `xml:"Предложение"`
}
