package generic

import "gin_example_with_generic/types/paginate"

type Paginate[T ModelInterface] struct {
	PaginateQuery *paginate.Query
	Items         []*T
}

func (p *Paginate[T]) GetPaginate() *paginate.Info {
	return p.PaginateQuery.Info
}

func (p *Paginate[T]) GetItems() any {
	return p.Items
}
