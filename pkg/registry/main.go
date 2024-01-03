package registry

import (
	"fmt"
	"terrapak/internal/config"
	"terrapak/internal/db"
	"terrapak/internal/router"

	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load(); if err != nil {
		fmt.Println("Error loading .env file")
	}

	config.Load()
	db.Start()
	router.Start()
}