package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/service"
)

func GetCaptcha(c *gin.Context) {
	var captchaService service.CaptchaService
	captchaService.GetCaptcha(c, 4)
}
