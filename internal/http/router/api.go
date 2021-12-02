package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"jim/config"
	"jim/internal/http/controller"
	"jim/internal/http/middleware"
	"jim/pkg"
	"sync"
)

func registerApi(r *gin.Engine) {
	index := controller.IndexController{}
	user := controller.UserController{}
	api := r.Group("/api")
	{
		api.GET("/ping", index.Pong)
		api.GET("/user/register", user.Register)
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
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			if err := v.RegisterValidation("mobile", pkg.MobileValidator); err != nil {
				panic(err)
			}
		}

		instance.Use(cors.Default())
		instance.Use(middleware.Recovery(), middleware.AppSetting(), middleware.AccessLog(), middleware.Jaeger())
		registerApi(instance)
	})
	return instance
}
