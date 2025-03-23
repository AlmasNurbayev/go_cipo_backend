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
	fmt.Println("============ start main ============")
	cfg := config.MustLoad()
	Log := logger.InitLogger(cfg.Env)
	p, err := utils.PrintAsJSON(cfg)
	if err != nil {
		panic(err)
	}
	Log.Info("load config: ")
	Log.Info(string(*p))
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

	app.Stop()

	Log.Info("server stopped")
}
