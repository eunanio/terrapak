package router

import (
	"terrapak/internal/api/discovery"

	"github.com/gin-gonic/gin"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)
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
		"message": "pong",
	})
}


func ModuleRoutes(r *gin.RouterGroup) {
	r.GET("/:namespace/:name/:provider/:version/download",nil)
	r.GET("/:namespace/:name/:provider/versions",nil)
}

func ApiRoutes(r *gin.RouterGroup) {
	r.POST("/:namespace/:name/:provider/:version/upload",nil)
	r.GET("/:namespace/:name/:provider/:version",nil)
	r.GET("/:namespace/:name/:provider/:version/close",nil)
	r.GET("/:namespace/:name/:provider/:version/publish",nil)
}
	