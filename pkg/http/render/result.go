package render

import (
	"encoding/xml"
	"gin_example_with_generic/pkg/ecode"
	"gin_example_with_generic/pkg/errors"
	"gin_example_with_generic/types"
	"gin_example_with_generic/types/paginate"
	"github.com/gin-gonic/gin"
	"sync"
)

type PaginateInterface interface {
	GetPaginate() *paginate.Info
	GetItems() any
}

type Result struct {
	XMLName xml.Name `json:"-" xml:"result" yaml:"-"`
	Data    any      `json:"data" xml:"data" yaml:"data"` // 返回数据/错误详细信息
	Msg     string   `json:"msg" xml:"msg" yaml:"msg"`    // 请求结果
	Code    int      `json:"code" xml:"code" yaml:"code"` // 状态码
}

type PaginateData struct {
	Items    any            `json:"items" xml:"items" yaml:"items"`          // 数据详情列表
	Paginate *paginate.Info `json:"paginate" xml:"paginate" yaml:"paginate"` // 分页信息
}

func (r *Result) reset() {
	r.Data = nil
	r.Msg = ""
	r.Code = 0
}

var resultPool = &sync.Pool{
	New: func() any {
		return new(Result)
	},
}

func Response(ctx *gin.Context, rawData ...any) {
	if ctx.IsAborted() {
		return
	}
	ctx.Abort()

	result := resultPool.Get().(*Result)
	defer func() {
		result.reset()
		resultPool.Put(result)
	}()

	var (
		httpStatus int
		data       any
	)

	if len(rawData) > 0 {
		data = rawData[0]
	} else {
		data = nil
	}

	switch d := data.(type) {
	case error:
		coder := errors.ParseCoder(d)
		result.Code = coder.Code()
		if coder.String() != "" {
			result.Msg = coder.String()
		} else {
			result.Msg = d.Error()
		}
		httpStatus = coder.HTTPStatus()
		if httpStatus >= 400 && httpStatus < 500 {
			ctx.Set(types.WarnAppLogLevel, d)
		} else {
			ctx.Set(types.ErrorAppLogLevel, d)
		}
		result.Data = d.Error()
	case PaginateInterface:
		p := d.GetPaginate()
		if p == nil {
			p = &paginate.Info{}
		}

		if p.PageSize == 0 {
			p.PageSize = int(p.Total)
		}
		result.Code = ecode.Success
		result.Msg = types.SuccessMsg
		result.Data = PaginateData{
			Items:    d.GetItems(),
			Paginate: p,
		}
		httpStatus = types.SuccessHttpStatus
	default:
		result.Code = ecode.Success
		result.Msg = types.SuccessMsg
		result.Data = d
		httpStatus = types.SuccessHttpStatus
	}

	switch ctx.GetHeader("accept") {
	case "application/xml":
		ctx.XML(httpStatus, result)
	case "application/x-yaml":
		ctx.YAML(httpStatus, result)
	case "application/json":
		ctx.JSON(httpStatus, result)
	default:
		ctx.JSON(httpStatus, result)
	}
}
