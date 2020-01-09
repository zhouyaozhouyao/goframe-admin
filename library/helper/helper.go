/**
 * 基础的方法
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/7 5:21 下午
 */

package helper

import (
	"time"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

// PageParam 检测分页参数，给定默认值与limit最大值
func PageParam(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	if limit >= 50 {
		limit = 50
	}

	return page, limit
}

// TimeToString 把系统时间转换为 2020-01-12 00:00:00 形式
func TimeToString(expire time.Time) string {
	t, err := gtime.StrToTime(expire.Format(time.RFC3339))
	if err != nil {
		glog.Error("服务器内部错误", err)
	}

	return gconv.String(t)
}
