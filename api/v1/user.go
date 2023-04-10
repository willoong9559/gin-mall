package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/service"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, err)
	}
	res := userRegisterService.Register(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, err)
	}
	res := userLogin.Login(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

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

func ValidEmail(c *gin.Context) {
	var vaildEmailService service.ValidEmailService
	if err := c.ShouldBind(&vaildEmailService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
	res := vaildEmailService.Valid(c.Request.Context(), c.GetHeader("Authorization"))
	c.JSON(200, res)
}

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
