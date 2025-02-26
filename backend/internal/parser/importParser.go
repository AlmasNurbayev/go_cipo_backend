package parser

import (
	"context"
	"strconv"

	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (p *Parser) ImportParser(ctx context.Context, mainStruct *xmltypes.ImportType, filePath string, newPath string) error {
	op := "parser.ImportParser"
	p.Log = p.Log.With(slog.String("op", op))
	// json, err := utils.PrintAsJSON(mainStruct)
	// if err != nil {
	// 	return err
	// }
	// p.Log.Info(string(*json))
	p.Log.Info("Starting import parsing")

	// получаем из XML и записываем регистратор
	registrator_id, err := p.service.RegistratorImportService(mainStruct, filePath, newPath)
	if err != nil {
		p.Log.Error("Error parse or saving registrator:", slog.String("err", err.Error()))
		return err
	}
	p.Log.Info("registrator_id: " + strconv.FormatInt(registrator_id, 10))

	// получаем из XML и записываем продуктовые группы
	err = p.service.ProductGroupService(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving product_groups:", slog.String("err", err.Error()))
		return err
	}

	// получаем из XML и записываем продуктовые виды моделей
	err = p.service.ProductVidService(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving product_vids:", slog.String("err", err.Error()))
		return err
	}

	// получаем из XML и записываем виды моделей
	err = p.service.VidModeliService(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving vid modeli:", slog.String("err", err.Error()))
		return err
	}

	err = p.service.ProductService(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving vid modeli:", slog.String("err", err.Error()))
		return err
	}

	// получаем из XML и записываем картинки
	err = p.service.ImageRegistryService(mainStruct, registrator_id, newPath)
	if err != nil {
		p.Log.Error("Error parse or saving images:", slog.String("err", err.Error()))
		return err
	}
	return nil
}
