package generic

import (
	"context"
	"gin_example_with_generic/pkg/ecode"
	"gin_example_with_generic/pkg/errors"
	"gin_example_with_generic/types/paginate"
	"github.com/tiancheng92/gf"
	"gorm.io/gorm"
	"strings"
)

type Repository[M ModelInterface] struct {
	DB           *gorm.DB
	PaginateData *Paginate[M]
}

func NewRepository[M ModelInterface](db *gorm.DB) *Repository[M] {
	return &Repository[M]{
		DB: db,
	}
}

func (r *Repository[M]) Get(ctx context.Context, pk any) (*M, error) {
	var ent M
	err := r.DB.WithContext(ctx).Where(gf.StringJoin("`", ent.GetPrimaryKeyName(), "` = ?"), pk).First(&ent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(ecode.ErrDataNotFound, err)
		} else {
			return nil, errors.WithCode(ecode.ErrGet, err)
		}
	}
	return &ent, nil
}

func (r *Repository[M]) Create(ctx context.Context, attributes M) (*M, error) {
	err := r.DB.WithContext(ctx).Create(&attributes).Error
	return &attributes, errors.WithCode(ecode.ErrCreate, err)
}

func (r *Repository[M]) Update(ctx context.Context, pk any, attributes M) (*M, error) {
	err := r.DB.WithContext(ctx).Where(gf.StringJoin("`", attributes.GetPrimaryKeyName(), "` = ?"), pk).Updates(&attributes).Error
	return &attributes, errors.WithCode(ecode.ErrUpdate, err)
}

func (r *Repository[M]) Delete(ctx context.Context, pk any) error {
	var ent M
	err := r.DB.WithContext(ctx).Where(gf.StringJoin("`", ent.GetPrimaryKeyName(), "` = ?"), pk).Delete(&ent).Error
	return errors.WithCode(ecode.ErrDelete, err)
}

func (r *Repository[M]) List(ctx context.Context, pq *paginate.Query) (*Paginate[M], error) {
	err := r.DB.WithContext(ctx).Scopes(r.Paginate(pq)).Find(&r.PaginateData.Items).Offset(-1).Limit(-1).Count(&r.PaginateData.PaginateQuery.Total).Error
	return r.PaginateData, errors.WithCode(ecode.ErrGet, err)
}

func (r *Repository[M]) Paginate(pq *paginate.Query) func(db *gorm.DB) *gorm.DB {
	r.PaginateData = &Paginate[M]{
		PaginateQuery: pq,
	}
	return func(db *gorm.DB) *gorm.DB {
		var fieldList []string
		stmt := &gorm.Statement{DB: db}
		_ = stmt.Parse(new(M))
		for i := range stmt.Schema.Fields {
			if stmt.Schema.Fields[i].DBName != "" {
				fieldList = append(fieldList, stmt.Schema.Fields[i].DBName)
			}
		}

		if !r.PaginateData.PaginateQuery.AllData {
			db = db.Offset((r.PaginateData.PaginateQuery.Page - 1) * r.PaginateData.PaginateQuery.PageSize).Limit(r.PaginateData.PaginateQuery.PageSize)
		}

		for k, v := range r.PaginateData.PaginateQuery.Params {
			if gf.ArrayContains(fieldList, k) {
				if len(v) == 1 {
					if v[0] != "" {
						db = db.Where(gf.StringJoin("`", k, "` = ?"), v[0])
					}
				} else {
					db = db.Where(gf.StringJoin("`", k, "` IN (?)"), v)
				}
			}

			fieldSlice := strings.Split(k, "__")
			if len(fieldSlice) >= 2 {
				suffix := fieldSlice[len(fieldSlice)-1]
				field := strings.ReplaceAll(k, gf.StringJoin("__", suffix), "")
				if gf.ArrayContains(fieldList, field) {
					for i := range v {
						if v[i] != "" {
							switch suffix {
							case "gte":
								db = db.Where(gf.StringJoin("`", field, "` >= ?"), v[i])
							case "gt":
								db = db.Where(gf.StringJoin("`", field, "` > ?"), v[i])
							case "lte":
								db = db.Where(gf.StringJoin("`", field, "` <= ?"), v[i])
							case "lt":
								db = db.Where(gf.StringJoin("`", field, "` < ?"), v[i])
							case "ne":
								db = db.Where(gf.StringJoin("`", field, "` != ?"), v[i])
							case "sw":
								db = db.Where(gf.StringJoin("`", field, "` LIKE ?"), gf.StringJoin(v[i], "%"))
							case "ew":
								db = db.Where(gf.StringJoin("`", field, "` LIKE ?"), gf.StringJoin("%", v[i]))
							case "like":
								db = db.Where(gf.StringJoin("`", field, "` LIKE ?"), gf.StringJoin("%", v[i], "%"))
							case "json_contains":
								db = db.Where(gf.StringJoin("JSON_CONTAINS(`", field, "`, CONCAT('\"', \"", v[i], "\", '\"'))"))
							case "json_extract":
								jsonKV := strings.Split(v[i], "//")
								if len(jsonKV) == 2 {
									if strings.Contains(jsonKV[0], "[*]") {
										db = db.Where(gf.StringJoin("JSON_CONTAINS(JSON_EXTRACT(`", field, "`, '", jsonKV[0], "'), CONCAT('\"', '", jsonKV[1], "', '\"'))"))
									} else {
										db = db.Where(gf.StringJoin("JSON_EXTRACT(`", field, "`, \"", jsonKV[0], "\") = ?"), jsonKV[1])
									}
								}
							}
						}
					}
				}
			}
		}

		fuzzySearchFieldList := (*new(M)).GetFuzzySearchFieldList()
		if r.PaginateData.PaginateQuery.Search != "" && len(fuzzySearchFieldList) > 0 {
			var searchField = make([]string, 0, len(fuzzySearchFieldList))
			for i := range fuzzySearchFieldList {
				searchField = append(searchField, gf.StringJoin("IFNULL(`", strings.TrimSpace(fuzzySearchFieldList[i]), "`, '')"))
			}
			db = db.Where(gf.StringJoin("CONCAT(", strings.Join(searchField, ", "), ") LIKE ?"), gf.StringJoin("%", r.PaginateData.PaginateQuery.Search, "%"))
		}

		if !gf.ArrayContains(fieldList, r.PaginateData.PaginateQuery.OrderBy) {
			r.PaginateData.PaginateQuery.OrderBy = paginate.DefaultOrderBy
		}

		if !gf.ArrayContains([]string{"desc", "asc"}, r.PaginateData.PaginateQuery.Order) {
			r.PaginateData.PaginateQuery.Order = paginate.DefaultOrder
		}

		db = db.Order(gf.StringJoin("`", r.PaginateData.PaginateQuery.OrderBy, "` ", strings.ToUpper(r.PaginateData.PaginateQuery.Order)))

		return db
	}
}

type RepositoryInterface[M ModelInterface] interface {
	Get(ctx context.Context, pk any) (*M, error)
	Create(ctx context.Context, attributes M) (*M, error)
	Update(ctx context.Context, pk any, attributes M) (*M, error)
	Delete(ctx context.Context, pk any) error
	List(ctx context.Context, pq *paginate.Query) (*Paginate[M], error)
}
