package clients

import (
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"path"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
)

type KaspiRequestParameters struct {
	method  string
	cfg     *config.Config
	log     *slog.Logger
	token   string
	query   url.Values
	urlMain string
	urlPart string
	body    string
}

// по нужной категории товара получаем список атрибутов
func KaspiGetAttributes(cfg *config.Config, log *slog.Logger, c string, token string) (int, string, error) {
	op := "KaspiGetAttributes"
	log = log.With(slog.String("op", op))

	query := url.Values{}
	query.Add("c", c)

	statusCode, body, err := apiKaspiSender(KaspiRequestParameters{
		method:  "GET",
		cfg:     cfg,
		log:     log,
		token:   token,
		query:   query,
		urlMain: cfg.Kaspi.KASPI_API_URL,
		urlPart: "products/classification/attributes",
	})
	if err != nil {
		return 0, "", err
	}

	return statusCode, body, nil
}

// по имени атрибута получаем список возможных значений
func KaspiGetValues(cfg *config.Config, log *slog.Logger, a string, c string, token string) (int, string, error) {
	op := "KaspiGetValues"
	log = log.With(slog.String("op", op))

	query := url.Values{}
	query.Add("a", a)
	query.Add("c", c)

	statusCode, body, err := apiKaspiSender(KaspiRequestParameters{
		method:  "GET",
		cfg:     cfg,
		log:     log,
		token:   token,
		query:   query,
		urlMain: cfg.Kaspi.KASPI_API_URL,
		urlPart: "products/classification/attribute/values",
	})
	if err != nil {
		return 0, "", err
	}

	return statusCode, body, nil
}

func KaspiGetCategories(cfg *config.Config, log *slog.Logger, token string) (int, string, error) {
	op := "KaspiGetCategories"
	log = log.With(slog.String("op", op))

	statusCode, body, err := apiKaspiSender(KaspiRequestParameters{
		method:  "GET",
		cfg:     cfg,
		log:     log,
		token:   token,
		urlMain: cfg.Kaspi.KASPI_API_URL,
		urlPart: "products/classification/categories",
	})
	if err != nil {
		log.Error(err.Error())
		return 0, "", err
	}

	return statusCode, body, nil
}

// Отправление запроса в Kaspi API по заданному URL, query и токену
func apiKaspiSender(params KaspiRequestParameters) (int, string, error) {
	base, err := url.Parse(params.urlMain)
	if err != nil {
		params.log.Error("Api error:", slog.String("err", err.Error()))
		return 0, "", err
	}
	base.Path = path.Join(base.Path, params.urlPart)

	if params.query != nil {
		base.RawQuery = params.query.Encode()
	}

	client := &http.Client{}
	req, err := http.NewRequest(params.method, base.String(), nil)
	if err != nil {
		params.log.Error("Api error:", slog.String("err", err.Error()))
		return 0, "", err
	}
	params.log.Debug("request", params.method, base.String())

	if params.token != "" {
		req.Header.Set("X-Auth-Token", params.token)
	}
	req.Header.Set("Content-Type", "application/json")
	//fmt.Println("X-Auth-Token", params.token)

	resp, err := client.Do(req)
	if err != nil {
		params.log.Error("Api error:", slog.String("err", err.Error()))
		return 0, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		params.log.Error("Api error:", slog.String("err", err.Error()))
		return 0, "", err
	}

	return resp.StatusCode, string(body), nil
}
