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
	common:=controller.CommonController{}
	group := controller.GroupController{}
	message := controller.MessageController{}
	ping := controller.PingController{}
	user := controller.UserController{}

	apiV1Login := r.Group("/api/v1").Use(middleware.JwtVerify())
	{
		apiV1Login.POST("/group/create", group.Create)
		apiV1Login.POST("/group/update", group.Update)
		apiV1Login.DELETE("/group/delete", group.Delete)
		apiV1Login.POST("/group/join", group.Join)
		apiV1Login.POST("/group/quit", group.Quit)

		apiV1Login.POST("/message/user/text", message.SendUserTextMessage)
		apiV1Login.POST("/message/user/location", message.SendUserLocationMessage)
		apiV1Login.POST("/message/user/face", message.SendUserFaceMessage)
		apiV1Login.POST("/message/user/sound", message.SendUserSoundMessage)
		apiV1Login.POST("/message/user/video", message.SendUserVideoMessage)
		apiV1Login.POST("/message/user/image", message.SendUserImageMessage)
		apiV1Login.POST("/message/user/file", message.SendUserFileMessage)
		apiV1Login.POST("/message/group/text", message.SendGroupTextMessage)
		apiV1Login.POST("/message/group/location", message.SendGroupLocationMessage)
		apiV1Login.POST("/message/group/face", message.SendGroupFaceMessage)
		apiV1Login.POST("/message/group/sound", message.SendGroupSoundMessage)
		apiV1Login.POST("/message/group/video", message.SendGroupVideoMessage)
		apiV1Login.POST("/message/group/image", message.SendGroupImageMessage)
		apiV1Login.POST("/message/group/file", message.SendGroupFileMessage)

		apiV1Login.POST("/user/password", user.UpdatePassword)
		apiV1Login.POST("/usr/avatar", user.UpdateAvatar)
	}
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/common/image",common.UploadImage)

		apiV1.GET("/group/get", group.Get)
		apiV1.GET("/group/query", group.Query)

		apiV1.GET("/ping", ping.Pong)

		apiV1.GET("/user/get", user.Get)
		apiV1.GET("/user/query", user.Query)
		apiV1.POST("/usr/login", user.Login)
	}
}

func NewGin(conf *config.Config) *gin.Engine {
	var once sync.Once
	var instance *gin.Engine
	once.Do(func() {
		env := conf.App.Env
		switch env {
		case config.EnvLocal:
			gin.SetMode(gin.DebugMode)
			break
		case config.EnvTest:
			gin.SetMode(gin.TestMode)
			break
		case config.EnvStage:
			gin.SetMode(gin.TestMode)
		case config.EnvProduction:
		default:
			gin.SetMode(gin.ReleaseMode)
			break
		}
		instance = gin.Default()
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			if err := v.RegisterValidation("mobile", pkg.MobileValidator); err != nil {
				panic(err)
			}
			if err := v.RegisterValidation("image", pkg.ImageValidator); err != nil {
				panic(err)
			}
		}

		instance.Use(cors.Default())
		instance.Use(middleware.Recovery(), middleware.AppSetting(), middleware.AccessLog(), middleware.Jaeger())
		registerApi(instance)
	})
	return instance
}
