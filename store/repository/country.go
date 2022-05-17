package repository

import (
	"gin_example_with_generic/generic"
	"gin_example_with_generic/store/model"
	"gorm.io/gorm"
)

type CountryRepository struct {
	*generic.Repository[model.Country]
}

func NewCountryRepository(db *gorm.DB) CountryInterface {
	return &CountryRepository{
		generic.NewRepository[model.Country](db),
	}
}
