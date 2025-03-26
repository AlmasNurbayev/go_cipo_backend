package middleware

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	//"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddlewareFunc(httpRequestCounter *prometheus.CounterVec, httpRequestDuration *prometheus.HistogramVec) fiber.Handler {
	return func(c fiber.Ctx) error {

		start := time.Now()

		// Выполняем следующий обработчик
		err := c.Next()

		// Засекаем время выполнения
		duration := time.Since(start).Seconds()

		// Записываем метрики
		httpRequestCounter.WithLabelValues(c.Method(), strconv.Itoa(c.Response().StatusCode()), c.Route().Path, c.OriginalURL()).Inc()
		httpRequestDuration.WithLabelValues(c.Method(), c.Route().Path, strconv.Itoa(c.Response().StatusCode()), c.OriginalURL()).Observe(duration)

		//fmt.Println("Duration: ", duration)

		return err
	}

}
