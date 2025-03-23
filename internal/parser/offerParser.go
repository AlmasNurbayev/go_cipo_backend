package parser

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (p *Parser) OfferParser(mainStruct *xmltypes.OfferType, filePath string, newPath string) error {
	// json, err := utils.PrintAsJSON(mainStruct)
	// if err != nil {
	// 	return err
	// }
	// p.Log.Info(string(*json))
	op := "parser.OfferParser"
	log := p.Log.With(slog.String("op", op))
	log.Info("Starting offer parsing")

	// получаем из XML и записываем регистратор
	registrator_id, err := p.service.RegistratorOfferService(mainStruct, filePath, newPath)
	if err != nil {
		log.Error("Error parse or saving registrator:", slog.String("err", err.Error()))
		return err
	}
	log.Info("registrator_id: " + strconv.FormatInt(registrator_id, 10))

	// получаем из XML и записываем размеры
	err = p.service.SizeService(mainStruct, registrator_id)
	if err != nil {
		log.Error("Error parse or saving sizes:", slog.String("err", err.Error()))
		return err
	}

	// получаем из XML и записываем виды цены
	err = p.service.PriceVidService(mainStruct, registrator_id)
	if err != nil {
		log.Error("Error parse or saving price vids:", slog.String("err", err.Error()))
		return err
	}

	// получаем из XML и записываем склады
	err = p.service.StoreService(mainStruct, registrator_id)
	if err != nil {
		log.Error("Error parse or saving stores:", slog.String("err", err.Error()))
		return err
	}

	err = p.service.QntPriveRegistryService(mainStruct, registrator_id)
	if err != nil {
		log.Error("Error parse or saving qnt price registry:", slog.String("err", err.Error()))
		return err
	}

	return nil
}
