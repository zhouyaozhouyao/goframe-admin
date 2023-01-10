package kernel

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type ApplicationInterface interface {
	GetContainer() *ServiceContainer
	GetConfig() *Config
	GetAccessToken() *AccessToken
}

type ServiceContainer struct {
	UserConfig *g.Map
	Config     *g.Map
	Ctx        context.Context
}

func NewServiceContainer(ctx context.Context, config *g.Map) (*ServiceContainer, error) {
	return &ServiceContainer{
		UserConfig: config,
		Ctx:        ctx,
	}, nil
}

func (container *ServiceContainer) getConfig() *g.Map {
	return &g.Map{
		"http": &g.Map{
			"timeout":  30,
			"base_uri": "https://api.weixin.qq.com",
		},
	}
}

func (container *ServiceContainer) GetConfig() *g.Map {
	basicConfig := container.getConfig()
	container.Config = ReplaceHashMapRecursive(container.Config, basicConfig, container.UserConfig)
	return container.Config
}

func ReplaceHashMapRecursive(toMap *g.Map, subMaps ...*g.Map) *g.Map {
	if toMap == nil {
		toMap = &g.Map{}
	}
	// 拍平subMaps
	for _, subMap := range subMaps {
		if subMap != nil {
			for k, v := range *subMap {
				m := gconv.Map(v)
				if len(m) > 0 {
					for kk, vv := range m {
						(*toMap)[kk] = vv
					}
				} else {
					(*toMap)[k] = v
				}
			}
		}

	}
	return toMap
}
