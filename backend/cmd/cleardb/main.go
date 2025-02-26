package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/logger"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/storage/postgres"
)

func main() {
	op := "cleardb.main"

	// считываем период из параметров командной строки
	var period string
	flag.StringVar(&period, "period", "5d", "period before clear")
	flag.Parse()

	Cfg := config.MustLoad()
	Logger := logger.InitLogger(Cfg.Env)
	Logger.With("op", op)
	Logger.Info("init cleardb on env: " + Cfg.Env)

	pastTime, err := parsePeriod(period)
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(1)
	}
	Logger.Info("period before clear: " + pastTime.String())

	// подключаем БД без сервиса и DI, т.к. не требуются слои абстракции
	Storage, err := postgres.NewStorage(Cfg.Dsn, Logger, Cfg.HTTP.HTTP_WRITE_TIMEOUT)
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.HTTP.HTTP_WRITE_TIMEOUT)

	pgxTransaction, err := Storage.Db.Begin(ctx)
	if err != nil {
		Logger.Error("Not created transaction:", slog.String("err", err.Error()))
		Storage.Close()
		defer cancel()
		os.Exit(1)
	}
	Storage.Tx = &pgxTransaction

	// получаем масссив элементов для удаления
	list, err := Storage.GetQntPriceRegistryBeforeDate(ctx, pastTime)
	if err != nil {
		// если ошибка, то откатываем транзакцию
		Logger.Error(err.Error())
		pgxTransaction.Rollback(ctx)
		Storage.Close()
		os.Exit(1)
	}
	Logger.Info("count items to delete: " + strconv.Itoa(len(list)))

	if len(list) == 0 {
		// если ничего нет для удаления, то откатываем транзакцию
		Logger.Info("No items to delete")
		pgxTransaction.Rollback(ctx)
		Storage.Close()
		defer cancel()
		return
	} else {
		Logger.Info("list first id: " + strconv.Itoa(int(list[0].Id)))
		Logger.Info("list last id: " + strconv.Itoa(int(list[len(list)-1].Id)))
	}

	// преобразуем масссив в формат []int64
	var ids []int64
	for _, item := range list {
		ids = append(ids, item.Id)
	}

	// удаляем записи регистра с полученными id
	err = Storage.DeleteQntPriceRegistryById(ctx, ids)
	if err != nil {
		// если ошибка, то откатываем транзакцию
		Logger.Error(err.Error())
		pgxTransaction.Rollback(ctx)
		Storage.Close()
		defer cancel()
		os.Exit(1)
	}

	// коммитим транзакцию
	err = pgxTransaction.Commit(ctx)
	if err != nil {
		Logger.Error("Error commit all db changes:", slog.String("err", err.Error()))
	} else {
		Logger.Info("DB changes committed, DONE")
	}

	Storage.Close()
	defer cancel()

}

// получаем время в прошлом на заданное кол-во времени назад
func parsePeriod(period string) (time.Time, error) {
	multipliers := map[string]time.Duration{
		"s": time.Second,
		"m": time.Minute,
		"h": time.Hour,
		"d": 24 * time.Hour,
	}

	if len(period) < 2 {
		return time.Time{}, fmt.Errorf("invalid period format: %s", period)
	}

	// Разделяем число и суффикс
	valueStr := period[:len(period)-1]
	unit := period[len(period)-1:]

	multiplier, exists := multipliers[unit]
	if !exists {
		return time.Time{}, fmt.Errorf("unknown time unit: %s", unit)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid number: %s", valueStr)
	}

	duration := time.Duration(value) * multiplier
	pastTime := time.Now().Add(-duration)
	//fmt.Printf("Time after %s: %s\n", period, futureTim

	return pastTime, nil
}
