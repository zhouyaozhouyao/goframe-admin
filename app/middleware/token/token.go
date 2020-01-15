/**
 * 校验 token是否有效 中间件
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/9 4:58 下午
 */

// Package token 检测token是否有效中间件
package token

import (
	"gadmin/app/api/admin/login"
	"gadmin/library/base"
	"gadmin/library/e"
	"gadmin/library/redis"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// Validator 验证token有效性
func Validator(r *ghttp.Request) {
	login.GfJWTMiddleware.MiddlewareFunc()(r)
	// 解析token
	parseToken, _ := login.GfJWTMiddleware.ParseToken(r)
	var token = parseToken.Raw
	var claims = gconv.Map(parseToken.Claims)
	r.SetParam("username", claims["username"])
	if !GetRedisToken(gconv.String(claims["uuid"]), token) {
		base.Fail(r, e.ErrorAuthCheckTokenFail)
	}
	r.Middleware.Next()
}

// GetRedisToken 获取缓存中的token与客户端token对比
func GetRedisToken(uuid string, oldToken string) bool {
	redisPrefix := gconv.String(g.Cfg("redis").Get("APP.LOGIN_PREFIX"))
	key := redisPrefix + uuid
	if redis.Get(key) != oldToken {
		return false
	}
	return gconv.Bool(true)
}
