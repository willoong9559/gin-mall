package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/service"
)

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
