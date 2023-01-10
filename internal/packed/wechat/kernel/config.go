package kernel

import "github.com/gogf/gf/v2/frame/g"

type Config struct {
	*g.Map
}

func NewConfig(item *g.Map) *Config {
	return &Config{
		item,
	}
}
