package pkg

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type HttpResponse struct {
}

func (r *HttpResponse) JsonSuccess(data map[string]interface{}, jumpUrl string) gin.H {
	return gin.H{"code": 0, "message": "success", "data": data, "jump_url": jumpUrl}
}

func (r *HttpResponse) JsonFail(code int32, businessCode int32, message string, data map[string]interface{}, jumpUrl string) gin.H {
	return gin.H{"code": code, "business_code": businessCode, "message": message, "data": data, "jump_url": jumpUrl}
}

func (r *HttpResponse) Pager() {
}

var (
	onceResponse     sync.Once
	instanceResponse *HttpResponse
)

func GetResponseInstance() *HttpResponse {
	onceResponse.Do(func() {
		instanceResponse = &HttpResponse{}
	})
	return instanceResponse
}
