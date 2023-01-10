package liberr

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// IsNil 判断是否为空
func IsNil(ctx context.Context, err error, msg ...string) {
	if !g.IsNil(err) {
		if len(msg) > 0 {
			g.Log().Error(ctx, err.Error())
			panic(msg[0])
		} else {
			panic(err.Error())
		}
	}
}

// ValueIsNil 判断值
func ValueIsNil(value interface{}, msg string) {
	if g.IsNil(value) {
		panic(msg)
	}
}
