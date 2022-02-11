package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"jim/pkg"
)

type PingController struct {
	BaseController
}

func (c *PingController) Pong(ctx *gin.Context) {
	rsp:=gin.H{"message": "pong"}

	tracerCtx,_:=ctx.Get("tracer_ctx")
	span,_:=opentracing.StartSpanFromContext(tracerCtx.(context.Context),pkg.GetCallerName())
	span.SetTag("request",ctx.Params)
	span.SetTag("rsp",rsp)

	defer func() {
		span.Finish()
	}()
	c.JsonSuccess(ctx, gin.H{"message": "pong"}, "")
}
