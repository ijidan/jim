package router

import (
	"github.com/gin-gonic/gin"
	"jim/internal/controller"
	"jim/internal/middleware"
	"sync"
)

func registerApi(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", controller.Pong)

	}
}

func NewGin() *gin.Engine {
	var once sync.Once
	var instance *gin.Engine
	once.Do(func() {
		instance = gin.Default()
		instance.Use(gin.Recovery())
		instance.Use(middleware.Logger())
		registerApi(instance)
	})
	return instance
}
