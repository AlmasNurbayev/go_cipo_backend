package httphandlers

import (
	"bytes"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/common/expfmt"
)

func (h *Handler) GetMetrics(c fiber.Ctx) error {

	op := "HttpHandlers.GetMetrics"
	log := h.log.With(slog.String("op", op))

	// Собираем метрики из глобального реестра
	mfs, err := h.promRegistry.Gather()
	if err != nil {
		log.Error("Error gathering metrics:", slog.String("err", err.Error()[:20]))
		return c.Status(500).SendString("Error gathering metrics")
	}

	// Кодируем метрики
	var buf bytes.Buffer
	enc := expfmt.NewEncoder(&buf, expfmt.NewFormat(expfmt.TypeTextPlain))
	for _, mf := range mfs {
		if err := enc.Encode(mf); err != nil {
			log.Error("Error encoding metrics:", slog.String("err", err.Error()))
			return c.Status(500).SendString("Error encoding metrics")
		}
	}

	return c.Status(200).Send(buf.Bytes())
}
