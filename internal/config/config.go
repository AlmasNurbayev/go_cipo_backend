package config

import (
	"log"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Env               string        `env:"ENV"`
	Dsn               string        `env:"DSN" json:"-"`
	GOMAXPROCS        int           `env:"GOMAXPROCS"`
	POSTGRES_USER     string        `env:"POSTGRES_USER" json:"-"`
	POSTGRES_PASSWORD string        `env:"POSTGRES_PASSWORD" json:"-"`
	POSTGRES_DB       string        `env:"POSTGRES_DB"`
	POSTGRES_PORT     string        `env:"POSTGRES_PORT"`
	TokenTTL          time.Duration `env:"TOKEN_TTL"`
	HTTP              struct {
		HTTP_PORT          string        `env:"HTTP_PORT"`
		HTTP_READ_TIMEOUT  time.Duration `env:"HTTP_READ_TIMEOUT"`
		HTTP_WRITE_TIMEOUT time.Duration `env:"HTTP_WRITE_TIMEOUT"`
		HTTP_IDLE_TIMEOUT  time.Duration `env:"HTTP_IDLE_TIMEOUT"`
		HTTP_PREFORK       bool          `env:"HTTP_PREFORK"`
	}
	Parser struct {
		PARSER_CLASSIFICATOR_FILE string `env:"PARSER_CLASSIFICATOR_FILE"`
		PARSER_OFFER_FILE         string `env:"PARSER_OFFER_FILE"`
		PARSER_IMAGE_FOLDER       string `env:"PARSER_IMAGE_FOLDER"`
		PARSER_DEFAULT_USER_ID    int64  `env:"PARSER_DEFAULT_USER_ID"`
		PARSER_ASSETS_PATH        string `env:"PARSER_ASSETS_PATH"`
		PARSER_INPUT_PATH         string `env:"PARSER_INPUT_PATH"`
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

	if cfg.Env == "" {
		log.Fatal("not load config: env is empty")
	} else if cfg.Dsn == "" {
		log.Fatal("not load config: Dsn is empty")
	}

	return &cfg
}
