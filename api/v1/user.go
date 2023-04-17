package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/service"
)

// @Summary 用户注册
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, err)
	}
	res := userRegisterService.Register(c)
	c.JSON(http.StatusOK, res)
}

// @Summary 用户登录
// @Produce  json
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, err)
	}
	res := userLogin.Login(c)
	c.JSON(http.StatusOK, res)
}

// @Summary 更新用户昵称
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
	var userUpdate service.UserService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
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
	uploadAvatarService := service.UserService{}
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatarService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
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
	var sendEmail service.SendEmailService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmail); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
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
	var vaildEmailService service.ValidEmailService
	if err := c.ShouldBind(&vaildEmailService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
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
	var showMoneyService service.ShowMoneyService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoneyService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, err)
	}
	res := showMoneyService.Show(c.Request.Context(), claims.ID)
	c.JSON(200, res)
}
