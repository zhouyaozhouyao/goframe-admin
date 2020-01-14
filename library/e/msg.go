/**
 * 定义消息提示返回
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/11 5:25 下午
 */

package e

var Message = map[int]string{
	Success:                 "请求成功",
	Fail:                    "请求失败",
	Error:                   "服务器内部错误",
	Unauthorized:            "身份未授权",
	Forbidden:               "没有访问权限",
	ErrorExist:              "数据已存在",
	ErrorNotExist:           "数据不存在",
	ErrorCreateFail:         "数据创建失败",
	ErrorUpdateFail:         "数据更新失败",
	ErrorDeleteFail:         "数据删除失败",
	ErrorSelectFail:         "数据查询失败",
	ErrorAuthCheckTokenFail: "Token鉴权失败",
	ErrorLoadCasBinFail:     "加载用户权限失败",
}

// GetMsg 根据状态码转换消息体
func GetMsg(code int) string {
	if msg, err := Message[code]; err {
		return msg
	} else {
		return Message[Fail]
	}
}
