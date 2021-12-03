package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jim/global"
	"os"
	"time"
)

func AccessLog() gin.HandlerFunc {
	return func(context *gin.Context) {

		// Start timer
		start := time.Now()
		path := context.Request.URL.Path
		raw := context.Request.URL.RawQuery

		// Process request
		context.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)
		clientIP := context.ClientIP()
		method := context.Request.Method
		statusCode := context.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}

		requestId := context.GetString(global.RequestId)
		fields := logrus.Fields{
			"startTime":  start,
			"endTime":    end,
			"latency":    latency,
			"requestId":  requestId,
			"statusCode": statusCode,
			"clientIP":   clientIP,
			"method":     method,
			"path":       path,
		}
		log := fmt.Sprintf("[GIN] %v | %15s | %3d | %13v | %15s | %-7s %s\n", end.Format("2006/01/02 - 15:04:05"),
			requestId, statusCode, latency, clientIP, method, path)
		global.Logger.WithFields(fields).Info("")
		if gin.IsDebugging() {
			_, _ = fmt.Fprintf(os.Stdout, log)
		}
	}
}
