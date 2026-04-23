package parserJSON

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/logger"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/storage/postgres"
	"github.com/kr/pretty"
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

	jsonPath, err := MovedFTPFiles(p.Cfg, p.Log)
	if err != nil {
		p.Log.Error("error moved FTP files: ", slog.String("error", err.Error()))
		return
	}
	p.Log.Info("Moved FTP files successfully: ", slog.String("jsonPath", jsonPath))

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

	err = parserClass(result)
	if err != nil {
		p.Log.Error("error parser class: ", slog.String("error", err.Error()))
		return
	}

	pretty.Log("unmarshal result: ", result)
}
