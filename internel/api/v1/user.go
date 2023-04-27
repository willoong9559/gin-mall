package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/internel/service/email"
	"github.com/willoong9559/gin-mall/internel/service/user"
	service "github.com/willoong9559/gin-mall/internel/service/user"
	"github.com/willoong9559/gin-mall/pkg/utils"
)

// @Summary 用户注册
// @Accept  json
// @Produce  json
// @Param nick_name body string true "昵称" minlength(5) maxlength(8)
// @Param user_name body string true "用户名" minlength(5) maxlength(12)
// @Param password body string true "密码" minlength(5) maxlength(12)
// @Param re_password body string true "重复密码" minlength(5) maxlength(12)
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserRegisterService
	if err := c.ShouldBind(&userRegisterService); err != nil {
		handleBindingErr(c, err)
		return
	}
	res := userRegisterService.Register(c)
	c.JSON(http.StatusOK, res)
}

// @Summary 用户登录
// @Accept  json
// @Produce  json
// @Param user_name body string true "用户名" minlength(5) maxlength(12)
// @Param password body string true "密码" minlength(5) maxlength(12)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var userLogin service.UserLoginService
	if err := c.ShouldBind(&userLogin); err != nil {
		handleBindingErr(c, err)
		return
	}
	res := userLogin.Login(c)
	c.JSON(http.StatusOK, res)
}

// @Summary 更新用户昵称
// @Accept  json
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/user [put]
func UserUpdate(c *gin.Context) {
	var userUpdate service.UserUpdateService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err != nil {
		handleBindingErr(c, err)
		return
	}
	res := userUpdate.Update(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

// @Summary 更改用户头像
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/avatar [post]
func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	uploadAvatarService := user.UserService{}
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatarService); err != nil {
		handleBindingErr(c, err)
		return
	}
	res := uploadAvatarService.Post(c.Request.Context(), claims.ID, file, fileSize)
	c.JSON(http.StatusOK, res)
}

// @Summary 绑定邮箱
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/user/sending-email [post]
func SendEmail(c *gin.Context) {
	var sendEmail email.SendEmailService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmail); err != nil {
		handleBindingErr(c, err)
		return
	}
	res := sendEmail.Send(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

// @Summary 验证邮箱
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/user/valid-email [post]
func ValidEmail(c *gin.Context) {
	var vaildEmailService email.ValidEmailService
	if err := c.ShouldBind(&vaildEmailService); err != nil {
		handleBindingErr(c, err)
		return
	}
	res := vaildEmailService.Valid(c.Request.Context(), c.GetHeader("Authorization"))
	c.JSON(200, res)
}

// @Summary 获取用户金额
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/money [post]
func ShowMoney(c *gin.Context) {
	var showMoneyService user.ShowMoneyService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoneyService); err != nil {
		handleBindingErr(c, err)
		return
	}
	res := showMoneyService.Show(c.Request.Context(), claims.ID)
	c.JSON(200, res)
}
