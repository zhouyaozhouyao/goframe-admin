package admin

import (
	v1 "api/api/v1"
	"api/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// SetUserReq 添加与编辑用户公共参数
type SetUserReq struct {
	UserName     string `json:"userName" v:"required#用户名不能为空" dc:"用户名"`
	UserPassword string `json:"userPassword" v:"required#密码不能为空" dc:"密码"`
	UserSalt     string `json:"userSalt"`
	UserStatus   int    `json:"userStatus" dc:"状态"`
	Mobile       string `json:"mobile" v:"required#手机号不能为空" dc:"手机号"`
	UserNickName string `json:"userNickName" v:"required#用户昵称不能为空" dc:"用户姓名"`
	Birthday     string `json:"birthday" dc:"生日"`
	UserEmail    string `json:"userEmail" dc:"邮箱"`
	Sex          int    `json:"sex" dc:"性别"`
	Avatar       string `json:"avatar" dc:"头像"`
	DeptId       int    `json:"deptId" json:"deptId" v:"required#部门不能为空" dc:"部门"`
	Remark       string `json:"remark" dc:"备注"`
	IsAdmin      int    `json:"isAdmin" dc:"是否为管理员 1是 0否"`
	Address      string `json:"address" dc:"地址"`
	Describe     string `json:"describe" dc:"描述信息"`
	PostIds      []int  `json:"postIds" dc:"岗位"`
	RoleIds      []int  `json:"roleIds" dc:"角色"`
}

type UserCreateReq struct {
	g.Meta `path:"/user/create" tags:"用户管理" method:"POST" summary:"添加用户" dc:"添加员工或管理员"`
	SetUserReq
}

type UserUpdateReq struct {
	g.Meta `path:"/user/update" tags:"用户管理" method:"PUT" summary:"编辑用户" dc:"编辑员工或管理员"`
	SetUserReq
	Id int `json:"id" v:"required#用户id不能为空" dc:"用户id"`
}

type UserDeleteReq struct {
	g.Meta `path:"/user/delete" tags:"用户管理" method:"DELETE" summary:"删除用户" dc:"删除员工或管理员"`
	Ids    []int `json:"id" v:"required#用户id不能为空" dc:"用户id"`
}

type UserListReq struct {
	g.Meta `path:"/user/list" tags:"用户管理" method:"GET" summary:"获取用户列表" dc:"获取用户列表"`
	v1.PageReq
	UserName     string `json:"userName" dc:"用户名"`
	Mobile       string `json:"mobile" dc:"手机号"`
	UserNickName string `json:"userNickName" dc:"真实姓名"`
	UserStatus   int    `json:"userStatus" dc:"用户状态"`
	DeptId       int    `json:"deptId" dc:"部门id"`
}

type UserListRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
	v1.ListRes
	List []*entity.User `json:"list" dc:"角色列表"`
}

type UserRes struct {
	g.Meta `mime:"application/json"`
}
