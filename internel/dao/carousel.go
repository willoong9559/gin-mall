package dao

import (
	"context"

	"github.com/willoong9559/gin-mall/internel/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

func (dao *CarouselDao) ListAddress() (carousels []*model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return
}
