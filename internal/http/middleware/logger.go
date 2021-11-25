package middleware

import (
	"github.com/gin-gonic/gin"
	"jim/internal/global"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		global.Logger.WithField(global.RequestId, context.GetString(global.RequestId)).Info("x-request-id")
		context.Next()
	}
}
