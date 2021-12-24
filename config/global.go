package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Environment struct {
	Port     int    `env:"PORT"`
	DB_URL   string `env:"DB_URL"`
	Database string `env:"DB_DATABASE"`
	APP_ID_1 string `env:"APP_ID_1"`
}

var ENV Environment

func Start() {
	_, err := env.UnmarshalFromEnviron(&ENV)
	if err != nil {
		log.Fatal(err)
	}
}
