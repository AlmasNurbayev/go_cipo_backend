package parserService

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) QntPriveRegistryService(mainStruct *xmltypes.OfferType,
	registrator_id int64) error {

	op := "parserService.QntPriveRegistryService"
	log := s.log.With(slog.String("op", op))

	products, err := s.storage.ListProduct(s.ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	priceVids, err := s.storage.ListPriceVid(s.ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	sizes, err := s.storage.ListSize(s.ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	stores, err := s.storage.ListStore(s.ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	registrator, err := s.storage.GetRegistratorById(s.ctx, registrator_id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	operation_date := registrator.Operation_date

	NewQntPrices, err := partParsers.QntPriceRegistryParser(mainStruct, registrator_id,
		operation_date, stores, priceVids, sizes, products)
	if err != nil {
		log.Error("Error parsing qnt_prices:", slog.String("err", err.Error()))
		return err
	}
	for _, val := range NewQntPrices {
		_, err := s.storage.CreateQntPriceRegistry(s.ctx, val)
		if err != nil {
			log.Error("Error inserting qnt_price_registry:", slog.String("err", err.Error()))
			return err
		}
	}
	log.Info("qnt_prices add count: " + strconv.Itoa(len(NewQntPrices)))
	return nil

}
