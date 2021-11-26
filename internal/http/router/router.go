package router

import (
	"github.com/gin-gonic/gin"
	"jim/config"
	"jim/internal/http/controller"
	"jim/internal/http/middleware"
	"sync"
)

func registerApi(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", controller.Pong)
	}
}

func NewGin(conf *config.Config) *gin.Engine {
	var once sync.Once
	var instance *gin.Engine
	once.Do(func() {
		env := conf.App.Env
		switch env {
		case "local":
			gin.SetMode(gin.DebugMode)
			break
		case "test":
			gin.SetMode(gin.TestMode)
			break
		case "production":
		default:
			gin.SetMode(gin.ReleaseMode)
			break
		}
		instance = gin.Default()
		instance.Use(middleware.Recovery(), middleware.AppSetting(), middleware.AccessLog(), middleware.Jaeger(), middleware.JwtVerify())
		registerApi(instance)
	})
	return instance
}
