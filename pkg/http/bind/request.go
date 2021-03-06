package bind

import (
	"gin_example_with_generic/pkg/http/render"
	"gin_example_with_generic/pkg/validator"
	"gin_example_with_generic/types/paginate"
	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/gf"
	"strconv"
)

func Body(ctx *gin.Context, ptr any) error {
	err := ctx.ShouldBind(ptr)
	if err != nil {
		render.Response(ctx, validator.HandleValidationErr(err))
	}
	return err
}

func Query(ctx *gin.Context, ptr any) error {
	err := ctx.ShouldBindQuery(ptr)
	if err != nil {
		render.Response(ctx, validator.HandleValidationErr(err))
	}
	return err
}

func Params(ctx *gin.Context, ptr any) error {
	err := ctx.ShouldBindUri(ptr)
	if err != nil {
		render.Response(ctx, validator.HandleValidationErr(err))
	}
	return err
}

func PaginateQuery(ctx *gin.Context) *paginate.Query {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", paginate.DefaultPageStr))
	if err != nil || page < 1 {
		page = paginate.DefaultPage
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", paginate.DefaultPageSizeStr))
	if err != nil || pageSize < 1 {
		pageSize = paginate.DefaultPageSize
	}

	allData := !gf.ArrayContains([]string{"", "false", "False", "FALSE", "0"}, ctx.DefaultQuery("all_data", ""))
	if allData {
		page = 1
		pageSize = 0
	}

	return &paginate.Query{
		Page:     page,
		PageSize: pageSize,
		Order:    ctx.DefaultQuery("order", paginate.DefaultOrder),
		OrderBy:  ctx.DefaultQuery("order_by", paginate.DefaultOrderBy),
		Search:   ctx.DefaultQuery("search", paginate.DefaultSearch),
		Params:   ctx.Request.URL.Query(),
		AllData:  allData,
	}
}
