package service

import (
	"context"
	"gin_example_with_generic/generic"
	"gin_example_with_generic/store/model"
	"gin_example_with_generic/types/request"
)

type CountryInterface interface {
	generic.ServiceInterface[request.Country, model.Country]
}

type UserInterface interface {
	generic.ServiceInterface[request.User, model.User]
	ListByName(ctx context.Context, name string) ([]*model.User, error)
}
