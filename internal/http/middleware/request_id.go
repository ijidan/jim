package middleware

import (
	"github.com/gin-gonic/gin"
	"jim/internal/global"
	"jim/pkg"
)

func RequestId() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.Request.Header.Get(global.RequestId)
		if requestId == "" {
			requestId = pkg.GetUUId()
		}
		context.Set(global.RequestId, requestId)
		context.Writer.Header().Set(global.RequestId, requestId)

		context.Next()
	}
}
