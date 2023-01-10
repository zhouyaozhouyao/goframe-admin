// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"api/internal/model"
	"context"
)

type ILoginLog interface {
	Invoke(ctx context.Context, data *model.LoginLogInput)
}

var localLoginLog ILoginLog

func LoginLog() ILoginLog {
	if localLoginLog == nil {
		panic("implement not found for interface ILoginLog, forgot register?")
	}
	return localLoginLog
}

func RegisterLoginLog(i ILoginLog) {
	localLoginLog = i
}