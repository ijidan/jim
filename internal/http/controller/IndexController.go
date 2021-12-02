package controller

import (
	"github.com/gin-gonic/gin"
	"jim/global"
	"jim/pkg"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Pong(content *gin.Context) {
	userId := int64(1)
	token := pkg.GenJwtToken(userId, global.Config.Jwt.Secret)
	c.JsonSuccess(content, gin.H{"message": "pong:" + token}, "")
}
