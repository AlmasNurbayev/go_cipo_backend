package parser

import (
	"context"
	"encoding/xml"
	"log/slog"
	"os"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/logger"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/moved"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/parserService"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/storage/postgres"
)

type Parser struct {
	Version string
	Cfg     *config.Config
	storage *postgres.Storage
	Log     *slog.Logger
	service *parserService.ParserService
}

func New(version string) *Parser {
	return &Parser{
		Version: version,
	}
}

func (p *Parser) Init() {
	op := "parser.Init"

	p.Version = "v0.1.0"
	p.Cfg = config.MustLoad()
	p.Log = logger.InitLogger(p.Cfg.Env, nil)
	p.Log.With("op", op)
	p.Log.Info("==== init parser on env: " + p.Cfg.Env)
	p.Log.Debug("debug message is enabled")
	postgresStorage, err := postgres.NewStorage(p.Cfg.Dsn, p.Log, p.Cfg.HTTP.HTTP_WRITE_TIMEOUT)
	if err != nil {
		p.Log.Error(err.Error())
		panic(err)
	}
	p.storage = postgresStorage
}

func (p *Parser) Run() {
	op := "parser.Run"
	p.Log.With("op", op).Info("start parser")

	movedFiles, err := moved.MovedInputFiles(p.Cfg, p.Log)

	// var err error = nil
	// movedFiles := &moved.MovedInputFilesT{
	// 	NewPath: "/input/",
	// 	Files: []moved.InputFilesT{
	// 		{TypeFile: "classificator", PathFile: "input/import0_1.xml"},
	// 		{TypeFile: "offer", PathFile: "input/offers0_1.xml"},
	// 		{TypeFile: "imageFolder", PathFile: "import_files"},
	// 	},
	// }

	if err != nil {
		p.Log.Error(err.Error())
		os.Exit(1)
	}
	assets_path := "assets"
	if p.Cfg.Parser.PARSER_ASSETS_PATH != "" {
		assets_path = p.Cfg.Parser.PARSER_ASSETS_PATH
	}
	if err = moved.CopyImages(assets_path, movedFiles.NewPath, p.Log); err != nil {
		p.Log.Error(err.Error())
	}
	p.Log.Info("movedFiles", slog.Any("movedFiles", movedFiles))

	ctx, cancel := context.WithTimeout(context.Background(), p.Cfg.HTTP.HTTP_WRITE_TIMEOUT)
	defer cancel()
	pgxTransaction, err := p.storage.Db.Begin(ctx)
	if err != nil {
		p.Log.Error("Not created transaction:", slog.String("err", err.Error()))
		os.Exit(1)
	}
	p.storage.Tx = &pgxTransaction
	p.service = parserService.NewParserService(ctx, p.storage, p.Log, p.Cfg)

	for index, fileItem := range movedFiles.Files {
		if index == 2 {
			// папку с картинками не парсим
			continue
		}

		file, err := os.Open(fileItem.PathFile)
		if err != nil {
			p.Log.Error("Error open file:", slog.String("err", err.Error()))
			os.Exit(1)
		}
		defer file.Close()

		switch fileItem.TypeFile {
		case "classificator":
			var temp xmltypes.ImportType             // создаем экземпляр структуры
			xmlStruct := temp.КоммерческаяИнформация // создаем подчиненный узел для декодирования
			decoder := xml.NewDecoder(file)
			decoder.Strict = false
			if err := decoder.Decode(&xmlStruct); err != nil {
				p.Log.Error("Error decode file:", slog.String("err", err.Error()))
				panic(err)
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			err := p.ImportParser(ctx, &xmltypes.ImportType{КоммерческаяИнформация: xmlStruct}, fileItem.PathFile, movedFiles.NewPath)
			if err != nil {
				p.Log.Error("Error import, now rollback all db changes:", slog.String("err", err.Error()))
				err = pgxTransaction.Rollback(p.storage.Ctx)
				if err != nil {
					p.Log.Error("Error rollback all db changes:", slog.String("err", err.Error()))
				}
				p.storage.Close()
				panic(err)
			}
		case "offer":
			var temp xmltypes.OfferType
			xmlStruct := temp.КоммерческаяИнформация
			decoder := xml.NewDecoder(file)
			if err := decoder.Decode(&xmlStruct); err != nil {
				p.Log.Error("Error rollback all db changes:", slog.String("err", err.Error()))
				panic(err)
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			err := p.OfferParser(&xmltypes.OfferType{КоммерческаяИнформация: xmlStruct}, fileItem.PathFile, movedFiles.NewPath)
			if err != nil {
				p.Log.Error("Error offer, now rollback all db changes:", slog.String("err", err.Error()))
				err = pgxTransaction.Rollback(p.storage.Ctx)
				if err != nil {
					p.Log.Error("Error rollback all db changes:", slog.String("err", err.Error()))
				}
				p.storage.Close()
				panic(err)
			}
		}
	}
	err = pgxTransaction.Commit(ctx)
	if err != nil {
		p.Log.Error("Error commit all db changes:", slog.String("err", err.Error()))
	} else {
		p.Log.Info("DB changes committed")
	}
	p.storage.Close()
	p.Log.Debug("DB shutdown")
	p.Log.Info("==== Parser success finished")
}
