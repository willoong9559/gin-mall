package global

import (
	"log"

	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func setupLogger() error {
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  conf.LogSavePath + "/" + conf.LogFileName + conf.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
