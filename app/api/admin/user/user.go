/**
 * 用户管理相关操作
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/9 6:14 下午
 */

// Package user 用户管理相关操作
package user

import (
	"gadmin/app/model/users"
	"gadmin/library/base"
	"gadmin/library/helper"
	"gadmin/library/input"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Controller 定义操作的结构体
type Controller struct{}

// Index 显示用户列表信息
func (c *Controller) Index(r *ghttp.Request) {
	var req = input.ListParams(r)

	// 设置分页的默认值与limit最大值
	page, limit := helper.PageParam(req.Page, req.Limit)

	// 获取数据
	total, result := users.GetList(page, limit, req.Where, req.OrderBy)

	// 返回结果集
	base.Success(r, g.Map{"total": total, "page": req.Page, "lists": result})
}
