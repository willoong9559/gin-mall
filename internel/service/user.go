package service

import (
	"context"
	"mime/multipart"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/internel/dao"
	"github.com/willoong9559/gin-mall/internel/model"
	e "github.com/willoong9559/gin-mall/pkg/errcode"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/serializer"
	"gopkg.in/mail.v2"
)

// UserService 管理用户服务
type UserService struct {
	NickName    string `form:"nick_name" json:"nick_name"`
	UserName    string `form:"user_name" json:"user_name"`
	Password    string `form:"password" json:"password"`
	Key         string `form:"key" json:"key"` // 前端进行判断
	CaptchaCode string `form:"captcha_code" json:"captcha_code"`
}

type SendEmailService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (service *UserService) Register(ctx *gin.Context) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	if !utils.CaptchaVerify(ctx, service.CaptchaCode) {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "验证码错误",
		}
	}
	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密钥长度不足",
		}
	}
	// 10000 => 密文加密
	utils.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
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
		utils.LogrusObj.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Login(ctx *gin.Context) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	if !utils.CaptchaVerify(ctx, service.CaptchaCode) {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "验证码错误",
		}
	}
	userDao := dao.NewUserDao(ctx)
	// 判断用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist || err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册",
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新输入",
		}
	}
	// http 无状态(认证token签发)
	token, err := utils.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	tokenData := serializer.TokenData{
		User:  serializer.BuildUser(user),
		Token: token,
	}
	return serializer.Response{
		Status: code,
		Data:   tokenData,
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var err error
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if service.NickName != "" {
		user.NickName = service.NickName
	}

	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.SUCCESS
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 保存图片到本地
	path, err := utils.UploadToLocalStatic(file, uId, user.UserName, utils.Avatar)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

// Send 发送邮件
func (service *SendEmailService) Send(ctx context.Context, id uint) serializer.Response {
	code := e.SUCCESS
	address := "" // 验证地址
	var notice *model.Notice
	token, err := utils.GenerateEmailToken(id, service.OperationType, service.Email, service.Password)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = conf.ValidEmail + token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	// gopkg.in/mail.v2
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "邮箱验证")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Valid 验证邮箱
func (service ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS

	//验证token
	if token == "" {
		code = e.InvalidParams
	}
	claims, err := utils.ParseEmailToken(token)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckTokenTimeout
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userID = claims.UserID
	email = claims.Email
	password = claims.Password
	operationType = claims.OperationType

	//获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userID)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	switch operationType {
	case 1:
		//1:绑定邮箱
		user.Email = email
	case 2:
		//2：解绑邮箱
		user.Email = ""
	case 3:
		//3：修改密码
		err = user.SetPassword(password)
		if err != nil {
			utils.LogrusObj.Info(err)
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userID, user)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 成功则返回用户的信息
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// 显示用户金额
func (service *ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    e.GetMsg(code),
	}
}
