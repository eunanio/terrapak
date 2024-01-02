package router

import (
	"terrapak/internal/api/auth"
	"terrapak/internal/api/discovery"
	"terrapak/internal/api/modules"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load(); if err != nil {
		panic("Error loading .env file")
	}

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp","localhost:6379", "redis", []byte("iqjwjoidjwqh28eu282837982731jadhwahdiawhud"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/ping", Ping)
	moduleRouter := r.Group("v1/modules")
	ApiRouter 	 := r.Group("v1/api")
	authRouter 	 := r.Group("v1/auth")
	ModuleRoutes(moduleRouter)
	ApiRoutes(ApiRouter)
	AuthRoutes(authRouter)

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
	r.GET("/:namespace/:name/:provider/:version",modules.Read)
	r.POST("/:namespace/:name/:provider/:version/upload",modules.Upload)
	r.GET("/:namespace/:name/:provider/:version/close",modules.RemoveDraft)
	r.GET("/:namespace/:name/:provider/:version/publish",modules.PublishDraft)
}

func AuthRoutes(r *gin.RouterGroup) {
	r.GET("/authorize", auth.Authorize)
	r.GET("/token", auth.Token)
	r.GET("/callback", auth.Callback)
}