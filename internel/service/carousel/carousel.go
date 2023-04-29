package service

import (
	"context"

	"github.com/willoong9559/gin-mall/global"
	"github.com/willoong9559/gin-mall/internel/dao"
	e "github.com/willoong9559/gin-mall/pkg/errcode"
	"github.com/willoong9559/gin-mall/serializer"
)

type ListCarouselsService struct{}

func (service *ListCarouselsService) List(ctx context.Context) serializer.Response {
	carouselsCtx := dao.NewCarouselDao(context.Background())
	carousels, err := carouselsCtx.ListAddress()
	if err != nil {
		global.Logger.Error(err)
		return *serializer.GetResponse(e.ErrorDatabase, "")
	}
	return serializer.GetListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
