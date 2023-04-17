package main

import (
	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/routers"
)

// @title 商城系统
// @version 1.0
// @description gin-mall
// @termsOfService https://github.com/willoong9559/gin-mall
func main() {
	conf.Init()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}
