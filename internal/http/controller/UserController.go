package controller

import (
	"github.com/gin-gonic/gin"
	"jim/internal/http/request"
	"net/http"
)

type UserController struct {
	BaseController
}

func (c *UserController) Register(content *gin.Context) {
	var req request.UserRegisterRequest
	if err := content.ShouldBindJSON(&req); err != nil {
		c.JsonFail(content, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	c.JsonSuccess(content, nil, "")
}
