package parserService

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) RegistratorImportService(mainStruct *xmltypes.ImportType,
	filePath string, newPath string) (int64, error) {

	op := "RegistratorImportService"
	log := s.log.With(slog.String("op", op))

	registrator, err := partParsers.RegistratorParserFromImport(s.cfg.Parser.PARSER_DEFAULT_USER_ID, mainStruct, filePath, newPath, s.log)
	if err != nil {
		log.Error("Error parsing registrator:", slog.String("err", err.Error()))
		return 0, err
	}
	registrator_id, err := s.storage.CreateRegistrator(s.ctx, registrator)
	//registrator_id, err := registerRepo.Create(*registrator)
	if err != nil {
		log.Error("Error inserting registrators:", slog.String("err", err.Error()))
		return 0, err
	}
	log.Info("added registrator with id: " + strconv.Itoa(int(registrator_id)))
	return registrator_id, nil
}

func (s *ParserService) RegistratorOfferService(mainStruct *xmltypes.OfferType,
	filePath string, newPath string) (int64, error) {

	op := "RegistratorOfferService"
	log := s.log.With(slog.String("op", op))

	registrator, err := partParsers.RegistratorParserFromOffer(s.cfg.Parser.PARSER_DEFAULT_USER_ID, mainStruct, filePath, newPath, s.log)
	if err != nil {
		log.Error("Error parsing registrator:", slog.String("err", err.Error()))
		return 0, err
	}
	registrator_id, err := s.storage.CreateRegistrator(s.ctx, registrator)
	//registrator_id, err := registerRepo.Create(*registrator)
	if err != nil {
		log.Error("Error inserting registrators:", slog.String("err", err.Error()))
		return 0, err
	}
	log.Info("added registrator with id: " + strconv.Itoa(int(registrator_id)))
	return registrator_id, nil
}
