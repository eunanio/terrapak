package db

import (
	"fmt"
	"terrapak/internal/db/client"
	"terrapak/internal/db/entity"

	"gorm.io/gorm"
)

type Table interface {
	TableName() string
	Up(client *gorm.DB)
}

func Start() {
	client.CreateDBIfNotExists()
	//gc := config.GetDefault()
	client := client.NewDBClient()
	// Migrate tables on startup
	tables := []Table{
		&entity.User{},
		&entity.Module{},
		&entity.ApiKeys{},
		&entity.Organization{},
	}

	for _, table := range tables {
		fmt.Printf("[SETUP] - migrating %s table\n", table.TableName())
		table.Up(client)
	}

	// if gc.AuthProvider.Type == "pat" || gc.AuthProvider.Type == "PAT" {
	// 	pat.CreateDefaultPAT(client)
	// }
	

}