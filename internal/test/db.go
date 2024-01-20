package test

import (
	"context"
	"fmt"
	"terrapak/internal/config"
	"terrapak/internal/config/mid"
	"terrapak/internal/db"
	"terrapak/internal/db/entity"
	"terrapak/internal/db/services"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func PostgresDB(config *config.Config, ctx context.Context) *postgres.PostgresContainer {

	postgresContainer, err  := postgres.RunContainer(ctx,
								testcontainers.WithImage("postgres:16"),
								postgres.WithUsername(config.Database.Username),
								postgres.WithPassword(config.Database.Password),
								testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections")))
	if err != nil {
		panic(err)
	}
	fmt.Println("starting postgres container")
	postgresContainer.Start(ctx)
	fmt.Println("postgres container started")
	db.Start()
	fmt.Println("database started")
	return postgresContainer
	
}

func SeedModule(mid ...mid.MID){
	ms := services.ModulesService{}
	for _, m := range mid {
		module := entity.Module{
			Name: m.Name,
			Namespace: m.Namespace,
			Provider: m.Provider,
			Version: m.Version,
			DownloadCount: 0,
			Readme: "test readme",
			SHAChecksum: "123456789",
		}
		ms.Create(module)
	}
}

func SeedUser(user ...entity.User){
	us := services.UserService{}
	for _, u := range user {
		us.Create(u)
	}
}

func SeedOrganization(org ...entity.Organization){
	os := services.OrganizationService{}
	for _, o := range org {
		os.Create(o)
	}
}