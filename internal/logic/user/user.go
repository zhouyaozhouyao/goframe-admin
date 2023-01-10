package user

import (
	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/library/casbin"
	"api/internal/library/libUtils"
	"api/internal/library/liberr"
	"api/internal/model"
	"api/internal/model/do"
	"api/internal/model/entity"
	"api/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

type sUser struct {
	CasBinUserPrefix string
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{
		CasBinUserPrefix: "u_",
	}
}

// GetAdminUserByUsernamePassword 检测账号密码是否匹配
func (s *sUser) GetAdminUserByUsernamePassword(ctx context.Context, in *admin.LoginReq) (user *model.UserLoginOutput, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, err = s.GetUserByUsername(ctx, in.Username)
		liberr.IsNil(ctx, err)
		liberr.ValueIsNil(user, "账号密码错误")
		// 验证密码
		if libUtils.EncryptPassword(in.Password, user.UserSalt) != user.UserPassword {
			liberr.IsNil(ctx, gerror.New("账号密码错误"))
		}
		// 账号状态
		if user.UserStatus == 0 {
			liberr.IsNil(ctx, gerror.New("账号已被禁用"))
		}
	})
	return user, err
}

// GetUserByUsername 根据用户名获取用户信息
func (s *sUser) GetUserByUsername(ctx context.Context, userName string) (user *model.UserLoginOutput, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user = &model.UserLoginOutput{}
		err = dao.User.Ctx(ctx).Fields(user).Where(dao.User.Columns().UserName, userName).Scan(user)
		liberr.IsNil(ctx, err, "账号密码错误")
	})
	return
}

// UpdateLoginInfo 更新用户登录信息
func (s *sUser) UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.User.Ctx(ctx).WherePri(id).Data(do.User{
			LastLoginIp:   ip,
			LastLoginTime: gtime.Now(),
		}).Update()
		liberr.IsNil(ctx, err)
	})
	return err
}

func (s *sUser) Create(ctx context.Context, req *admin.UserCreateReq) (err error) {
	// 检测手机号和用户名称是否已存在
	err = s.userNameOrMobileExists(ctx, req.UserName, req.Mobile)
	if err != nil {
		return
	}
	req.UserSalt = grand.S(10) // 生成随机数 参与密码加密
	req.UserPassword = libUtils.EncryptPassword(req.UserPassword, req.UserSalt)
	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			userId, e := dao.User.Ctx(ctx).TX(tx).InsertAndGetId(req)
			liberr.IsNil(ctx, e, "创建用户失败")
			// 添加员工与岗位关联表
			e = s.AddUserPost(ctx, tx, req.PostIds, gconv.Int(userId))
			liberr.IsNil(ctx, e)
			// 添加员工与角色关联表
			e = s.AddUserRole(ctx, req.RoleIds, gconv.Int(userId))
			liberr.IsNil(ctx, e)
		})
		return err
	})
	return err
}

func (s *sUser) Update(ctx context.Context, req *admin.UserUpdateReq) (err error) {
	// 检测手机号和用户名称是否已存在
	err = s.userNameOrMobileExists(ctx, req.UserName, req.Mobile, gconv.Int64(req.Id))
	if err != nil {
		return
	}
	// 查询当前用户信息
	user, err := s.GetUserByUsername(ctx, req.UserName)
	if err != nil {
		return
	}
	// 检测密码是否有修改
	if req.UserPassword != "" {
		req.UserPassword = libUtils.EncryptPassword(req.UserPassword, user.UserSalt)
	}

	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 更新用户信息
			_, err = dao.User.Ctx(ctx).TX(tx).WherePri(req.Id).Data(req).Update()
			// 设置用户与角色关联信息
			_ = s.EditRoles(ctx, req.RoleIds, req.Id)
			liberr.IsNil(ctx, err, "设置用户权限失败")
			// 添加员工与角色关联表
			err = s.AddUserRole(ctx, req.RoleIds, gconv.Int(req.Id))
			liberr.IsNil(ctx, err, "设置用户岗位信息失败")
		})
		return nil
	})
	return
}

func (s *sUser) Delete(ctx context.Context, id []int) (err error) {
	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 删除用户
			_, err = dao.User.Ctx(ctx).TX(tx).WhereIn(dao.User.Columns().Id, id).Delete()
			liberr.IsNil(ctx, err, "删除用户失败")
			// 删除用户与岗位关联表
			_, err = dao.UserPost.Ctx(ctx).TX(tx).Where(dao.UserPost.Columns().UserId, id).Delete()
			liberr.IsNil(ctx, err, "删除用户与角色关联表失败")

			enforcer, e := casbin.CabinEnforcer(ctx)
			liberr.IsNil(ctx, e, "获取权限失败")

			for _, v := range id {
				_, _ = enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.CasBinUserPrefix, v))
			}
		})
		return err
	})
	return
}

func (s *sUser) List(ctx context.Context, req *admin.UserListReq) (res *admin.UserListRes, err error) {
	res = new(admin.UserListRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.User.Ctx(ctx)
		if req.UserName != "" {
			m = m.WhereLike(dao.User.Columns().UserName, "%"+req.UserName+"%")
		}
		if req.Mobile != "" {
			m = m.WhereLike(dao.User.Columns().Mobile, "%"+req.Mobile+"%")
		}
		if req.UserStatus > 0 {
			m = m.Where(dao.User.Columns().UserStatus, req.UserStatus)
		}
		if req.UserNickName != "" {
			m = m.WhereLike(dao.User.Columns().UserNickname, "%"+req.UserNickName+"%")
		}
		if req.DeptId > 0 {
			m = m.Where(dao.User.Columns().DeptId, req.DeptId)
		}

		res.Total, err = m.Count()
		liberr.IsNil(ctx, err, "获取用户列表失败")
		res.CurrentPage = req.PageNum
		err = m.Page(req.PageNum, req.PageSize).Scan(&res.List)
		liberr.IsNil(ctx, err, "获取用户列表失败")
	})
	return
}

func (s *sUser) AddUserRole(ctx context.Context, roleIds []int, userId int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := casbin.CabinEnforcer(ctx)
		liberr.IsNil(ctx, e)
		for _, v := range roleIds {
			_, e = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", s.CasBinUserPrefix, userId), gconv.String(v))
			liberr.IsNil(ctx, e)
		}
	})
	return err
}

func (s *sUser) EditRoles(ctx context.Context, rolesId []int, userId int) (err error) {
	// 实例化casbin
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := casbin.CabinEnforcer(ctx)
		liberr.IsNil(ctx, e)
		// 删除用户所有角色
		_, e = enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.CasBinUserPrefix, userId))
		// 添加用户角色关联
		for _, v := range rolesId {
			_, e = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", s.CasBinUserPrefix, userId), gconv.String(v))
			liberr.IsNil(ctx, e)
		}
	})
	return
}

// AddUserPost 添加员工与岗位关联表
func (s *sUser) AddUserPost(ctx context.Context, tx *gdb.TX, postIds []int, userId int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 删除员工与岗位关联表
		_, err = dao.UserPost.Ctx(ctx).TX(tx).Where(dao.UserPost.Columns().UserId, userId).Delete()
		liberr.IsNil(ctx, err, "设置用户岗位失败")
		// 添加员工与岗位关联表
		if len(postIds) == 0 {
			return
		}
		data := g.List{}
		for _, v := range postIds {
			data = append(data, g.Map{
				dao.UserPost.Columns().UserId: userId,
				dao.UserPost.Columns().PostId: v,
			})
		}
		_, err = dao.UserPost.Ctx(ctx).TX(tx).Data(data).Insert()
		liberr.IsNil(ctx, err, "设置用户岗位失败")
	})
	return err
}

// userNameOrMobileExists 检测手机号和用户名称是否已存在
func (s *sUser) userNameOrMobileExists(ctx context.Context, userName string, mobile string, id ...int64) error {
	user := (*entity.User)(nil)
	err := g.Try(ctx, func(ctx context.Context) {
		m := dao.User.Ctx(ctx)
		if len(id) > 0 {
			m = m.Where(dao.User.Columns().Id+"!= ?", id[0])
		}
		err := m.Wheref("%s = '%s' OR %s = '%s'", dao.User.Columns().UserName, userName, dao.User.Columns().Mobile, mobile).Scan(&user)
		liberr.IsNil(ctx, err, "获取用户信息失败")
		if user == nil {
			return // 用户不存在
		}
		if user.UserName == userName {
			liberr.IsNil(ctx, gerror.New("用户名已存在"))
		}
		if user.Mobile == mobile {
			liberr.IsNil(ctx, gerror.New("手机号已存在"))
		}
	})
	return err
}
