package generic

import (
	"gin_example_with_generic/pkg/http/bind"
	"gin_example_with_generic/pkg/http/render"
	"github.com/gin-gonic/gin"
)

func NewController[R RequestInterface, M ModelInterface](service ServiceInterface[R, M]) *Controller[R, M] {
	return &Controller[R, M]{service}
}

type Controller[R RequestInterface, M ModelInterface] struct {
	ServiceInterface[R, M]
}

func (c *Controller[R, M]) Get(ctx *gin.Context) {
	pk, err := bind.ParamsID(ctx, "pk")
	if err != nil {
		return
	}
	res, err := c.ServiceInterface.Get(ctx, pk)
	if err != nil {
		render.Response(ctx, err)
		return
	}
	render.Response(ctx, res)
}

func (c *Controller[R, M]) List(ctx *gin.Context) {
	res, err := c.ServiceInterface.List(ctx, bind.PaginateQuery(ctx))
	if err != nil {
		render.Response(ctx, err)
		return
	}
	render.Response(ctx, res)
}

func (c *Controller[R, M]) Create(ctx *gin.Context) {
	var r R
	if err := bind.Body(ctx, &r); err != nil {
		return
	}

	res, err := c.ServiceInterface.Create(ctx, &r)
	if err != nil {
		render.Response(ctx, err)
		return
	}
	render.Response(ctx, res)
}

func (c *Controller[R, M]) Update(ctx *gin.Context) {
	pk, err := bind.ParamsID(ctx, "pk")
	if err != nil {
		return
	}

	var r R
	if err := bind.Body(ctx, &r); err != nil {
		return
	}

	res, err := c.ServiceInterface.Update(ctx, pk, &r)
	if err != nil {
		render.Response(ctx, err)
		return
	}
	render.Response(ctx, res)
}

func (c *Controller[R, M]) Delete(ctx *gin.Context) {
	pk, err := bind.ParamsID(ctx, "pk")
	if err != nil {
		return
	}
	if err = c.ServiceInterface.Delete(ctx, pk); err != nil {
		render.Response(ctx, err)
		return
	}
	render.Response(ctx)
}

type ControllerInterface[R RequestInterface, M ModelInterface] interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewReadOnlyController[M ModelInterface](service ReadOnlyServiceInterface[M]) *ReadOnlyController[M] {
	return &ReadOnlyController[M]{service}
}

type ReadOnlyController[M ModelInterface] struct {
	ReadOnlyServiceInterface[M]
}

func (roc *ReadOnlyController[M]) Get(ctx *gin.Context) {
	pk, err := bind.ParamsID(ctx, "pk")
	if err != nil {
		return
	}
	res, err := roc.ReadOnlyServiceInterface.Get(ctx, pk)
	if err != nil {
		render.Response(ctx, err)
		return
	}
	render.Response(ctx, res)
}

func (roc *ReadOnlyController[M]) List(ctx *gin.Context) {
	res, err := roc.ReadOnlyServiceInterface.List(ctx, bind.PaginateQuery(ctx))
	if err != nil {
		render.Response(ctx, err)
		return
	}
	render.Response(ctx, res)
}

type ReadOnlyControllerInterface[M ModelInterface] interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
}
