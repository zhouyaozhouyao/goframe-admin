/**
 * 定义消息状态码返回
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/11 5:25 下午
 */

package e

const (
	Success                 = 200   // 请求成功
	Fail                    = 400   // 请求失败
	Error                   = 500   // 服务器内部错误
	Unauthorized            = 401   // 身份未授权
	ErrorExist              = 10001 // 数据已存在
	ErrorNotExist           = 10002 // 数据不存在
	ErrorCreateFail         = 10003 // 数据创建失败
	ErrorUpdateFail         = 10004 // 数据更新失败
	ErrorDeleteFail         = 10005 // 数据删除失败
	ErrorSelectFail         = 10006 // 数据查询失败
	ErrorAuthCheckTokenFail = 40001 // Token鉴权失败
	ErrorLoadCasBinFail     = 40002 // 加载用户权限失败
)
