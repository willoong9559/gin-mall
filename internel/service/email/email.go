package email

import (
	"context"
	"strings"
	"time"

	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/internel/dao"
	"github.com/willoong9559/gin-mall/internel/model"
	e "github.com/willoong9559/gin-mall/pkg/errcode"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/serializer"
	"gopkg.in/mail.v2"
)

type SendEmailService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ValidEmailService struct {
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
