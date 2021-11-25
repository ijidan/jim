package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jim/internal/global"
	"net/http"
)

// Pong pong
func Pong(content *gin.Context) {
	err := errors.New("TEST")
	panic(err)
	rsp := global.Response.JsonSuccess(gin.H{"message": "pong"}, "")
	content.JSON(http.StatusOK, rsp)
}
