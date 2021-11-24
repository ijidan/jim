package router

import (
	"github.com/gin-gonic/gin"
	"jim/internal/controller"
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
		registerApi(instance)
	})
	return instance
}
