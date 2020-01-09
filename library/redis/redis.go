/**
 * Redis 客户端操作使用
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/3 2:30 下午
 */

package redis

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// 定义全局变量
var (
	redis = g.Redis()
)

// Set 设置字符串
func Set(key string, value string) bool {
	if ok, err := redis.Do("SET", key, value); err != nil {
		panic(err)
	} else {
		return gconv.Bool(ok)
	}
}

// Get 获取字符串
func Get(key string) string {
	if data, err := redis.Do("GET", key); err != nil {
		panic(err)
	} else {
		return gconv.String(data)
	}
}
