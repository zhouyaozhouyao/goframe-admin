package auth

import (
	"api/internal/packed/wechat/kernel"
	"context"
)

func RegisterProvider(ctx context.Context, app kernel.ApplicationInterface) (*AccessToken, error) {
	return NewAccessToken(ctx, &app)
}

func RegisterAuthProvider(ctx context.Context, app kernel.ApplicationInterface) (*Client, error) {
	return NewClient(ctx, &app)
}
