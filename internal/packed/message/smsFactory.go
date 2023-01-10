package message

import (
	"api/internal/model"
	"api/internal/packed/message/aliyun"
	"api/internal/packed/message/dto"
	"api/internal/packed/message/wechat_template"
	"context"
)

type Message interface {
	// SendMessage 发送消息
	SendMessage(ctx context.Context, data *model.SmsMessageInput) (smsResponse dto.SmsResponse, err error)
}

func New(name string) Message {
	if name == dto.AliYun {
		return &aliyun.AliYun{}
	}

	if name == dto.WechatTemplate {
		return &wechat_template.WechatTemplate{}
	}

	return nil
}
