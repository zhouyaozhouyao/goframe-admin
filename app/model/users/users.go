package users

import (
	"gadmin/library/helper"

	"github.com/gogf/gf/os/glog"
)

var (
	entity *Entity
)

// GetOne 查询单条数据
func GetOne(where ...interface{}) *Entity {
	res, err := Model.FindOne(where...)
	if err != nil {
		glog.Error("数据查询失败", err)
		return entity
	}

	return res
}

// GetCount 显示查询结果总条数
func GetCount(where ...interface{}) int {
	total, err := Model.FindCount(where...)
	if err != nil {
		glog.Error("获取用户数量错误", err)
		return 0
	}

	return total
}

// GetList 显示用户信息列表
func GetList(page, limit int, where interface{}, orderBy ...map[string]interface{}) (int, interface{}) {
	// 获取总数
	total := GetCount(where)
	// 把排序参数转换为字符串
	orderByStr := helper.OrderByParam(orderBy)
	// 查询sql语句
	result, err := Model.M.Page(page, limit).Fields("id, username, email, uuid").Order(orderByStr).FindAll(where)
	if err != nil {
		glog.Error("数据查询失败", err)
		return total, []int{}
	}

	if result.IsEmpty() {
		return total, []int{}
	}

	return total, result
}
