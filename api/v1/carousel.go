package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/service"
)

// @Summary 获取轮播图
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/carousels [post]
func ListCarousels(c *gin.Context) {
	var listCarouselsService service.ListCarouselsService
	if err := c.ShouldBind(&listCarouselsService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
	res := listCarouselsService.List(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
