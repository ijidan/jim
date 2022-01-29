package request

import "mime/multipart"

type CommonImageRequest struct {
	Ext string `form:"ext" json:"ext" binding:"required,oneof=jpg|jpeg|png"`
}
type CommonUploadImageRequest struct {
	File *multipart.FileHeader `form:"file" json:"file" binding:"required,image" `
}
