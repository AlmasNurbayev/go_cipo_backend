package parserJSON

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/logger"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parserJSON/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/storage/postgres"
)

type ParserJSON struct {
	Version string
	Cfg     *config.Config
	storage *postgres.Storage
	Log     *slog.Logger
	//service *parserService.ParserService
}

func New(version string) *ParserJSON {
	return &ParserJSON{
		Version: version,
	}
}

func (p *ParserJSON) Init() {
	op := "parserJSON.Init"

	p.Version = "v0.1.0"
	p.Cfg = config.MustLoad()
	p.Log = logger.InitLogger(p.Cfg.Env, nil)
	p.Log.With("op", op)
	p.Log.Info("==== init parserJSON on env: " + p.Cfg.Env)
	p.Log.Debug("debug message is enabled")
	postgresStorage, err := postgres.NewStorage(p.Cfg.Dsn, p.Log, p.Cfg.HTTP.HTTP_WRITE_TIMEOUT, p.Cfg)
	if err != nil {
		p.Log.Error(err.Error())
		panic(err)
	}
	p.storage = postgresStorage
}

func (p *ParserJSON) Run() {
	op := "parserJSON.Run"
	p.Log.With("op", op).Info("start parserJSON")

	defer p.storage.Close()

	jsonPath, err := MovedFTPFiles(p.Cfg, p.Log)
	if err != nil {
		p.Log.Error("error moved FTP files: ", slog.String("error", err.Error()))
		return
	}
	p.Log.Info("Moved JSON file successfully: ", slog.String("jsonPath", jsonPath))

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		p.Log.Error("error read file: ", slog.String("error", err.Error()))
		return
	}

	var result dto.StockJSON
	// Десериализация (Unmarshalling)
	err = json.Unmarshal(data, &result)
	if err != nil {
		p.Log.Error("error unmarshal JSON: ", slog.String("error", err.Error()))
		return
	}
	p.Log.Info("Unmarshal JSON successfully")

	ctx, cancel := context.WithTimeout(context.Background(), p.Cfg.HTTP.HTTP_WRITE_TIMEOUT)
	defer cancel()
	pgxTransaction, err := p.storage.Db.Begin(ctx)
	if err != nil {
		p.Log.Error("Not created transaction:", slog.String("err", err.Error()))
		p.Log.Info("==== Parser ERROR finished")
		os.Exit(1)
	}
	p.storage.Tx = &pgxTransaction
	var is_commited bool = false

	defer func() {
		if pgxTransaction != nil && !is_commited {
			_ = pgxTransaction.Rollback(context.Background())
			p.Log.Warn("Try Rollback transaction")
		}
	}()

	registrator_id, err := partParsers.ParserRegistrator(p.Cfg, p.Log, p.storage, result, jsonPath)
	if err != nil {
		p.Log.Error("error parser registrator: ", slog.String("error", err.Error()))
		return
	}
	err = partParsers.ParserProductGroups(p.Log, ctx, p.storage, result, registrator_id)
	if err != nil {
		p.Log.Error("error parser product groups: ", slog.String("error", err.Error()))
		return
	}
	err = partParsers.ParserProductVids(p.Log, ctx, p.storage, result, registrator_id)
	if err != nil {
		p.Log.Error("error parser product vids: ", slog.String("error", err.Error()))
		return
	}
	err = partParsers.ParserVidModeli(p.Log, ctx, p.storage, result, registrator_id)
	if err != nil {
		p.Log.Error("error parser vid modeli: ", slog.String("error", err.Error()))
		return
	}

	err = partParsers.ParserStore(p.Log, ctx, p.storage, result, registrator_id)
	if err != nil {
		p.Log.Error("error parser store: ", slog.String("error", err.Error()))
		return
	}

	err = partParsers.ParserSize(p.Log, ctx, p.storage, result, registrator_id)
	if err != nil {
		p.Log.Error("error parser size: ", slog.String("error", err.Error()))
		return
	}

	err = partParsers.ParserBrend(p.Log, ctx, p.storage, result, registrator_id)
	if err != nil {
		p.Log.Error("error parser brend: ", slog.String("error", err.Error()))
		return
	}

	err = partParsers.ParserProduct(p.Log, ctx, p.storage, result, registrator_id)
	if err != nil {
		p.Log.Error("error parser product: ", slog.String("error", err.Error()))
		return
	}

	assetsPath := "assets/product_images"
	if p.Cfg.Parser.PARSER_ASSETS_PATH != "" {
		assetsPath = p.Cfg.Parser.PARSER_ASSETS_PATH + "/product_images"
	}
	err = partParsers.ImageRegistryParser(ctx, p.storage, p.Log, result, registrator_id, assetsPath)
	if err != nil {
		p.Log.Error("error parser image registry: ", slog.String("error", err.Error()))
		return
	}

	err = partParsers.ParserQnt(p.Log, ctx, p.storage, result, registrator_id)
	if err != nil {
		p.Log.Error("error parser qnt: ", slog.String("error", err.Error()))
		return
	}

	// Commit transaction
	err = pgxTransaction.Commit(context.Background())
	if err != nil {
		p.Log.Error("error commit transaction: ", slog.String("error", err.Error()))
		return
	}
	is_commited = true
	p.Log.Info("Commit transaction done")

	p.Log.Info("parserJSON success finished", slog.String("registrator_id", strconv.FormatInt(registrator_id, 10)))
}
