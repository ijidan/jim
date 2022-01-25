package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"jim/internal/http/request"
	"jim/internal/rpc/jim_proto/proto_build"
	"jim/internal/rpc/service"
	"net/http"
)

type GroupController struct {
	BaseController
}

func (c *GroupController) Create(ctx *gin.Context) {
	var req request.GroupCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceGroupClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.CreateGroupRequest{
		Name:         req.Name,
		Introduction: req.Introduction,
		AvatarUrl:    req.AvatarUrl,
	}
	serviceRsp, serviceErr := serviceClient.CreateGroup(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	rsp, _ := cast.ToStringMapE(serviceRsp.GetGroup())
	c.JsonSuccess(ctx, rsp, "")
}

func (c *GroupController) Update(ctx *gin.Context) {
	var req request.GroupUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceGroupClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.UpdateGroupRequest{}
	if req.Id > 0 {
		serviceReq.Id = req.Id
	}
	if len(req.Name) > 0 {
		serviceReq.Name = req.Name
	}
	if len(req.Introduction) > 0 {
		serviceReq.Introduction = req.Introduction
	}
	if len(req.AvatarUrl) > 0 {
		serviceReq.AvatarUrl = req.AvatarUrl
	}
	_, serviceErr := serviceClient.UpdateGroup(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	c.JsonSuccess(ctx, nil, "")
}

func (c *GroupController) Get(ctx *gin.Context) {
	var req request.GroupGetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceGroupClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.GetGroupRequest{
		Id: req.Id,
	}
	serviceRsp, serviceErr := serviceClient.GetGroup(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	rsp, _ := cast.ToStringMapE(serviceRsp.GetGroup())
	c.JsonSuccess(ctx, rsp, "")
}

func (c *GroupController) Query(ctx *gin.Context) {
	var req request.GroupQueryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceGroupClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.QueryGroupRequest{
		Keyword:  req.Keyword,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	serviceRsp, serviceErr := serviceClient.QueryGroup(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	pager, _ := cast.ToStringMapE(serviceRsp.GetPager())
	user, _ := cast.ToStringMapE(serviceRsp.GetGroup())
	rsp := map[string]interface{}{"pager": pager, "users": user}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *GroupController) Delete(ctx *gin.Context) {
	var req request.GroupDeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceGroupClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.DeleteGroupRequest{
		Id: req.Id,
	}
	_, serviceErr := serviceClient.DeleteGroup(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	c.JsonSuccess(ctx, nil, "")
}

func (c *GroupController) Join(ctx *gin.Context) {
	var req request.GroupJoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceGroupClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.JoinGroupRequest{
		Id:     req.Id,
		UserId: req.UserId,
	}
	_, serviceErr := serviceClient.JoinGroup(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")

	}
	c.JsonSuccess(ctx, nil, "")
}

func (c *GroupController) Quit(ctx *gin.Context) {
	var req request.GroupQuitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceGroupClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.QuitGroupRequest{
		Id:     req.Id,
		UserId: req.UserId,
	}
	_, serviceErr := serviceClient.QuitGroup(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	c.JsonSuccess(ctx, nil, "")
}
