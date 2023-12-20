package client

import (
	"fmt"
	"terrapak/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBResult struct {
	ID      int    `json:"id"`
	Datname string `json:"datname"`
}

const DB_NAME = "terrapak"

var defaultDbClient *gorm.DB

func NewDBClient() *gorm.DB {
	fmt.Println("[SETUP] - connecting to database")
	gc := config.GetDefault()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", gc.Database.Username, gc.Database.Password, gc.Database.Hostname, DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	setDefaultDBClient(db)
	return db
}

func NewDBServerClient() *gorm.DB {
	fmt.Println("[SETUP] - connecting to db server")
	gc := config.GetDefault()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432", gc.Database.Username, gc.Database.Password, gc.Database.Hostname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func setDefaultDBClient(client *gorm.DB) {
	defaultDbClient = client
}

func Default() *gorm.DB {
	return defaultDbClient
}

func CreateDBIfNotExists() {
	client := NewDBServerClient()
	var results []DBResult
	client.Raw(fmt.Sprintf("select * from pg_database where datname = '%s'", DB_NAME)).Scan(&results)

	if len(results) == 0 {
		fmt.Println("[SETUP] - DB missing.. creating")
		client.Exec(fmt.Sprintf("CREATE DATABASE %s", DB_NAME))
		client.Raw(fmt.Sprintf("select * from pg_database where datname = '%s'", DB_NAME)).Scan(&results)
		if len(results) > 0 {
			fmt.Println("[SETUP] - database created")
		} else {
			panic("ERROR: Cannot create database")

		}
	}
}