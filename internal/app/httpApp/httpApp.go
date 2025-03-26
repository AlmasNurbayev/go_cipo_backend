package httpApp

import (
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	httphandlers "github.com/AlmasNurbayev/go_cipo_backend/internal/http/httpHandlers"
	httproutes "github.com/AlmasNurbayev/go_cipo_backend/internal/http/httpRoutes"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/http/middleware"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/services"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/storage/postgres"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type structValidator struct {
	validate *validator.Validate
}

type HttpApp struct {
	Log             *slog.Logger
	Server          *fiber.App
	PostgresStorage *postgres.Storage
	Cfg             *config.Config
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func NewApp(
	log *slog.Logger,
	cfg *config.Config,
	storage *postgres.Storage) *HttpApp {

	server := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
		ReadTimeout:     cfg.HTTP.HTTP_READ_TIMEOUT,
		WriteTimeout:    cfg.HTTP.HTTP_WRITE_TIMEOUT,
		IdleTimeout:     cfg.HTTP.HTTP_IDLE_TIMEOUT,
	})

	if cfg.Env != "prod" {
		server.Use(middleware.RequestTracingMiddleware(log))
	}

	service := services.NewService(log, storage, cfg)

	registry := prometheus.NewRegistry()
	httpRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_milliseconds",
			Help:    "Duration of HTTP requests in milliseconds",
			Buckets: []float64{1, 10, 50, 100, 200, 500, 1000}, // Бакеты аналогичны JS
		},
		[]string{"method", "route", "statusCode", "originalUrl"},
	)
	httpRequestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "route", "statusCode", "originalUrl"},
	)

	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		httpRequestDuration,
		httpRequestCounter,
	)

	server.Use(middleware.PrometheusMiddlewareFunc(httpRequestCounter, httpRequestDuration))

	handlers := httphandlers.NewHandler(log, service, registry)
	httproutes.RegisterRoutes(server, handlers, log)

	server.Get("/healthz", func(c fiber.Ctx) error {
		return c.Status(200).SendString("OK")
	})

	return &HttpApp{
		Log:             log,
		Server:          server,
		PostgresStorage: storage,
		Cfg:             cfg,
	}
}
