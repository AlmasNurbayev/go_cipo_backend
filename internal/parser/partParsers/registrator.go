package partParsers

import (
	"log/slog"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

// ищет в структуре вложенную структуру "Классификатор", "Каталог" и возвращат ее поля
// а также сведения о загружаемом файле
func RegistratorParserFromImport(default_user_id int64, receiveStruct *xmltypes.ImportType, filePath string, newPath string, log *slog.Logger) (models.RegistratorEntity, error) {

	mainStruct := (*receiveStruct)
	var registratorStruct models.RegistratorEntity

	registratorStruct.Name_folder = newPath
	registratorStruct.Name_file = filePath
	registratorStruct.User_id = default_user_id
	registratorStruct.Id_catalog = mainStruct.КоммерческаяИнформация.Каталог.Ид
	registratorStruct.Id_class = mainStruct.КоммерческаяИнформация.Классификатор.Ид
	registratorStruct.Name_catalog = mainStruct.КоммерческаяИнформация.Каталог.Наименование
	registratorStruct.Name_class = mainStruct.КоммерческаяИнформация.Классификатор.Наименование
	registratorStruct.Operation_date = time.Now()
	registratorStruct.Ver_schema = mainStruct.КоммерческаяИнформация.ВерсияСхемы
	if mainStruct.КоммерческаяИнформация.Каталог.СодержитТолькоИзменения == "false" {
		registratorStruct.Only_change = false
	} else {
		registratorStruct.Only_change = true
	}

	layout := "2006-01-02T15:04:05"
	var time, err = time.Parse(layout, mainStruct.КоммерческаяИнформация.ДатаФормирования)
	if err != nil {
		log.Error("Error parsing time:", slog.String("err", err.Error()))
		return registratorStruct, err
	}
	registratorStruct.Date_schema = time
	//utils.PrintAsJSON(registratorStruct)

	return registratorStruct, nil
}

func RegistratorParserFromOffer(default_user_id int64, receiveStruct *xmltypes.OfferType, filePath string, newPath string, log *slog.Logger) (models.RegistratorEntity, error) {
	// parser.Parser()
	mainStruct := (*receiveStruct)
	var registratorStruct models.RegistratorEntity

	registratorStruct.Name_folder = newPath
	registratorStruct.Name_file = filePath
	registratorStruct.User_id = default_user_id
	registratorStruct.Id_catalog = mainStruct.КоммерческаяИнформация.Классификатор.Наименование
	registratorStruct.Id_class = mainStruct.КоммерческаяИнформация.Классификатор.Ид
	registratorStruct.Name_catalog = mainStruct.КоммерческаяИнформация.ПакетПредложений.Наименование
	registratorStruct.Name_class = mainStruct.КоммерческаяИнформация.ПакетПредложений.Ид
	registratorStruct.Operation_date = time.Now()
	registratorStruct.Ver_schema = mainStruct.КоммерческаяИнформация.ВерсияСхемы
	if mainStruct.КоммерческаяИнформация.ПакетПредложений.СодержитТолькоИзменения == "false" {
		registratorStruct.Only_change = false
	} else {
		registratorStruct.Only_change = true
	}

	layout := "2006-01-02T15:04:05"
	var time, err = time.Parse(layout, mainStruct.КоммерческаяИнформация.ДатаФормирования)
	if err != nil {
		log.Error("Error parsing time:", slog.String("err", err.Error()))
		return registratorStruct, err
	}
	registratorStruct.Date_schema = time
	//utils.PrintAsJSON(registratorStruct)

	return registratorStruct, nil
}
