package jobs

import (
	"fmt"
	"os"
	"terrapak/internal/api/auth/roles"
	"terrapak/internal/config"
	"terrapak/internal/db/client"
	"terrapak/internal/db/entity"
)

func CreateDBIfNotExists() {
	db_client := client.NewDBServerClient()
	var results []client.DBResult
	db_client.Raw(fmt.Sprintf("select * from pg_database where datname = '%s'", client.DB_NAME)).Scan(&results)

	if len(results) == 0 {
		fmt.Println("[SETUP] - DB missing.. creating")
		db_client.Exec(fmt.Sprintf("CREATE DATABASE %s", client.DB_NAME))
		db_client.Raw(fmt.Sprintf("select * from pg_database where datname = '%s'", client.DB_NAME)).Scan(&results)
		if len(results) > 0 {
			fmt.Println("[SETUP] - database created")
		} else {
			panic("ERROR: Cannot create database")

		}
	}
}

func CreateAdminUserIfNotExists() {
	org := entity.Organization{}
	var count int64
	userEnv, ext := os.LookupEnv(config.ENV_TP_USER)
	if !ext {
		panic("TP_USER var not set")
		
	}

	passwordEnv, ext := os.LookupEnv(config.ENV_TP_PASSWORD)
	if !ext {
		panic("TP_PASSWORD var not set")
	}

	client := client.NewDBClient()
	gc := config.GetDefault()
	client.Model(&entity.User{}).Where("name = ?", userEnv).Count(&count)
	client.Raw("SELECT * FROM organizations WHERE name = ?", gc.Organization).Scan(&org)
	fmt.Println(count)
	if count == 0 {
		user := entity.User{}
		user.Name = userEnv
		user.Email = "admin@terrapak.local"
		user.PasswordHash = config.HashSecret(passwordEnv)
		user.OrganizationID = org.ID
		user.Role = roles.Owner
		user.Create(client)
	}

}

func CreateDefaultOrganizationIfNotExists() {
	client := client.NewDBClient()
	gc := config.GetDefault()
	var count int64
	client.Table("organizations").Count(&count)

	if count == 0 {
		org := entity.Organization{}
		org.Name = gc.Organization
		org.Create(client)
	}
}