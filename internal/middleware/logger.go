package middleware

import (
	"github.com/gin-gonic/gin"
	"jim/internal/global"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Next()
		latency := time.Since(t)
		global.NewGlobal().Logger.Infoln(latency)
	}
}
