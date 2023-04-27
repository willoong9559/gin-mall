package user

import (
	"context"

	"github.com/willoong9559/gin-mall/internel/dao"
	e "github.com/willoong9559/gin-mall/pkg/errcode"
	"github.com/willoong9559/gin-mall/serializer"
)

type UserUpdateService struct {
	UserService
	NickName    string `form:"nick_name" json:"nick_name" binding:"required,min=5,max=8"`
	CaptchaCode string `form:"captcha_code" json:"captcha_code" binding:"required,len=4"`
}

func (service *UserUpdateService) Update(ctx context.Context, uId uint) serializer.Response {
	var err error
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if service.NickName != "" {
		user.NickName = service.NickName
	}

	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}
