package registry

import (
	"terrapak/internal/config"
	"terrapak/internal/db"
	"terrapak/internal/logger"
	"terrapak/internal/router"
)

func Start() {
	lf := logger.NewLogger()
	defer lf.Close()

	config.Load()
	db.Start()
	router.Start()
}