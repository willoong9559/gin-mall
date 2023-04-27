package main

import (
	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/internel/localization"
	"github.com/willoong9559/gin-mall/internel/routers"
)

// @title 商城系统
// @version 1.0
// @description gin-mall
// @termsOfService https://github.com/willoong9559/gin-mall
func main() {
	conf.Init()
	localization.LoadValidTrans()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}
