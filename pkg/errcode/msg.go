package e

var MsgFlags = [...]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	ErrorCaptcha:           "验证码错误",
	ErrorDatabase:          "数据库错误",
	ErrorExistUser:         "用户名已存在",
	ErrorFailEncryption:    "密码加密失败",
	ErrorExistUserNotFound: "用户不存在",
	ErrorNotCompare:        "密码不匹配",
	ErrorUnmarshalType:     "JSON类型不匹配",

	ErrorAuthToken:             "token认证失败",
	ErrorAuthCheckTokenTimeout: "token过期",
	ErrorUploadFail:            "图片上传失败",
	ErrorSendEmail:             "邮件发送失败",
	ErrorProductUpload:         "产品图片上传错误",
}

// GetMsg 获取状态码对应信息
func GetMsg(code CustomError) string {
	msg := "自定义错误未定义"
	msg = MsgFlags[code]
	return msg
}
