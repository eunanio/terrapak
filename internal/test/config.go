package test

import (
	"os"
	"terrapak/internal/config"
)

func SetupTestConfig() {
	os.Setenv(config.ENV_DB_USER, "test")
	os.Setenv(config.ENV_DB_PASS, "test")
	os.Setenv(config.ENV_DB_HOST, "127.0.0.1")
	os.Setenv(config.ENV_TP_HOST, "localhost:5551")
	os.Setenv(config.ENV_TP_AUTH_TYPE, "noop")
	config.Load()
}

func CleanupEnv() {
	os.Unsetenv(config.ENV_DB_USER)
	os.Unsetenv(config.ENV_DB_PASS)
	os.Unsetenv(config.ENV_DB_HOST)
	os.Unsetenv(config.ENV_TP_HOST)
	os.Unsetenv(config.ENV_TP_AUTH_TYPE)
}