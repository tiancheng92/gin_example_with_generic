package repository

import (
	"context"
	"gin_example_with_generic/generic"
	"gin_example_with_generic/store/model"
)

type CountryInterface interface {
	generic.RepositoryInterface[model.Country]
}

type UserInterface interface {
	generic.RepositoryInterface[model.User]
	ListByName(ctx context.Context, name string) ([]*model.User, error)
}
