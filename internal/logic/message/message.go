package message

import (
	"api/api/v1/common"
	"api/internal/library/libUtils"
	"api/internal/library/libcache"
	"api/internal/model"
	"api/internal/modules/admin/consts"
	"api/internal/packed/message"
	"api/internal/service"
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type sMessage struct {
}

func init() {
	service.RegisterMessage(New())
}

func New() *sMessage {
	return &sMessage{}
}

// SendMessage 发送短信验证码
func (s *sMessage) SendMessage(ctx context.Context, req *common.SendMessageReq) (boolean bool, err error) {
	code := libUtils.SmsRandCode(6)
	resp, err := message.New("aliyun").SendMessage(ctx, &model.SmsMessageInput{
		Code:         code,
		Phone:        req.Phone,
		TemplateCode: "SMS_264885777",
	})
	if err != nil {
		return false, err
	}

	if resp.Status.Code != consts.SmsOk {
		g.Log().Error(ctx, "发送短信失败", resp.Status.Message, g.Map{"RequestId": resp.RequestId})
		return false, err
	}

	prefix := g.Cfg().MustGet(ctx, "admin.cache.prefix").String()
	// 短信通过redis进行缓存，方便后续验证
	libcache.NewRedis(prefix).Set(ctx, "sms_"+req.Phone, code, 10*time.Minute)
	return boolean, err
}
