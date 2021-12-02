package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"jim/global"
	"jim/pkg"
	"net/http"
)

// Pong pong
func Pong(content *gin.Context) {
	tracerCtx, exists := content.Get("tracer_ctx")
	if exists {
		span, _ := opentracing.StartSpanFromContext(tracerCtx.(context.Context), "Controller:Pong")
		defer span.Finish()
	}

	userId := int64(1)
	token := pkg.GenJwtToken(userId, global.Config.Jwt.Secret)
	rsp := global.Response.JsonSuccess(gin.H{"message": "pong:" + token}, "")
	content.JSON(http.StatusOK, rsp)
}
