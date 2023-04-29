package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/willoong9559/gin-mall/internel/service"
)

// @Summary 获取验证码
// @Produce  json
// @Success 200 imageResponse"成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/user/register [post]
func GetCaptcha(c *gin.Context) {
	var captchaService service.CaptchaService
	captchaService.GetCaptcha(c, 4)
}
