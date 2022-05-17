package service

import (
	"gin_example_with_generic/generic"
	"gin_example_with_generic/store/model"
	"gin_example_with_generic/store/repository"
	"gin_example_with_generic/types/request"
	"gorm.io/gorm"
)

type CountryService struct {
	generic.Service[request.Country, model.Country]
}

func NewCountryService(db *gorm.DB) CountryInterface {
	return &CountryService{
		generic.Service[request.Country, model.Country]{
			repository.NewCountryRepository(db),
		},
	}
}
