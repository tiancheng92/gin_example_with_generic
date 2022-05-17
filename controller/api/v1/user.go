package v1

import (
	"gin_example_with_generic/generic"
	"gin_example_with_generic/pkg/http/render"
	"gin_example_with_generic/service"
	"gin_example_with_generic/store/model"
	"gin_example_with_generic/types/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	generic.Controller[request.User, model.User]
	UserService service.UserInterface
}

func NewUserController(db *gorm.DB) *UserController {
	userService := service.NewUserService(db)
	return &UserController{
		generic.Controller[request.User, model.User]{userService},
		userService,
	}
}

func (u *UserController) ListByName(ctx *gin.Context) {
	res, err := u.UserService.ListByName(ctx, ctx.Param("name"))
	if err != nil {
		render.Response(ctx, err)
	}
	render.Response(ctx, res)
}
