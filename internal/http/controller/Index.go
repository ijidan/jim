package controller

import (
	"github.com/gin-gonic/gin"
	"jim/internal/global"
)

// Pong pong
func Pong(content *gin.Context) {
	rsp := global.Reponse.JsonSuccess(gin.H{"message": "pong"}, "")
	content.JSON(200, rsp)
}
