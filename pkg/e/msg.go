package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	ErrorExistUser:             "用户名已存在",
	ErrorFailEncryption:        "密码加密失败",
	ErrorExistUserNotFound:     "用户不存在",
	ErrorNotCompare:            "密码不匹配",
	ErrorAuthToken:             "token认证失败",
	ErrorAuthCheckTokenTimeout: "token过期",
	ErrorUploadFail:            "图片上传失败",
	ErrorSendEmail:             "邮件发送失败",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[ERROR]
	}
	return msg
}
