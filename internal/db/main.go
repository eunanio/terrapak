package db

import (
	"fmt"
	"log/slog"
	"terrapak/internal/db/client"
	"terrapak/internal/db/entity"
	"terrapak/internal/db/jobs"

	"gorm.io/gorm"
)

type Table interface {
	TableName() string
	Up(client *gorm.DB)
}

func Start() {
	jobs.CreateDBIfNotExists()
	db_client := client.NewDBClient()

	// Migrate tables on startup
	tables := []Table{
		&entity.User{},
		&entity.Module{},
		&entity.ApiKeys{},
		&entity.Organization{},
	}

	for _, table := range tables {
		slog.Info(fmt.Sprintf("Migrating %s table", table.TableName()))
		table.Up(db_client)
	}

	jobs.CreateDefaultOrganizationIfNotExists()
}