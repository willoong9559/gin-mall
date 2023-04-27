package service

import (
	"context"
	"mime/multipart"
	"strconv"
	"sync"

	"github.com/willoong9559/gin-mall/internel/dao"
	"github.com/willoong9559/gin-mall/internel/model"
	e "github.com/willoong9559/gin-mall/pkg/errcode"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/serializer"
)

// 更新商品的服务
type ProductService struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" binding:"omitempty,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"`
	Num           int    `form:"num" json:"num"`
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(uId)
	if err != nil {
		utils.LogrusObj.Info(err)
	}
	// 以第一张作为封面图
	tmp, _ := files[0].Open()
	path, err := utils.UploadToLocalStatic(tmp, uId, service.Name, utils.ProductImg)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorProductUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryID:    uint(service.CategoryID),
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		Num:           service.Num,
		OnSale:        true,
		BossID:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(len(files))
	for i, file := range files {
		num := strconv.Itoa(i)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = utils.UploadToLocalStatic(tmp, uId, service.Name+num, utils.ProductImg)
		if err != nil {
			code = e.ErrorProductUpload
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		wg.Done()
	}
	wg.Wait()

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

func (service *ProductService) List(ctx context.Context) serializer.Response {
	var products []*model.Product
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.CategoryID != 0 {
		condition["categorg_id"] = service.CategoryID
	}

	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	productDao = dao.NewProductDaoByDB(productDao.DB)
	products, _ = productDao.ListProductByCondition(condition, service.BasePage)

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

func (service *ProductService) Search(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewProductDao(ctx)
	products, count, err := productDao.SearchProducts(service.Info, service.BasePage)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(count))
}
