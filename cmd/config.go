package main

import (
	"github.com/caarlos0/env"
)

// ENV TF_PLUGIN_CACHE_DIR="/var/lib/terraform/providers"
// ENV TF_PLUGIN_DIR="/var/lib/terraform/providers"
// ENV TF_MODULE_CACHE_DIR="/var/lib/terraform/modules"
// ENV TF_IN_AUTOMATION=true
// ENV CHECKPOINT_DISABLE=true

type config struct {
	Port           string `env:"TF_RUNNER_API_PORT" envDefault:"8080"`
	WorkDir        string `env:"TF_RUNNER_WORKDIR" envDefault:"/var/workspace"`
	PluginCacheDir string `env:"TF_PLUGIN_CACHE_DIR" envDefault:"/var/lib/terraform/providers"`
	TFBinaryPath   string `env:"TF_BINARY_PATH" envDefault:"/usr/local/bin/terraform"`
}

func (c *config) loadConfig() {
	if err := env.Parse(c); err != nil {
		recordErrorMetric()
		logger.Fatal("Unable to parse ENV vars")
	}
}
