/**
 * 定义返回格式
 * @email 994914376@qq.com
 * @Author: zhouyao
 * @Date: 2019/12/20 1:16 下午
 */

package response

import (
	"encoding/json"
	"gadmin/library/e"

	"github.com/gogf/gf/util/gconv"
)

// Resp 返回结构体
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
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
	var msg = e.GetMsg(code)
	return Resp{code, msg, []int{}}
}

// Error 错误
func Error(code int) Resp {
	var msg = e.GetMsg(code)
	return Resp{code, gconv.String(msg), []int{}}
}

// FailParam 校验参数提示信息
func FailParam(msg string) Resp {
	return Resp{400, msg, []int{}}
}
