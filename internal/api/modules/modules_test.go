package modules

import (
	"context"
	"terrapak/internal/config"
	"terrapak/internal/config/mid"
	"terrapak/internal/test"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestModuleRead(t *testing.T) {
	db_config := config.DatabaseConfig{
		Hostname:     "127.0,0,1",
		Username:    "test",
		Password:    "test",
	}

	config := config.Config{
		Database: db_config,
	}

	ctx := gin.Context{}
	bctx := context.Background()
	mid := mid.NewMID("myorg", "test-mod", "test", "0.0.1")
	ctx.Params = append(ctx.Params, gin.Param{Key: "namespace", Value: mid.Namespace})
	ctx.Params = append(ctx.Params, gin.Param{Key: "provider", Value: mid.Provider})
	ctx.Params = append(ctx.Params, gin.Param{Key: "name", Value: mid.Name})
	ctx.Params = append(ctx.Params, gin.Param{Key: "version", Value: mid.Version})
	test.PostgresDB(&config,bctx)
}