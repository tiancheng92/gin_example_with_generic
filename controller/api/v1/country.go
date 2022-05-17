package v1

import (
	"gin_example_with_generic/generic"
	"gin_example_with_generic/service"
	"gin_example_with_generic/store/model"
	"gin_example_with_generic/types/request"
	"gorm.io/gorm"
)

type CountryController struct {
	*generic.Controller[request.Country, model.Country]
}

func NewCountryController(db *gorm.DB) *CountryController {
	return &CountryController{
		generic.NewController[request.Country, model.Country](service.NewCountryService(db)),
	}
}
