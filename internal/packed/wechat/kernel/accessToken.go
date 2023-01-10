package kernel

import (
	"api/internal/packed/wechat/kernel/contract"
	response2 "api/internal/packed/wechat/kernel/response"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/crypto/gmd5"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

type AccessToken struct {
	App      *ApplicationInterface
	Token    *g.Map
	TokenKey string
	Ctx      context.Context

	EndpointToGetToken string

	*InteractWithCache

	GetCredentials func() *g.MapStrStr
	GetEndpoint    func() (string, error)
}

func NewAccessToken(ctx context.Context, app *ApplicationInterface) (*AccessToken, error) {
	// 读取配置
	config := (*app).GetContainer().GetConfig()
	//cacheClient := (*config)["cache"].(libcache.ICache)
	var cacheClient CacheInterface = nil
	if (*config)["cache"] != nil {
		cacheClient = (*config)["cache"].(CacheInterface)
	}
	token := &AccessToken{
		App:               app,
		Token:             nil,
		Ctx:               ctx,
		TokenKey:          "access_token",
		InteractWithCache: NewInteractsWithCache(cacheClient, "powerwechat.access_token."),
	}
	token.OverrideGetEndpoint()
	return token, nil
}

func (accessToken *AccessToken) ApplyToRequest(r *http.Request, requestOptions *g.Map) (*http.Request, error) {
	key := accessToken.TokenKey
	resToken, err := accessToken.GetToken(false)

	if err != nil {
		return nil, err
	}

	arrayReturn := &g.Map{
		key: resToken.AccessToken,
	}

	q := r.URL.Query()
	for k, value := range *arrayReturn {
		q.Set(k, gconv.String(value))
	}
	r.URL.RawQuery = q.Encode()
	return r, nil
}

// GetToken 获取token
func (accessToken *AccessToken) GetToken(refresh bool) (resToken *response2.TokenResponse, err error) {
	// 获取缓存key
	cacheKey := accessToken.GetCacheKey()
	cache := accessToken.Cache
	boolean, _ := cache.Has(accessToken.Ctx, cacheKey)
	if !refresh && boolean {
		// 读取缓存
		value, err := cache.Get(accessToken.Ctx, cacheKey)
		if err == nil && value != nil {
			token := value.Map()
			resToken = &response2.TokenResponse{
				ExpiresIn: gconv.Float64(token["expires_in"]),
			}
			// 小程序的token
			if accessToken.TokenKey == "access_token" && token["access_token"] != nil {
				resToken.AccessToken = token["access_token"].(string)
			} else if accessToken.TokenKey == "authorizer_access_token" && token["authorizer_access_token"] != nil {
				resToken.AccessToken = token[accessToken.TokenKey].(string)
				resToken.AuthorizerAccessToken = token[accessToken.TokenKey].(string)
			} else {
				return nil, gerror.New("no token found in cache")
			}
			return resToken, nil
		}
	}

	// 请求微信接口获取token
	resToken, err = accessToken.requestToken(accessToken.GetCredentials())
	if err != nil {
		return nil, err
	}
	// 缓存 token
	_, err = accessToken.SetToken(resToken)

	return resToken, nil
}

func (accessToken *AccessToken) requestToken(credentials *g.MapStrStr) (*response2.TokenResponse, error) {
	token, err := accessToken.sendRequest(credentials)
	if token == nil || token.AccessToken == "" {
		return nil, gerror.Newf("request access_token fail: %v", err)
	}
	return token, err
}

func (accessToken *AccessToken) Refresh() contract.AccessTokenInterface {
	_, _ = accessToken.GetToken(true)
	return accessToken
}

func (accessToken *AccessToken) sendRequest(credential *g.MapStrStr) (*response2.TokenResponse, error) {
	strEndpoint, err := accessToken.GetEndpoint()
	if err != nil {
		return nil, err
	}

	response := &response2.TokenResponse{}
	// 发送请求
	err = g.Client().GetVar(accessToken.Ctx, strEndpoint, credential).Scan(&response)
	return response, err
}

func (accessToken *AccessToken) SetToken(token *response2.TokenResponse) (tokenInterface contract.AccessTokenInterface, err error) {
	// 检测 token 是否过期
	if token.ExpiresIn <= 0 {
		token.ExpiresIn = 7200
	}
	// 开启缓存 根据入口文件的缓存来自动区分是cache还是redis
	cache := accessToken.GetCache()
	err = cache.Set(accessToken.Ctx, accessToken.GetCacheKey(), gconv.Map(token), time.Duration(token.ExpiresIn)*time.Second)
	if err != nil {
		return nil, err
	}

	ok, err2 := cache.Has(accessToken.Ctx, accessToken.GetCacheKey())
	if err2 != nil || !ok {
		return nil, gerror.New("failed to cache access token")
	}
	return accessToken, nil
}

func (accessToken *AccessToken) OverrideGetEndpoint() {
	accessToken.GetEndpoint = func() (string, error) {
		if accessToken.EndpointToGetToken == "" {
			return "", gerror.New("endpoint is empty")
		}
		return accessToken.EndpointToGetToken, nil
	}
}

// GetCacheKey 缓存的Key生成规则
func (accessToken *AccessToken) GetCacheKey() string {
	credentials := *accessToken.GetCredentials()
	data := fmt.Sprintf("%s%s", credentials["appid"], credentials["secret"])
	return gmd5.MustEncryptString(data)
}
