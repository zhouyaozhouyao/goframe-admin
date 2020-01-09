package users

import "github.com/gogf/gf/os/glog"

// Fill with you ideas below.

// GetOne 查询单条数据
func GetOne(where interface{}) *Entity {
	var entity *Entity
	res, err := Model.FindOne(where)
	if err != nil {
		glog.Error("数据查询失败", err)
		return entity
	}

	return res
}
