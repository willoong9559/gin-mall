package serializer

import e "github.com/willoong9559/gin-mall/pkg/errcode"

// Response 基础序列化器
type Response struct {
	Status e.CustomError `json:"status"`
	Data   interface{}   `json:"data"`
	Msg    string        `json:"msg"`
	Error  string        `json:"error"`
}

//DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}
