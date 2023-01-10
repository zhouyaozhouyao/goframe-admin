package wechat

type Configuration struct {
	MiniProgram MiniProgramCfg // 小程序配置文件
}

type MiniProgramCfg struct {
	AppID         string
	Secret        string
	RedisAddr     string
	MessageToken  string
	MessageAesKey string
}
