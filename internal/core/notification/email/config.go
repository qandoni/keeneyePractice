package core_email

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string `envconfig:"HOST" required:"true"`
	Port     int    `envconfig:"PORT" required:"true"`
	Username string `envconfig:"USERNAME" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	From     string `envconfig:"FROM" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("SMTP", &config); err != nil {
		return Config{}, fmt.Errorf("proccess envconfig: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get email sender config: %w", err)
		panic(err)
	}
	return config
}
