package registry

import (
	"terrapak/internal/config"
	"terrapak/internal/db"
	"terrapak/internal/router"
)

func Start() {

	config.Load()
	db.Start()
	router.Start()
}