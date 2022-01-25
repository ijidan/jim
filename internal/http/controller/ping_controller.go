package controller

import (
	"github.com/gin-gonic/gin"
)

type PingController struct {
	BaseController
}

func (c *PingController) Pong(ctx *gin.Context) {
	c.JsonSuccess(ctx, gin.H{"message": "pong"}, "")
}
