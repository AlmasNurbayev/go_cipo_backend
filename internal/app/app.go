package app

import (
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/app/httpApp"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/storage/postgres"
	"github.com/gofiber/fiber/v3"
)

type App struct {
	Log             *slog.Logger
	httpApp         *httpApp.HttpApp
	PostgresStorage *postgres.Storage
	Cfg             *config.Config
}

func NewApp(cfg *config.Config, log *slog.Logger) *App {
	storage, err := postgres.NewStorage(cfg.Dsn, log, cfg.HTTP.HTTP_WRITE_TIMEOUT)
	if err != nil {
		log.Error("not init postgres storage")
		panic(err)
	}
	httpApp := httpApp.NewApp(log, cfg, storage)
	return &App{
		Log:             log,
		httpApp:         httpApp,
		PostgresStorage: storage,
		Cfg:             cfg,
	}
}

func (a *App) Run() {
	err := a.httpApp.Server.Listen(":"+a.Cfg.HTTP.HTTP_PORT, fiber.ListenConfig{
		EnablePrefork:   a.Cfg.HTTP.HTTP_PREFORK,
		ShutdownTimeout: a.Cfg.HTTP.HTTP_WRITE_TIMEOUT + a.Cfg.HTTP.HTTP_IDLE_TIMEOUT,
	})
	if err != nil {
		a.Log.Error("not start server: ", slog.String("err", err.Error()))
		panic(err)
	}
}

func (a *App) Stop() {
	err := a.httpApp.Server.Shutdown()
	a.PostgresStorage.Close()
	if err != nil {
		a.Log.Error("error on stop server: ", slog.String("err", err.Error()))
		panic(err)
	}
}
