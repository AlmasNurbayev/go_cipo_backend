package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// InitLogger initializes a logger based on the environment.
//
// It takes a string parameter 'env' and returns a pointer to slog.Logger.
func InitLogger(env string, logFile *os.File) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(tint.NewHandler(os.Stdout, &tint.Options{
			Level: slog.LevelDebug,
		}))
	case envDev:
		log = slog.New(tint.NewHandler(os.Stdout, &tint.Options{
			Level: slog.LevelDebug,
		}))
	case envProd:
		// если передан файл лога - создаем и используем файл
		if logFile != nil {
			logFile, err := os.OpenFile("../assets/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0775)
			if err != nil {
				panic(err)
			}
			log = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
		} else {
			// иначе пишем в консоль ОС
			log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
		}
	}

	return log
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
