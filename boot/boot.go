package boot

import "github.com/gogf/gf/frame/g"

func init() {
	// 增加自定义配置文件
	g.Cfg("redis").SetFileName("redis.toml")
	g.Cfg("message").SetFileName("message.toml")
}
