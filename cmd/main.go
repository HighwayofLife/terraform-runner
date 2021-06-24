package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	info = AppInfo{
		Name:        "terraform-runner",
		Description: "API to run Terraform code in Go runner microservice",
		URL:         "https://github.com/highwayoflife/terraform-runner",
	}

	cfg config

	logger *zap.SugaredLogger
)

func main() {
	logger = InitLogger()
	cfg.loadConfig()

	runnerAPI := NewServer()
	runnerAPI.Run()
}

// InitLogger iniializes zap logger
func InitLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	corelogger, _ := config.Build()
	defer corelogger.Sync() // flushes buffer, if any
	return corelogger.Sugar()
}
