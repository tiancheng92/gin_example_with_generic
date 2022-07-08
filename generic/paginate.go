package generic

import "gin_example_with_generic/types/paginate"

type Paginate[M ModelInterface] struct {
	*paginate.Info
	Items []*M
}

func (p *Paginate[M]) GetPaginate() *paginate.Info {
	return p.Info
}

func (p *Paginate[M]) GetItems() any {
	return p.Items
}
