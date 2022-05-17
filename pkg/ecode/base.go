package ecode

//go:generate codegen -type=int

// Common: basic errors.
// Code must start with 1xxxxx.
const (
	// Success - 200: OK.
	Success int = iota + 100000
	// ErrUnknown - 500: 服务端错误.
	ErrUnknown
	// ErrGet - 500: 数据获取失败.
	ErrGet
	// ErrCreate - 500: 创建失败.
	ErrCreate
	// ErrUpdate - 500: 更新失败.
	ErrUpdate
	// ErrDelete - 500: 删除失败.
	ErrDelete
	// ErrParam - 400: 参数异常.
	ErrParam
	// ErrAuth - 401: 权限异常.
	ErrAuth
	// ErrDataNotFound - 404: 数据未找到.
	ErrDataNotFound
	// ErrPageNotFound - 404: 页面未找到.
	ErrPageNotFound
	// ErrValidation - 400: 参数验证失败.
	ErrValidation
	// ErrTimeOut - 503: 服务端响应超时.
	ErrTimeOut
)
