package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/willoong9559/gin-mall/internel/model"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, basePage model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).Offset((basePage.PageNum - 1) * basePage.PageSize).Limit(basePage.PageSize).Find(&products).Error
	return
}

func (dao *ProductDao) SearchProducts(info string, basePage model.BasePage) (products []*model.Product, count int64, err error) {
	errC := dao.DB.Model(&model.Product{}).Where("name LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").Count(&count).Error
	err = dao.Model(&model.Product{}).
		Offset((basePage.PageNum - 1) * basePage.PageSize).
		Limit(basePage.PageSize).Find(&products).Error
	if err != nil {
		err = errC
	}
	return
}
