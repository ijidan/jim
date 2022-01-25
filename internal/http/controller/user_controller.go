package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"jim/internal/http/request"
	"jim/internal/rpc/jim_proto/proto_build"
	"jim/internal/rpc/service"
	"net/http"
)

type UserController struct {
	BaseController
}

func (c *UserController) Create(ctx *gin.Context) {
	var req request.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceUserClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.UserCreateRequest{
		Nickname:    req.Nickname,
		AvatarUrl:   "https://cdn.libravatar.org/static/img/nobody/80.png",
		Password:    req.Password,
		PasswordRpt: req.Password,
	}
	serviceRsp, serviceErr := serviceClient.CreateUser(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	rsp, _ := cast.ToStringMapE(serviceRsp.GetUser())
	c.JsonSuccess(ctx, rsp, "")
}

func (c *UserController) Get(ctx *gin.Context) {
	var req request.UserGetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceUserClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.UserGetRequest{Id: req.Id}
	serviceRsp, serviceErr := serviceClient.GetUser(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	rsp, _ := cast.ToStringMapE(serviceRsp.GetUser())
	c.JsonSuccess(ctx, rsp, "")
}

func (c *UserController) Query(ctx *gin.Context) {
	var req request.UserQueryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceUserClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.UserQueryRequest{
		Keyword:  req.Keyword,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	serviceRsp, serviceErr := serviceClient.QueryUser(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	pager, _ := cast.ToStringMapE(serviceRsp.GetPager())
	user, _ := cast.ToStringMapE(serviceRsp.GetUser())
	rsp := map[string]interface{}{"pager": pager, "users": user}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *UserController) Login(ctx *gin.Context) {
	var req request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceUserClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.UserLoginRequest{
		Nickname: req.Nickname,
		Password: req.Password,
	}
	serviceRsp, serviceErr := serviceClient.UserLogin(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	rsp := map[string]interface{}{
		"token": serviceRsp.GetToken(),
	}
	c.JsonSuccess(ctx, rsp, "")

}

func (c *UserController) UpdatePassword(ctx *gin.Context) {
	var req request.UserUpdatePassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceUserClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.UpdatePasswordRequest{
		Password:    req.Password,
		PasswordRpt: req.PasswordRpt,
	}
	_, serviceErr := serviceClient.UpdatePassword(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	c.JsonSuccess(ctx, nil, "")

}

func (c *UserController) UpdateAvatar(ctx *gin.Context) {
	var req request.UserUpdateAvatar
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceUserClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.UpdateAvatarRequest{Url: req.AvatarUrl}
	_, serviceErr := serviceClient.UpdateAvatar(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	c.JsonSuccess(ctx, nil, "")

}
