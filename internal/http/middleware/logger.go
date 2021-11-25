package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jim/internal/global"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.GetString(global.RequestId)
		global.Logger.WithFields(logrus.Fields{global.RequestId: requestId}).Info(global.RequestId)
		context.Next()
	}
}
