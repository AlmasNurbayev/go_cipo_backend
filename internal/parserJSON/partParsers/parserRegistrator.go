package partParsers

import (
	"context"
	"log/slog"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
)

type storageR interface {
	CreateRegistrator(ctx context.Context, registrator models.RegistratorEntity) (int64, error)
}

// заполняем и записываем регистратор
func ParserRegistrator(Cfg *config.Config, Log *slog.Logger, storage storageR, data dto.StockJSON, jsonPath string) (registrator_id int64, err error) {
	op := "parserJSON.parserRegistrator"
	Log = Log.With(slog.String("op", op))

	var registratorStruct models.RegistratorEntity
	registratorStruct.Name_folder = jsonPath
	registratorStruct.Name_file = jsonPath
	registratorStruct.User_id = Cfg.Parser.PARSER_DEFAULT_USER_ID
	registratorStruct.Id_catalog = data.BasePrefix
	registratorStruct.Id_class = data.BasePrefix
	registratorStruct.Name_catalog = data.NameLog
	registratorStruct.Name_class = data.NameLog
	registratorStruct.Operation_date = time.Now()
	registratorStruct.Ver_schema = data.BasePrefix
	registratorStruct.Only_change = false

	location := time.Local
	layout := "2006-01-02T15:04:05"
	date, err := time.ParseInLocation(layout, data.ExportDate, location)
	if err != nil {
		Log.Error("Error parsing time:", slog.String("err", err.Error()))
		return 0, err
	}
	registratorStruct.Date_schema = date

	registrator_id, err2 := storage.CreateRegistrator(context.Background(), registratorStruct)
	if err2 != nil {
		Log.Error("Error inserting registrator:", slog.String("err", err2.Error()))
		return 0, err2
	}

	return registrator_id, nil
}
