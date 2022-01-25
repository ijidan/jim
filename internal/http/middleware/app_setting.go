package middleware

import (
	sContext "context"
	"github.com/gin-gonic/gin"
	"jim/pkg"
	"time"
)

func AppSetting() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set(
			"app_name",
			pkg.Conf.App.Name,
		)
		context.Set("app_ver", pkg.Conf.App.Ver)
		requestId := context.Request.Header.Get(pkg.RequestId)
		if requestId == "" {
			requestId = pkg.GetUUId()
		}
		context.Set(pkg.RequestId, requestId)
		context.Writer.Header().Set(pkg.RequestId, requestId)

		ctx, cancel := sContext.WithTimeout(context, 2*time.Second)
		defer cancel()
		context.Set("cancel_ctx", ctx)

		context.Next()
	}
}
