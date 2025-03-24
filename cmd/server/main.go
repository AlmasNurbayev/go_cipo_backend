package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/app"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/logger"
)

func main() {
	// ключевые сообщения дублируем и в консоль и в логгер (он может писать в файл)
	fmt.Println("============ start main ============")
	cfg := config.MustLoad()

	var logFile *os.File
	if cfg.Env == "prod" {
		var err error
		logFile, err = os.OpenFile("../assets/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0775)
		if err != nil {
			panic(err)
		}
		defer logFile.Close()
	} else {
		logFile = nil
	}

	Log := logger.InitLogger(cfg.Env, logFile)
	Log.Info("============ start main ============")
	Log.Info("load: ", slog.Any("config", cfg))
	Log.Debug("debug message is enabled")

	runtime.GOMAXPROCS(cfg.GOMAXPROCS)
	app := app.NewApp(cfg, Log)
	go func() {
		app.Run()
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	signalString := <-done
	Log.Info("received signal " + signalString.String())
	fmt.Println("received signal " + signalString.String())

	app.Stop()
	Log.Info("server stopped")
	fmt.Println("server stopped")
}
