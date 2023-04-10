package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/pkg/utils"
	"github.com/willoong9559/gin-mall/service"
)

func ListCarousels(c *gin.Context) {
	var listCarouselsService service.ListCarouselsService
	if err := c.ShouldBind(&listCarouselsService); err != nil {
		utils.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
	res := listCarouselsService.List(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
