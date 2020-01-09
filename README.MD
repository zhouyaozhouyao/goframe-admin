### 当前项目会长期更新当作学习记录，有需求的小伙伴可以自行下载，也可以参考以下项目，其他人开发的管理后台，也是基于GF

gmanager: 基于gf框架的管理平台，支持登录、认证、组织机构、用户、角色、菜单、日志 
https://gitee.com/goflyfox/gmanager



### 简要说明
登录账号信息

用户名：admin

密码：123456

项目数据库结构在[docfile/sql](https://github.com/zhouyaozhouyao/goframe-admin/tree/master/docker)目录下

### 前端界面搭建中




自述：第一次设计，整体结构可能有些问题，望见谅，要是能给我个建议什么的就更好了。暂时也没想到增加什么功能，有需要可以提个issues，看看我能不能做。



----

### 技术栈


#### 后端

框架：[GoFrame](https://github.com/gogf/gf)

登录加密：RSA

API授权方式：JWT，使用[gf-jwt](https://github.com/gogf/gf-jwt)

----

#### 前端
框架：[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)

----
----
### 如何访问接口

```
 gf run main.go
```

```
  SERVER  | DOMAIN  | ADDRESS | METHOD |         ROUTE         |                           HANDLER                           |   MIDDLEWARE     
|---------|---------|---------|--------|-----------------------|-------------------------------------------------------------|-----------------|
  default | default | :8199   | POST   | /v1/admin/loginSubmit | github.com/gogf/gf-jwt.(*GfJWTMiddleware).LoginHandler-fm   |                  
|---------|---------|---------|--------|-----------------------|-------------------------------------------------------------|-----------------|
  default | default | :8199   | GET    | /v1/admin/refresh     | github.com/gogf/gf-jwt.(*GfJWTMiddleware).RefreshHandler-fm | token.Validator  
|---------|---------|---------|--------|-----------------------|-------------------------------------------------------------|-----------------|
  default | default | :8199   | GET    | /v1/admin/user        | gadmin/router.init.0.func1.1.1                              | token.Validator  
|---------|---------|---------|--------|-----------------------|-------------------------------------------------------------|-----------------|
```

使用postman发送post请求


