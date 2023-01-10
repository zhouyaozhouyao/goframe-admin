package phoneNumber

import (
	"api/internal/packed/wechat/kernel"
	"context"
)

func RegisterProvider(ctx context.Context, app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(ctx, &app, nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		baseClient,
		ctx,
	}, nil
}
