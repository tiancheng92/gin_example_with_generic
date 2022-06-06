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
	Order   string
	OrderBy string
	Search  string
	Params  map[string][]string
	AllData bool
}

type Info struct {
	Total    int64 `json:"total" xml:"total" yaml:"total"`
	Page     int   `json:"page" xml:"page" yaml:"page"`
	PageSize int   `json:"page_size" xml:"page_size" yaml:"page_size"`
}
