package paginate

const (
	DefaultPageSizeStr = "20"
	DefaultPageStr     = "1"
	DefaultPageSize    = 20
	DefaultPage        = 1
	DefaultSearch      = ""
	DefaultOrder       = "DESC"
	DefaultOrderBy     = "id"
)

type Query struct {
	*Info
	Order   string `default:"desc" enums:"desc,asc"`
	OrderBy string `default:"id" json:"order_by" form:"order_by"`
	Search  string
	Params  map[string][]string
	AllData bool `json:"all_data" form:"all_data" xml:"all_data" yaml:"all_data"`
}

type Info struct {
	Total    int64 `json:"total" xml:"total" yaml:"total" `
	Page     int   `json:"page" xml:"page" yaml:"page" default:"1" minimum:"0"`
	PageSize int   `json:"page_size" xml:"page_size" form:"page_size" yaml:"page_size" default:"20" minimum:"0" maximum:"1000"`
}

type List[T any] struct {
	PaginateQuery        *Query
	Items                []*T
	FuzzySearchFieldList []string
}

func (l *List[T]) GetPaginate() *Info {
	return l.PaginateQuery.Info
}

func (l *List[T]) GetItems() any {
	return l.Items
}
