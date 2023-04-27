package user

import (
	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/internel/dao"
	"github.com/willoong9559/gin-mall/internel/model"
	e "github.com/willoong9559/gin-mall/pkg/errcode"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/serializer"
)

type UserRegisterService struct {
	UserService
	NickName    string `form:"nick_name" json:"nick_name" binding:"required,min=5,max=8"`
	RePassword  string `form:"re_password" json:"re_password" binding:"required,eqfield=Password"`
	CaptchaCode string `form:"captcha_code" json:"captcha_code" binding:"required,len=4"`
}

func (service *UserRegisterService) Register(ctx *gin.Context) *serializer.Response {
	var user *model.User
	if !utils.CaptchaVerify(ctx, service.CaptchaCode) {
		return serializer.GetResponse(e.ErrorCaptcha, "")
	}
	// 设置加密秘钥
	utils.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		return serializer.GetResponse(e.ErrorDatabase, "")
	}
	if exist {
		return serializer.GetResponse(e.ErrorExistUser, "")
	}
	user = &model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    utils.Encrypt.AesEncoding("10000"), // 初始金额
	}
	// 密码加密
	if err = user.SetPassword(service.Password); err != nil {
		return serializer.GetResponse(e.ErrorFailEncryption, "")
	}
	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		return serializer.GetResponse(e.ErrorDatabase, "")
	}
	return serializer.GetResponse(e.SUCCESS, "")
}
