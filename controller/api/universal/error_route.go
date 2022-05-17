package universal

import (
	"gin_example_with_generic/pkg/ecode"
	"gin_example_with_generic/pkg/errors"
	"gin_example_with_generic/pkg/http/render"
	"github.com/gin-gonic/gin"
)

func NoRoute(ctx *gin.Context) {
	render.Response(ctx, errors.WithCode(ecode.ErrPageNotFound, "Page not found"))
}
