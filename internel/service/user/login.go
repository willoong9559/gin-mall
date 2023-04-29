package user

import (
	"github.com/gin-gonic/gin"

	"github.com/willoong9559/gin-mall/internel/dao"
	"github.com/willoong9559/gin-mall/internel/model"
	e "github.com/willoong9559/gin-mall/pkg/errcode"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/serializer"
)

type UserLoginService struct {
	UserService
	CaptchaCode string `form:"captcha_code" json:"captcha_code" binding:"required,len=4"`
}

func (service *UserLoginService) Login(ctx *gin.Context) *serializer.Response {
	var user *model.User
	if !utils.CaptchaVerify(ctx, service.CaptchaCode) {
		return serializer.GetResponse(e.ErrorCaptcha, "")
	}
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		return serializer.GetResponse(e.ErrorDatabase, "")
	}
	if !exist {
		return serializer.GetResponse(e.ErrorExistUserNotFound, "")
	}
	if !user.CheckPassword(service.Password) {
		return serializer.GetResponse(e.ErrorNotCompare, "")
	}
	// http 无状态(认证token签发)
	token, err := utils.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		return serializer.GetResponse(e.ErrorAuthToken, "")
	}
	tokenData := serializer.TokenData{
		User:  serializer.BuildUser(user),
		Token: token,
	}
	return serializer.GetResponse(e.SUCCESS, tokenData)
}
