package router

import (
	"terrapak/internal/api/discovery"
	"terrapak/internal/api/modules"

	"github.com/gin-gonic/gin"
)

func Start() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/ping", Ping)
	moduleRouter := r.Group("v1/modules")
	ApiRouter := r.Group("v1/api")
	ModuleRoutes(moduleRouter)
	ApiRoutes(ApiRouter)

	//Service Discovery
	r.GET("/.well-known/terraform.json", discovery.Serve)
	
	r.Run(":5551")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}


func ModuleRoutes(r *gin.RouterGroup) {
	r.GET("/:namespace/:name/:provider/:version/download",modules.Download)
	r.GET("/:namespace/:name/:provider/versions",modules.Versions)
}

func ApiRoutes(r *gin.RouterGroup) {
	r.POST("/:namespace/:name/:provider/:version/upload",modules.Upload)
	r.GET("/:namespace/:name/:provider/:version",modules.Read)
	r.GET("/:namespace/:name/:provider/:version/close",modules.RemoveDraft)
	r.GET("/:namespace/:name/:provider/:version/publish",modules.PublishDraft)
}
	