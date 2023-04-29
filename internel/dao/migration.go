package dao

import (
	"github.com/willoong9559/gin-mall/global"
	"github.com/willoong9559/gin-mall/internel/model"
)

// Migration 执行数据迁移
func Migration() error {
	//自动迁移模式
	err := global.DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Product{},
			&model.Carousel{},
			&model.Category{},
			&model.Favorite{},
			&model.ProductImg{},
			&model.Order{},
			&model.Cart{},
			&model.Admin{},
			&model.Address{},
			&model.Notice{},
		)
	return err
}
