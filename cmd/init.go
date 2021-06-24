package main

import (
	"fmt"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func Init(options *terraform.Options) string {
	out, err := InitE(options)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return out
}

func InitE(options *terraform.Options) (string, error) {
	args := []string{"init", fmt.Sprintf("-upgrade=%t", options.Upgrade)}
	args = append(args, terraform.FormatTerraformBackendConfigAsArgs(options.BackendConfig)...)
	args = append(args, terraform.FormatTerraformPluginDirAsArgs(options.PluginDir)...)

	return RunTerraformCommandE(options, args...)
}
