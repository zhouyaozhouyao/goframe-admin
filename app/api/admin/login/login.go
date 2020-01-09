/**
 * 后台用户登录管理
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/9 2:59 下午
 */

// Package login 后台用户登录管理
package login

import (
	"errors"
	"gadmin/app/model/users"
	"gadmin/library/base"
	"gadmin/library/helper"
	"gadmin/library/input"
	"gadmin/library/redis"
	"time"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gvalid"

	"github.com/gogf/gf/net/ghttp"

	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/os/glog"
)

var (
	GfJWTMiddleware *jwt.GfJWTMiddleware // 声明jwt包的全局变量
)

type SignRequest struct {
	Username string `v:required#账号不能为空 json:"username"`
	Password string `v:required#密码不能为空 json:"password"`
}

func init() {
	authMiddleWare, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "zhouyao",
		Key:             []byte("zhouyao"),
		Timeout:         time.Minute * 60 * 24 * 30,
		MaxRefresh:      time.Minute * 60 * 24 * 30,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   PostLogin,
		RefreshResponse: RefreshResponse,
		Unauthorized:    Unauthorized,
		IdentityHandler: IdentityHandler,
		PayloadFunc:     PayloadFunc,
	})

	if err != nil {
		glog.Error("JWT Error:" + err.Error())
	}

	GfJWTMiddleware = authMiddleWare
}

// Authenticator 检测身份信息是否正常
func Authenticator(r *ghttp.Request) (interface{}, error) {
	var req *SignRequest
	// 接收参数
	input.JSONToStruct(r, &req)

	// 校验数据参数
	if err := gvalid.CheckStruct(req, nil); err != nil {
		base.FailParam(r, err.String())
	}

	// 查询数据
	res := users.GetOne(g.Map{"username": req.Username})

	if res.Id <= 0 {
		return nil, errors.New("用户名或密码错误")
	}

	reqPwd, errPwd := gmd5.Encrypt(req.Password + res.Salt)
	if errPwd != nil {
		glog.Error("md5加密异常", errPwd)
		return nil, errors.New("服务器异常")
	}

	if reqPwd != res.Password {
		return nil, errors.New("用户名或密码错误~")
	}

	// 设置参数保存到请求中
	r.SetParam("uuid", res.Uuid)

	return g.Map{"username": res.Username, "uuid": res.Uuid}, nil
}

// PostLogin 返回对应的用户信息
func PostLogin(r *ghttp.Request, code int, token string, expire time.Time) {
	j, _ := r.GetJson()
	// 格式化时间
	t := helper.TimeToString(expire)

	// 获取配置文件中的redis前缀
	var loginPrefix = g.Cfg("redis").Get("APP.LOGIN_PREFIX")
	redis.Set(gconv.String(loginPrefix)+gconv.String(r.GetParam("uuid")), token)

	base.Success(r, g.Map{
		"username": j.GetString("username"),
		"token":    token,
		"expire":   t,
	})
}

// RefreshResponse 刷新token信息
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	base.Success(r, g.Map{"token": token, "expire": helper.TimeToString(expire)})
}

// Unauthorized 返回验证错误信息
func Unauthorized(r *ghttp.Request, code int, message string) {
	// TODO 英文提示可在此处转换为中文，暂没找到好的方式后续单做一个配置文件
	base.FailParam(r, message)
}

// IdentityHandler sets the identity for JWT.
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["id"]
}

// PayloadFunc 给token中添加其它字段的数据
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}
