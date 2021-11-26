package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"jim/internal/global"
	"net/http"
)

// Pong pong
func Pong(content *gin.Context) {
	tracerCtx, exists := content.Get("tracer_ctx")
	if exists {
		span, _ := opentracing.StartSpanFromContext(tracerCtx.(context.Context), "Pong")
		defer span.Finish()
	}

	rsp := global.Response.JsonSuccess(gin.H{"message": "pong"}, "")
	content.JSON(http.StatusOK, rsp)
}
