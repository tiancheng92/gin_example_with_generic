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

// Query 分页查询
type Query struct {
	Page     int                 // 页数
	PageSize int                 // 每页数据量
	Order    string              // 排序方式
	OrderBy  string              // 排序字段
	Search   string              // 关键字搜索
	Params   map[string][]string // 其他参数
	AllData  bool                // 是否查询所有数据
}

type Info struct {
	Total    int64 `json:"total" xml:"total" yaml:"total"`             // 数据总数
	Page     int   `json:"page" xml:"page" yaml:"page"`                // 页数
	PageSize int   `json:"page_size" xml:"page_size" yaml:"page_size"` // 每页数据量
}
