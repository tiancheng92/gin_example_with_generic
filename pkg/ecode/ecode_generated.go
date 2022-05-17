// Code generated by "codegen -type=int ../pkg/ecode"; DO NOT EDIT.

package ecode

// init register error codes
func init() {
	register(Success, 200, "OK")
	register(ErrUnknown, 500, "服务端错误")
	register(ErrGet, 500, "数据获取失败")
	register(ErrCreate, 500, "创建失败")
	register(ErrUpdate, 500, "更新失败")
	register(ErrDelete, 500, "删除失败")
	register(ErrParam, 400, "参数异常")
	register(ErrAuth, 401, "权限异常")
	register(ErrDataNotFound, 404, "数据未找到")
	register(ErrPageNotFound, 404, "页面未找到")
	register(ErrValidation, 400, "参数验证失败")
	register(ErrTimeOut, 503, "服务端响应超时")
}
