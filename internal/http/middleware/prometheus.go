package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	//"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	promauto "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Количество HTTP-запросов",
		},
		[]string{"method", "status_code", "route", "original_url"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Длительность HTTP-запросов",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route", "status_code", "original_url"},
	)
)

func PrometheusMiddleware(c fiber.Ctx) error {

	start := time.Now()

	// Выполняем следующий обработчик
	err := c.Next()

	// Засекаем время выполнения
	duration := time.Since(start).Seconds()

	// Записываем метрики
	httpRequestCounter.WithLabelValues(c.Method(), strconv.Itoa(c.Response().StatusCode()), c.Route().Path, c.OriginalURL()).Inc()
	httpRequestDuration.WithLabelValues(c.Method(), c.Route().Path, strconv.Itoa(c.Response().StatusCode()), c.OriginalURL()).Observe(duration)

	fmt.Println("Duration: ", duration)

	return err
}
