package controller

import (
	"api/api/v1/admin"
	"api/internal/service"
	"context"
)

type cPost struct {
}

var (
	Post = cPost{}
)

// List 岗位列表
func (c *cPost) List(ctx context.Context, req *admin.PostListReq) (res *admin.PostListRes, err error) {
	res, err = service.Post().List(ctx, req)
	return
}

func (c *cPost) Create(ctx context.Context, req *admin.PostCreateReq) (res *admin.PostCreateRes, err error) {
	err = service.Post().Create(ctx, req)
	return
}

func (c *cPost) Update(ctx context.Context, req *admin.PostUpdateReq) (res *admin.PostUpdateRes, err error) {
	err = service.Post().Update(ctx, req)
	return
}

func (c *cPost) Delete(ctx context.Context, req *admin.PostDeleteReq) (res *admin.PostDeleteRes, err error) {
	err = service.Post().Delete(ctx, req)
	return
}
