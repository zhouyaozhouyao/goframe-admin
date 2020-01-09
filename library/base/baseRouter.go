/**
 * 全局的常归操作
 * @Author: zhouyao
 * @Date: 2019/12/23 5:19 下午
 */

package base

import (
	"gadmin/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// Success 成功返回
func Success(r *ghttp.Request, data interface{}) {
	r.Response.WriteJson(response.Success(data))
	r.Exit()
}

// Fail 操作失败返回
func Fail(r *ghttp.Request, code int) {
	r.Response.WriteJson(response.Fail(code))
	r.Exit()
}

// Error 请求错误返回
func Error(r *ghttp.Request, code int) {
	r.Response.WriteJson(response.Error(code))
	r.Exit()
}

func FailParam(r *ghttp.Request, msg string) {
	r.Response.WriteJson(response.FailParam(msg))
	r.Exit()
}
