package service

import (
	"jim/internal/repository"
	"jim/internal/rpc/jim_proto/proto_build"
)

var TypeMap = map[uint64]proto_build.IMAGE_TYPE{
	repository.ImageTypeOrigin:    proto_build.IMAGE_TYPE_ORIGIN,
	repository.ImageTypeBig:       proto_build.IMAGE_TYPE_BIG,
	repository.ImageTypeThumbnail: proto_build.IMAGE_TYPE_THUMBNAIL,
}
var FormatMap = map[uint64]proto_build.IMAGE_FORMAT{
	repository.ImageFormatJpg:   proto_build.IMAGE_FORMAT_JPG,
	repository.ImageFormatGif:   proto_build.IMAGE_FORMAT_GIF,
	repository.ImageFormatPng:   proto_build.IMAGE_FORMAT_PNG,
	repository.ImageFormatBmp:   proto_build.IMAGE_FORMAT_BMP,
	repository.ImageFormatOther: proto_build.IMAGE_FORMAT_OTHER,
}
