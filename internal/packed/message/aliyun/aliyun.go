package aliyun

import (
	"api/internal/model"
	"api/internal/packed/message/dto"
	"context"
	"encoding/json"
	"os"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/alibabacloud-go/tea/tea"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
)

type AliYun struct {
}

func (a *AliYun) SendMessage(ctx context.Context, data *model.SmsMessageInput) (smsResponse dto.SmsResponse, err error) {
	// 读取配置文件
	aliyunCfg := g.Cfg().MustGet(ctx, "ali").MapStrStr()
	client, err := a.client(tea.String(aliyunCfg["accessKey"]), tea.String(aliyunCfg["secret"]))
	if err != nil {
		return dto.SmsResponse{}, err
	}

	var mapJson = make(map[string]string)
	mapJson["code"] = data.Code
	param, _ := json.Marshal(mapJson)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(data.Phone),
		TemplateCode:  tea.String(data.TemplateCode),
		TemplateParam: tea.String(string(param)),
		SignName:      tea.String("北京快乐就好科技"),
	}

	sms, err := client.SendSms(sendSmsRequest)
	os.Exit(0)
	return dto.SmsResponse{
		Status: dto.SmsStatus{
			Phone:   "phone",
			Code:    *sms.Body.Code,
			Message: *sms.Body.Message,
		},
		RequestId: *sms.Body.RequestId,
	}, err

}

func (a *AliYun) client(accessKeyId *string, accessKeySecret *string) (result *dysmsapi20170525.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	result = &dysmsapi20170525.Client{}
	result, err = dysmsapi20170525.NewClient(config)
	return result, err
}
