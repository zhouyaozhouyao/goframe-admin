package phoneNumber

import (
	"api/internal/packed/wechat/kernel"
	"api/internal/packed/wechat/miniProgram/phoneNumber/response"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type Client struct {
	*kernel.BaseClient
	Ctx context.Context
}

func (c *Client) GetUserPhoneNumber(code string) (*response.PhoneNumberResponse, error) {
	result := &response.PhoneNumberResponse{}
	_, _ = c.HttpPost("/wxa/business/getuserphonenumber", g.Map{"code": code}, result)
	return result, nil
}
