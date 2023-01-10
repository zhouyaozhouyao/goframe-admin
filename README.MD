# goframe-admin

Goravel 是一个基础功能完善的、具有良好扩展能力的 Web 应用程序框架。 作为一个起始脚手架帮助 Golang 开发者快速构建自己的应用。

# 主要功能

- [x] Config
- [x] Http
- [x] Orm
- [x] Logger
- [x] Cache
- [x] Queue
- [x] Casbin
- [x] WeChat
- [x] AliSms
- [x] Oss
- [ ] Schedule
- [ ] Event
- [ ] Mail
- [] Migrate

# 感谢

框架使用的是 `gofrmae/v2` 为基础进行开发
Wechat 是参考 `PowerWeChat` 进行改版 方便自已持续集成


# 接口文档

gf run main.go 启动项目后访问：

http://127.0.0.1:8000/swagger/

# 说明

main分支使用目前goframe最新版V2.2开发实现

# 运行流程

## 1. 下载项目

git@github.com:zhouyaozhouyao/goframe-admin.git

## 2. 配置数据库

把resource/databases/admin.sql导入你的数据库中

## 3. 修改配置文件

修改hack/config.yaml文件中的数据库密码

修改manifest/config/config.yaml中的数据库密码

redis的密码可以不改，gtoken已经使用gcache模式，如果你需要使用redis，请配置配置文件中的redis

oss需要进行配置或上传改为本地文件，不影响项目启动，如果你需要图片上传功能，请修改配置文件中qiniu相关的参数

## 4. 启动项目

在项目根目录下执行：

go run main.go

如果你需要自动编译，可以执行：

gf run main.go

# 项目启动失败可能的原因

1. Go或者GoFrame安装的版本不一致
2. 配置文件问题，密码不正确等等

# 出现问题可以联系我


## 微信

zhouyao999


# 交叉编译

## for Linux

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

## for windows

```
CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build
```

# 热更新启动项目

```
gf run main.go
```
