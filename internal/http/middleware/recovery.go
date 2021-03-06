package middleware

import (
	"github.com/gin-gonic/gin"
	"jim/pkg"
	"net/http"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				message := error2Message(r)
				if gin.IsDebugging() {
					debug.PrintStack()
				}
				pkg.Logger.Fatal(r)
				rsp := pkg.Rsp.JsonFail(pkg.ServerError, pkg.OK, message, nil, "")
				context.JSON(http.StatusOK, rsp)
				context.Abort()
			}
		}()
		context.Next()
	}
}

func error2Message(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
