package controller

import (
	"github.com/gin-gonic/gin"
	"jim/internal/http/request"
	"jim/internal/rpc/jim_proto/proto_build"
	"jim/internal/rpc/service"
	"net/http"
)

type MessageController struct {
	BaseController
}

func (c *MessageController) SendUserTextMessage(ctx *gin.Context) {
	var req request.MessageSendUserTextMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserTextMessageRequest{
		ToUserId: req.ToUserId,
		Text: &proto_build.TextMessage{
			Content: req.Content,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserTextMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendUserLocationMessage(ctx *gin.Context) {
	var req request.MessageSendUserLocationMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserLocationMessageRequest{
		ToUserId: req.ToUserId,
		Location: &proto_build.LocationMessage{
			CoverImage: req.CoverImage,
			Lat:        req.Lat,
			Lng:        req.Lng,
			MapLink:    req.MapLink,
			Desc:       req.Desc,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserLocationMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendUserFaceMessage(ctx *gin.Context) {
	var req request.MessageSendUserFaceMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserFaceMessageRequest{
		ToUserId: req.ToUserId,
		Face: &proto_build.FaceMessage{
			Symbol: req.Symbol,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserFaceMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendUserSoundMessage(ctx *gin.Context) {
	var req request.MessageSendUserSoundMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserSoundMessageRequest{
		ToUserId: req.ToUserId,
		Sound: &proto_build.SoundMessage{
			Url:     req.Url,
			Size:    req.Size,
			Seconds: req.Seconds,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserSoundMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendUserVideoMessage(ctx *gin.Context) {
	var req request.MessageSendUserVideoMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserVideoMessageRequest{
		ToUserId: req.ToUserId,
		Video: &proto_build.VideoMessage{
			Size:        req.Size,
			Seconds:     req.Seconds,
			Url:         req.Url,
			Format:      req.Format,
			ThumbUrl:    req.ThumbUrl,
			ThumbSize:   req.ThumbSize,
			ThumbWidth:  req.ThumbWidth,
			ThumbHeight: req.ThumbHeight,
			ThumbFormat: service.FormatMap[req.ThumbFormat],
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserVideoMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendUserImageMessage(ctx *gin.Context) {
	var req request.MessageSendUserImageMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}

	serviceReq := &proto_build.SendUserImageMessageRequest{
		ToUserId: req.ToUserId,
		Image: &proto_build.ImageMessageItem{
			Type:   service.TypeMap[req.Type],
			Format: service.FormatMap[req.Format],
			Size:   req.Size,
			Width:  req.Width,
			Height: req.Height,
			Url:    req.Url,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserImageMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendUserFileMessage(ctx *gin.Context) {
	var req request.MessageSendUserFileMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserFileMessageRequest{
		ToUserId: req.ToUserId,
		File: &proto_build.FileMessage{
			Size: req.Size,
			Name: req.Name,
			Url:  req.Url,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserFileMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendGroupTextMessage(ctx *gin.Context) {
	var req request.MessageSendGroupTextMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendGroupTextMessageRequest{
		ToGroupId: req.ToGroupId,
		AtUserId: req.AtUserId,
		Text: &proto_build.TextMessage{
			Content: req.Content,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendGroupTextMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendGroupLocationMessage(ctx *gin.Context) {
	var req request.MessageSendUserLocationMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserLocationMessageRequest{
		ToUserId: req.ToUserId,
		Location: &proto_build.LocationMessage{
			CoverImage: req.CoverImage,
			Lat:        req.Lat,
			Lng:        req.Lng,
			MapLink:    req.MapLink,
			Desc:       req.Desc,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserLocationMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendGroupFaceMessage(ctx *gin.Context) {
	var req request.MessageSendUserFaceMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserFaceMessageRequest{
		ToUserId: req.ToUserId,
		Face: &proto_build.FaceMessage{
			Symbol: req.Symbol,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserFaceMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendGroupSoundMessage(ctx *gin.Context) {
	var req request.MessageSendUserSoundMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserSoundMessageRequest{
		ToUserId: req.ToUserId,
		Sound: &proto_build.SoundMessage{
			Url:     req.Url,
			Size:    req.Size,
			Seconds: req.Seconds,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserSoundMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendGroupVideoMessage(ctx *gin.Context) {
	var req request.MessageSendUserVideoMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendUserVideoMessageRequest{
		ToUserId: req.ToUserId,
		Video: &proto_build.VideoMessage{
			Size:        req.Size,
			Seconds:     req.Seconds,
			Url:         req.Url,
			Format:      req.Format,
			ThumbUrl:    req.ThumbUrl,
			ThumbSize:   req.ThumbSize,
			ThumbWidth:  req.ThumbWidth,
			ThumbHeight: req.ThumbHeight,
			ThumbFormat: service.FormatMap[req.ThumbFormat],
		},
	}
	serviceRsp, serviceErr := serviceClient.SendUserVideoMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendGroupImageMessage(ctx *gin.Context) {
	var req request.MessageSendGroupImageMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}

	serviceReq := &proto_build.SendGroupImageMessageRequest{
		ToGroupId: req.ToGroupId,
		AtUserId: req.AtUserId,
		Image: &proto_build.ImageMessageItem{
			Type:   service.TypeMap[req.Type],
			Format: service.FormatMap[req.Format],
			Size:   req.Size,
			Width:  req.Width,
			Height: req.Height,
			Url:    req.Url,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendGroupImageMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}

func (c *MessageController) SendGroupFileMessage(ctx *gin.Context) {
	var req request.MessageSendGroupFileMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceMessageClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	serviceReq := &proto_build.SendGroupFileMessageRequest{
		ToGroupId: req.ToGroupId,
		AtUserId: req.AtUserId,
		File: &proto_build.FileMessage{
			Size: req.Size,
			Name: req.Name,
			Url:  req.Url,
		},
	}
	serviceRsp, serviceErr := serviceClient.SendGroupFileMessage(ctx, serviceReq)
	if serviceErr != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
	}
	rsp := map[string]interface{}{"id": serviceRsp.Id}
	c.JsonSuccess(ctx, rsp, "")
}
