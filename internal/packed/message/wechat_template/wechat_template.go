package wechat_template

import (
	"api/internal/model"
	"api/internal/packed/message/dto"
	"context"
	"fmt"
)

type WechatTemplate struct {
}

func (w *WechatTemplate) SendMessage(ctx context.Context, input *model.SmsMessageInput) (smsResponse dto.SmsResponse, err error) {
	fmt.Println("我发送微信模板消息")
	return smsResponse, err

}
