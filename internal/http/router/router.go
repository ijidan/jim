package router

import (
	"github.com/gin-gonic/gin"
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

func NewGin() *gin.Engine {
	var once sync.Once
	var instance *gin.Engine
	once.Do(func() {
		instance = gin.Default()
		instance.Use(middleware.Recovery(), middleware.Logger(), middleware.RequestId())
		registerApi(instance)
	})
	return instance
}
