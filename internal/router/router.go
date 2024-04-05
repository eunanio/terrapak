package router

import (
	"fmt"
	"terrapak/internal/api/auth"
	"terrapak/internal/api/auth/roles"
	"terrapak/internal/api/discovery"
	"terrapak/internal/api/middleware"
	"terrapak/internal/api/webhook"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	
	r.GET("/ping", Ping)
	moduleRouter := r.Group("v1/modules")
	ApiRouter 	 := r.Group("v1/api")
	authRouter 	 := r.Group("v1/auth")
	webhookRouter := r.Group("v1/vcs")
	ModuleRoutes(moduleRouter)
	ApiRoutes(ApiRouter)
	AuthRoutes(authRouter)
	WebhookRoutes(webhookRouter)
	//Service Discovery
	r.GET("/.well-known/terraform.json", discovery.Serve)
	fmt.Println("Terrapak Started")
	r.Run(":5551")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func ModuleRoutes(r *gin.RouterGroup) {
	endpoints := Endpoint{}
	r.Use(middleware.HasAuthenticatedRole(roles.Reader,roles.Editor))
	r.GET("/:namespace/:name/:provider/:version/download",endpoints.Download)
	r.GET("/:namespace/:name/:provider/versions",endpoints.Version)
}

func ApiRoutes(r *gin.RouterGroup) {
	endpoints := Endpoint{}
	r.Use(middleware.HasAuthenticatedRole(roles.Editor))
	r.GET("/:namespace/:name/:provider/:version",endpoints.Read)
	r.POST("/:namespace/:name/:provider/:version/upload",endpoints.Upload)
	r.GET("/:namespace/:name/:provider/:version/close",endpoints.Remove)
	r.GET("/:namespace/:name/:provider/:version/publish",endpoints.Publish)
}

func AuthRoutes(r *gin.RouterGroup) {
	r.GET("/authorize", auth.Authorize)
	r.POST("/token", auth.Token)
	r.GET("/callback", auth.Callback)
}

func WebhookRoutes(r *gin.RouterGroup) {
	r.POST("/webhook", webhook.HandleGithubWebhook)
}