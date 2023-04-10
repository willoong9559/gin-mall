package main

import (
	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/routers"
)

func main() {
	conf.Init()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}
