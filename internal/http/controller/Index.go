package controller

import (
	"github.com/gin-gonic/gin"
	"jim/internal/global"
	"net/http"
)

// Pong pong
func Pong(content *gin.Context) {
	rsp := global.Response.JsonSuccess(gin.H{"message": "pong"}, "")
	content.JSON(http.StatusOK, rsp)
}
