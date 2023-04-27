package main

import (
	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/internel/localization"
	"github.com/willoong9559/gin-mall/internel/routers"
)

// @title gin-mall
// @version 1.0
// @description This is a sample server  :) .
// @termsOfService https://github.com/willoong9559/gin-mall

// @contact.name   API Support
// @contact.email  willoongl@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	conf.Init()
	localization.LoadValidTrans()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}
