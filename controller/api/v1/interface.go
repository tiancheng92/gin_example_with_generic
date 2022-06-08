package v1

import (
	"gin_example_with_generic/generic"
	"gin_example_with_generic/store/model"
	"gin_example_with_generic/types/request"
	"github.com/gin-gonic/gin"
)

type CountryInterface interface {
	generic.ControllerInterface[request.Country, model.Country]
}

type UserInterface interface {
	generic.ControllerInterface[request.User, model.User]
	ListByName(ctx *gin.Context)
}
