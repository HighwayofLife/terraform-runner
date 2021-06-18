package main

import (
	"github.com/caarlos0/env"
)

type config struct {
	Port string `env:"TF_RUNNER_API_PORT" envDefault:"8080"`
}

func (c *config) loadConfig() {
	if err := env.Parse(c); err != nil {
		recordErrorMetric()
		logger.Fatal("Unable to parse ENV vars")
	}
}
