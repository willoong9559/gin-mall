package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willoong9559/gin-mall/internel/localization"
	"github.com/willoong9559/gin-mall/serializer"
)

func handleBindingErr(c *gin.Context, err error) {
	errs, ok := localization.GetValidationErrors(err)
	// validator binding error
	if ok {
		c.JSON(http.StatusBadRequest, serializer.GetResponse(http.StatusBadRequest, errs.Translate()))
		return
	}
	// others errors
	c.JSON(http.StatusBadRequest, serializer.GetResponse(http.StatusBadRequest, ""))
	return
}

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "JSON类型不匹配",
		}
	}
	return serializer.Response{
		Status: http.StatusBadRequest,
		Msg:    "参数错误",
	}
}
