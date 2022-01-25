package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"jim/internal/http/request"
	"jim/internal/rpc/jim_proto/proto_build"
	"jim/internal/rpc/service"
	"net/http"
)

type CommonController struct {
	BaseController
}

func (c *CommonController) UploadImage(ctx *gin.Context) {
	var req request.CommonUploadImageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
	}
	file, err1 := ctx.FormFile("file")
	if err1 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err1.Error(), nil, "")
	}
	reader, err2 := file.Open()
	if err2 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err2.Error(), nil, "")
	}
	fileBytes, err3 := ioutil.ReadAll(reader)
	if err3 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err3.Error(), nil, "")
	}
	serviceClient, err := service.GetServiceCommonClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
	}
	uploadImageClient, err5 := serviceClient.UploadImage(ctx)
	if err5 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err5.Error(), nil, "")
	}
	serviceReq := &proto_build.UploadImageRequest{
		Content: fileBytes,
	}
	serviceErr := uploadImageClient.Send(serviceReq)
	if serviceErr != nil {
		if serviceErr == io.EOF {
			err6 := uploadImageClient.CloseSend()
			if err6 != nil {
				c.JsonFail(ctx, http.StatusInternalServerError, 0, err6.Error(), nil, "")
			}
		} else {
			c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
		}
	}
	serviceRsp, err7 := uploadImageClient.CloseAndRecv()
	if err7 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err7.Error(), nil, "")
	}
	rsp := map[string]interface{}{"url": serviceRsp.Url}
	c.JsonSuccess(ctx, rsp, "")
}
