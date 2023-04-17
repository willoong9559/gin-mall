package v1

import (
	"encoding/json"
	"net/http"

	"github.com/willoong9559/gin-mall/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusBadRequest,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
