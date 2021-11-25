package middleware

import (
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
