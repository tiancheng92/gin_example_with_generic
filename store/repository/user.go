package repository

import (
	"context"
	"gin_example_with_generic/generic"
	"gin_example_with_generic/pkg/ecode"
	"gin_example_with_generic/pkg/errors"
	"gin_example_with_generic/store/model"
	"gin_example_with_generic/types/paginate"
	"github.com/tiancheng92/gf"
	"gorm.io/gorm"
)

type UserRepository struct {
	*generic.Repository[model.User]
}

func NewUserRepository(db *gorm.DB) UserInterface {
	return &UserRepository{
		generic.NewRepository[model.User](db),
	}
}

func (u *UserRepository) Get(ctx context.Context, pk any) (*model.User, error) {
	var ent model.User
	err := u.DB.WithContext(ctx).Preload("Country").Where(gf.StringJoin("`", ent.GetPrimaryKeyName(), "` = ?"), pk).First(&ent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(ecode.ErrDataNotFound, err)
		} else {
			return nil, errors.WithCode(ecode.ErrGet, err)
		}
	}
	return &ent, nil
}

func (u *UserRepository) List(ctx context.Context, pq *paginate.Query) (*generic.Paginate[model.User], error) {
	err := u.DB.WithContext(ctx).Scopes(u.Paginate(pq)).Preload("Country").Find(&u.PaginateData.Items).Offset(-1).Limit(-1).Count(&u.PaginateData.PaginateQuery.Total).Error
	return u.PaginateData, errors.WithCode(ecode.ErrGet, err)
}

func (u *UserRepository) ListByName(ctx context.Context, name string) ([]*model.User, error) {
	var ents []*model.User
	err := u.DB.WithContext(ctx).Preload("Country").Where("`name` = ?", name).Find(&ents).Error
	return ents, errors.WithCode(ecode.ErrGet, err)
}
