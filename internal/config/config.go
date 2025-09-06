package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Env               string `env:"ENV"`
	Dsn               string `env:"DSN" json:"-"`
	GOMAXPROCS        int    `env:"GOMAXPROCS"`
	POSTGRES_USER     string `env:"POSTGRES_USER" json:"-"`
	POSTGRES_PASSWORD string `env:"POSTGRES_PASSWORD" json:"-"`
	POSTGRES_DB       string `env:"POSTGRES_DB"`
	POSTGRES_PORT     string `env:"POSTGRES_PORT"`
	HTTP              struct {
		HTTP_LOG_FILE          string        `env:"HTTP_LOG_FILE"`
		HTTP_PORT              string        `env:"HTTP_PORT"`
		HTTP_READ_TIMEOUT      time.Duration `env:"HTTP_READ_TIMEOUT"`
		HTTP_WRITE_TIMEOUT     time.Duration `env:"HTTP_WRITE_TIMEOUT"`
		HTTP_IDLE_TIMEOUT      time.Duration `env:"HTTP_IDLE_TIMEOUT"`
		HTTP_PREFORK           bool          `env:"HTTP_PREFORK"`
		CORS_ALLOW_ORIGINS     []string      `env:"HTTP_CORS_ALLOW_ORIGINS" envSeparator:","` // разделенные запятыми
		CORS_ALLOW_HEADERS     []string      `env:"HTTP_CORS_ALLOW_HEADERS" envSeparator:","` // разделенные запятыми
		CORS_ALLOW_CREDENTIALS bool          `env:"HTTP_CORS_ALLOW_CREDENTIALS"`
		EXCLUDE_VIDS_IN_LIST   []string      `env:"EXCLUDE_VIDS_IN_LIST" envSeparator:","` // разделенные запятыми, исключить эти виды моделей из выдачи productSearch
		EXCLUDES_SIZES_LEN_MIN int           `env:"EXCLUDES_SIZES_LEN_MIN"`                // исключить из выдачи ProductsFilter размеры с количеством символами больше
	}
	Parser struct {
		PARSER_CLASSIFICATOR_FILE string `env:"PARSER_CLASSIFICATOR_FILE"`
		PARSER_OFFER_FILE         string `env:"PARSER_OFFER_FILE"`
		PARSER_IMAGE_FOLDER       string `env:"PARSER_IMAGE_FOLDER"`
		PARSER_DEFAULT_USER_ID    int64  `env:"PARSER_DEFAULT_USER_ID"`
		PARSER_ASSETS_PATH        string `env:"PARSER_ASSETS_PATH"`
		PARSER_INPUT_PATH         string `env:"PARSER_INPUT_PATH"`
	}
	Auth struct {
		TokenTTL    time.Duration `env:"TOKEN_TTL"`
		SECRET_KEY  string        `env:"SECRET_KEY"  json:"-"`
		SECRET_BYTE []byte        `json:"-"`
	}
	Kaspi struct {
		KASPI_API_URL string `env:"KASPI_API_URL"`
	}
}

func MustLoad() *Config {

	var cfg Config

	// загружаем сначала из переменных окружения
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal("not load config: ", err)
	}

	// если в конфиге нет Env или Dsn, то переменные окружения не заданы
	// и ищем файл ./config/.env для получения DSN
	if (cfg.Env == "") || (cfg.Dsn == "") {
		err := godotenv.Load("./config/.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		err = env.Parse(&cfg)
		if err != nil {
			log.Fatal("not load config: ", err)
		}
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		host := os.Getenv("POSTGRES_HOST")
		db := os.Getenv("POSTGRES_DB")
		port := os.Getenv("POSTGRES_PORT")
		cfg.Dsn = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + db + "?sslmode=disable"
	}

	if cfg.Auth.SECRET_KEY != "" {
		cfg.Auth.SECRET_BYTE = utils.DeriveKeyFromSecret(cfg.Auth.SECRET_KEY)
	}

	if cfg.Env == "" {
		log.Fatal("not load config: env is empty")
	} else if cfg.Dsn == "" {
		log.Fatal("not load config: Dsn is empty")
	}

	// убрать кавычки и обратные слэши
	for i, origin := range cfg.HTTP.CORS_ALLOW_ORIGINS {
		cfg.HTTP.CORS_ALLOW_ORIGINS[i] = strings.ReplaceAll(origin, "\"", "")                         // удаляем кавычки
		cfg.HTTP.CORS_ALLOW_ORIGINS[i] = strings.ReplaceAll(cfg.HTTP.CORS_ALLOW_ORIGINS[i], "\\", "") // удаляем обратные слэши
	}

	for i, origin := range cfg.HTTP.CORS_ALLOW_HEADERS {
		cfg.HTTP.CORS_ALLOW_HEADERS[i] = strings.ReplaceAll(origin, "\"", "")                         // удаляем кавычки
		cfg.HTTP.CORS_ALLOW_HEADERS[i] = strings.ReplaceAll(cfg.HTTP.CORS_ALLOW_HEADERS[i], "\\", "") // удаляем обратные слэши
	}

	return &cfg
}
