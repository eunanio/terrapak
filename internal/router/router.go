package router

import (
	"fmt"
	"terrapak/internal/api/auth"
	"terrapak/internal/api/auth/roles"
	"terrapak/internal/api/discovery"
	"terrapak/internal/api/middleware"
	"terrapak/internal/api/modules"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	
	r.GET("/ping", Ping)
	moduleRouter := r.Group("v1/modules")
	ApiRouter 	 := r.Group("v1/api")
	authRouter 	 := r.Group("v1/auth")
	ModuleRoutes(moduleRouter)
	ApiRoutes(ApiRouter)
	AuthRoutes(authRouter)

	//Service Discovery
	r.GET("/.well-known/terraform.json", discovery.Serve)
	fmt.Println("[SETUP] - Terrapak Started")
	r.Run(":5551")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func ModuleRoutes(r *gin.RouterGroup) {
	r.Use(middleware.HasAuthenticatedRole(roles.Editor))
	r.GET("/:namespace/:name/:provider/:version/download",modules.Download)
	r.GET("/:namespace/:name/:provider/versions",modules.Versions)
}

func ApiRoutes(r *gin.RouterGroup) {
	r.Use(middleware.HasAuthenticatedRole(roles.Editor))
	r.GET("/:namespace/:name/:provider/:version",modules.Read)
	r.POST("/:namespace/:name/:provider/:version/upload",modules.Upload)
	r.GET("/:namespace/:name/:provider/:version/close",modules.RemoveDraft)
	r.GET("/:namespace/:name/:provider/:version/publish",modules.PublishDraft)
}

func AuthRoutes(r *gin.RouterGroup) {
	r.GET("/authorize", auth.Authorize)
	r.POST("/token", auth.Token)
	r.GET("/callback", auth.Callback)
}