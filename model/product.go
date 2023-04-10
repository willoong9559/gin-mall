package model

import (
	"strconv"

	"github.com/willoong9559/gin-mall/cache"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint   `gorm:"not null"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

// View 获取点击数
func (product *Product) View() int64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseInt(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
