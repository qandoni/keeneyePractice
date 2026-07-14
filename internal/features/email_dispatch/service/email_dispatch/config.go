package email_dispatch_service

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	CompleteURL       string `envconfig:"COMPLETE_URL" required:"true"`
	EmailSubject      string `envconfig:"EMAIL_SUBJECT" required:"true"`
	EmailBody         string `envconfig:"EMAIL_BODY" required:"true"`
	EmailDispatchMode string `envconfig:"EMAIL_DISPATCH_MODE" required:"true"`
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process(
		"REGISTRATION",
		&cfg,
	); err != nil {
		return Config{}, fmt.Errorf(
			"process registration config: %w",
			err,
		)
	}

	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
