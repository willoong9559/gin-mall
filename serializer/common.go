package serializer

import e "github.com/willoong9559/gin-mall/pkg/errcode"

// Response 基础序列化器
type Response struct {
	Status e.CustomError `json:"status"`
	Msg    string        `json:"msg"`
	Data   interface{}   `json:"data"`
}

func GetResponse(status e.CustomError, data interface{}) *Response {
	return &Response{
		Status: status,
		Msg:    e.GetMsg(status),
		Data:   data,
	}
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

func GetListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}
