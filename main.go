package main

import (
	"log"

	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/global"
	"github.com/willoong9559/gin-mall/internel/dao"
	"github.com/willoong9559/gin-mall/internel/localization"
	"github.com/willoong9559/gin-mall/internel/routers"
	"github.com/willoong9559/gin-mall/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
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
	setupLogger()
	go setupGlobalDB()
	go localization.LoadValidTrans()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}

func setupLogger() {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  conf.LogSavePath + "/" + conf.LogFileName + conf.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
}

func setupGlobalDB() {
	global.DB = dao.NewGlobalDB()
}

func dbMigration() {
	err := dao.Migration()
	if err != nil {
		global.Logger.Fatalf("register table fail: %s", err)
		return
	}
	global.Logger.Info("register table success")
}
