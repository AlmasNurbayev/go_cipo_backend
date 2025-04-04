package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/app"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/logger"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
)

func main() {
	// ключевые сообщения дублируем и в консоль и в логгер (он может писать в файл)
	fmt.Println("============ start main ============")
	cfg := config.MustLoad()

	var logFile *os.File
	if cfg.Env == "prod" {
		var err error
		fmt.Println("append to log file: " + cfg.HTTP.HTTP_LOG_FILE)
		logFile, err = os.OpenFile(cfg.HTTP.HTTP_LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0775)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := logFile.Close(); err != nil {
				fmt.Println("Error closing file:", err)
			}
		}()
	} else {
		logFile = nil
	}

	Log := logger.InitLogger(cfg.Env, logFile)
	Log.Info("============ start main ============")

	cfgBytes, err := utils.PrintAsJSON(cfg)
	if err != nil {
		panic(err)
	}

	Log.Info("load config: ")
	fmt.Println(string(*cfgBytes))

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
	fmt.Println("============ server stopped ============")
}
