package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/internel/service"
	"github.com/willoong9559/gin-mall/pkg/utils"
)

// @Summary 创建商品
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/product [post]
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := c.ShouldBind(&createProductService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
	res := createProductService.Create(c.Request.Context(), claim.ID, files)
	c.JSON(http.StatusOK, res)
}

// @Summary 展示用户商品
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/listProducts [post]
func ListProducts(c *gin.Context) {
	listProductService := service.ProductService{}
	if err := c.ShouldBind(&listProductService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
	res := listProductService.List(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

// @Summary 搜索商品
// @Produce  json
// @Param nick_name body string true "昵称" maxlength(100)
// @Param user_name body string true "用户名" maxlength(100)
// @Param password body int true "密码"
// @Param key body string true "密码加密key" minlength(16) maxlength(16)
// @Param captcha body string true "验证码" minlength(4) maxlength(4)
// @Success 200 {object} serializer.Response "成功"
// @Failure 400 {object} serializer.Response "请求错误"
// @Failure 500 {object} serializer.Response "内部错误"
// @Router /api/v1/searchProducts [post]
func SearchProducts(c *gin.Context) {
	searchProductsService := service.ProductService{}
	if err := c.ShouldBind(&searchProductsService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
	res := searchProductsService.Search(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
