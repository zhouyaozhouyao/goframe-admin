/**
 * 定义返回格式
 * @email 994914376@qq.com
 * @Author: zhouyao
 * @Date: 2019/12/20 1:16 下午
 */

package response

import (
	"encoding/json"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gconv"
)

// Resp 返回结构体
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 获取Data值转字符串
func (resp Resp) Success() bool {
	return resp.Code == 200
}

// DataString 获取Data转字符串
func (resp Resp) DataString() string {
	return gconv.String(resp.Data)
}

// DataInt 获取Data转Int
func (resp Resp) DataInt() int {
	return gconv.Int(resp.Data)
}

// GetString 获取Data值转字符串
func (resp Resp) GetString(key string) string {
	return gconv.String(resp.Get(key))
}

// GetInt 获取Data值转Int
func (resp Resp) GetInt(key string) int {
	return gconv.Int(resp.Get(key))
}

// Get 获取Data值
func (resp Resp) Get(key string) interface{} {
	m := gconv.Map(resp.Data)
	if m == nil {
		return ""
	}
	return m[key]
}

// JSON 返回json字符串
func (resp Resp) JSON() string {
	str, _ := json.Marshal(resp)
	return string(str)
}

// Success 成功
func Success(data interface{}) Resp {
	return Resp{200, "success", data}
}

// Fail  失败
func Fail(code int) Resp {
	var msg = gconv.String(g.Cfg("message").Get("msg." + string(code)))
	return Resp{code, msg, []int{}}
}

// FailData 失败设置Data
func FailData(code int, data interface{}) Resp {
	var msg = gconv.String(g.Cfg("message").Get("msg." + string(code)))
	return Resp{code, msg, data}
}

// Error 错误
func Error(code int) Resp {
	var msg = gconv.String(g.Cfg("message").Get("msg." + string(code)))
	return Resp{code, msg, []int{}}
}

// ErrorData 错误设置Data
func ErrorData(code int, data interface{}) Resp {
	var msg = gconv.String(g.Cfg("message").Get("msg." + string(code)))
	return Resp{code, msg, data}
}

// FailParam 校验参数提示信息
func FailParam(msg string) Resp {
	return Resp{1003, msg, []int{}}
}
