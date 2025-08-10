package config

import "insider/src/infra/env"

type ConfigType struct {
	APP_PORT      string
	APP_ENV       string
	POSTGREDB_URL string
	REDIS_URL     string
	WEBHOOK_URL   string
}

func LoadConfig() *ConfigType {
	var cfg ConfigType
	env.LoadEnv(&cfg)
	return &cfg
}
