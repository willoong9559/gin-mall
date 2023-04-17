# gin-mall
电子商城系统
----
实现一些如用户登录，校验，邮箱绑定，购物支付等功能，使用MySQL读写分离、redis数据缓存等。
## 主要依赖：
- gin
- mysql
- redis
- ini
- gorm
- swaggo
- jwt-go
- dchest/captcha

## 项目结构
```
gin-mall/
├── api
├── cache
├── conf
├── dao
├── doc
├── middleware
├── model
├── pkg
│  ├── e
│  └── util
├── routers
├── serializer
└── service
```

## 简要接口文档(swaggo生成)
![已实现的api](https://raw.githubusercontent.com/willoong9559/gin-mall/main/docs/api.PNG)