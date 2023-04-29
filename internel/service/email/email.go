package email

import (
	"context"
	"strings"
	"time"

	"gopkg.in/mail.v2"

	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/global"
	"github.com/willoong9559/gin-mall/internel/dao"
	"github.com/willoong9559/gin-mall/internel/model"
	e "github.com/willoong9559/gin-mall/pkg/errcode"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/serializer"
)

type SendEmailService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	//OperationType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ValidEmailService struct {
}

// Send 发送邮件
func (service *SendEmailService) Send(ctx context.Context, id uint) serializer.Response {
	address := "" // 验证地址
	var notice *model.Notice
	token, err := utils.GenerateEmailToken(id, service.OperationType, service.Email, service.Password)
	if err != nil {
		global.Logger.Error(err)
		return *serializer.GetResponse(e.ErrorAuthToken, "")
	}

	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		global.Logger.Error(err)
		return *serializer.GetResponse(e.ERROR, "")
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
		global.Logger.Warn("返送邮件失败:%s", err)
		return *serializer.GetResponse(e.ErrorSendEmail, "")
	}
	return *serializer.GetResponse(e.SUCCESS, "")
}

// Valid 验证邮箱
func (service ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint

	//验证token
	if token == "" {
		global.Logger.Warn("客户端未携带token")
	}
	claims, err := utils.ParseEmailToken(token)
	if err != nil {
		global.Logger.Error(err)
		return *serializer.GetResponse(e.ErrorAuthToken, "")
	}
	if time.Now().Unix() > claims.ExpiresAt {
		global.Logger.Error(err)
		return *serializer.GetResponse(e.ErrorAuthCheckTokenTimeout, "")
	}
	userID = claims.UserID
	email = claims.Email
	password = claims.Password
	operationType = claims.OperationType

	//获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userID)
	if err != nil {
		global.Logger.Error(err)
		return *serializer.GetResponse(e.ERROR, "")
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
			global.Logger.Error(err)
			return *serializer.GetResponse(e.ERROR, "")
		}
	}
	err = userDao.UpdateUserById(userID, user)
	if err != nil {
		global.Logger.Error(err)
		return *serializer.GetResponse(e.ERROR, "")
	}
	// 成功则返回用户的信息
	return *serializer.GetResponse(e.SUCCESS, serializer.BuildUser(user))
}
