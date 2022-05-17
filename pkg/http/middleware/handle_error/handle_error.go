package handle_error

import (
	"gin_example_with_generic/pkg/log"
	"gin_example_with_generic/types"
	"github.com/gin-gonic/gin"
)

// HandleError 错误处理
func HandleError(ctx *gin.Context) {
	ctx.Next()
	if warnInfo, exist := ctx.Get(types.WarnAppLogLevel); exist {
		log.Warnf("%v", warnInfo)
	}
	if errInfo, exist := ctx.Get(types.ErrorAppLogLevel); exist {
		log.Errorf("%+v", errInfo)
	}
}
