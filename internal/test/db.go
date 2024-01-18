package test

import (
	"context"
	"terrapak/internal/config"
	"terrapak/internal/config/mid"
	"terrapak/internal/db"
	"terrapak/internal/db/entity"
	"terrapak/internal/db/services"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartTestDB(config *config.Config, ctx context.Context) *postgres.PostgresContainer {

	postgresContainer, err  := postgres.RunContainer(ctx,
								testcontainers.WithImage("postgres:16"),
								postgres.WithUsername(config.Database.Username),
								postgres.WithPassword(config.Database.Password),
								testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections")))
	if err != nil {
		panic(err)
	}
	postgresContainer.Start(ctx)
	db.Start()
	return postgresContainer
	
}

func SeedModule(mid mid.MID){
	ms := services.ModulesService{}
	module := entity.Module{
		Name: mid.Name,
		Namespace: mid.Namespace,
		Provider: mid.Provider,
		Version: mid.Version,
		DownloadCount: 0,
		Readme: "test readme",
		SHAChecksum: "123456789",
	}
	ms.Create(module)
}