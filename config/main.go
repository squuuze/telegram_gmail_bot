package config

import (
	"encoding/json"
	"sync"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var (
	once sync.Once
	cfg  *Config
)

type Config struct {
	TelegramApiToken string `env:"TELEGRAM_API_TOKEN"`
	ClientID         string `env:"CLIENT_ID"`
	ClientSecret     string `env:"CLIENT_SECRET"`
}

func (c *Config) String() string {
	b, _ := json.MarshalIndent(c, "", "	")
	return string(b)
}

func Get() *Config {
	_ = godotenv.Load()

	once.Do(func() {
		config := Config{}
		if err := env.Parse(&config); err != nil {
			panic(err)
		}

		cfg = &config
	})

	return cfg
}
