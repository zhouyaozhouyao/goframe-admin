package controller

import (
	"api/api/v1/admin"
	"api/internal/service"
	"context"
)

type cUser struct {
}

var User = cUser{}

func (c *cUser) Create(ctx context.Context, req *admin.UserCreateReq) (res *admin.UserRes, err error) {
	err = service.User().Create(ctx, req)
	return
}

func (c *cUser) Update(ctx context.Context, req *admin.UserUpdateReq) (res *admin.UserRes, err error) {
	err = service.User().Update(ctx, req)
	return
}

func (c *cUser) Delete(ctx context.Context, req *admin.UserDeleteReq) (res *admin.UserRes, err error) {
	err = service.User().Delete(ctx, req.Ids)
	return
}

func (c *cUser) List(ctx context.Context, req *admin.UserListReq) (res *admin.UserListRes, err error) {
	res, err = service.User().List(ctx, req)
	return
}
