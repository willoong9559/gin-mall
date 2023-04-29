package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/willoong9559/gin-mall/internel/localization"
	e "github.com/willoong9559/gin-mall/pkg/errcode"
	"github.com/willoong9559/gin-mall/serializer"
)

func handleBindingErr(c *gin.Context, err error) {
	// validator binding error
	if errs, ok := localization.GetValidationErrors(err); ok {
		c.JSON(http.StatusBadRequest, serializer.GetResponse(http.StatusBadRequest, errs.Translate()))
		return
	}
	// UnmarshalTypeError（unusual）
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		c.JSON(http.StatusBadRequest, serializer.GetResponse(e.ErrorUnmarshalType, ""))
	}
	// others errors
	c.JSON(http.StatusBadRequest, serializer.GetResponse(http.StatusBadRequest, ""))
	return
}
