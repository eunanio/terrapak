package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv(ENV_CONFIG_FILE,"../../config.yml")
	os.Setenv(ENV_TP_AUTH_KEY_FILE,"../../private.pem")
	config := Load();
	assert.NotNil(t,config)
}