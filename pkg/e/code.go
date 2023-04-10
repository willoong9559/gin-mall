package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	// user错误
	ErrorExistUser = 30001 + iota
	ErrorFailEncryption
	ErrorExistUserNotFound
	ErrorNotCompare
	ErrorAuthToken
	ErrorAuthCheckTokenTimeout
	ErrorUploadFail
	ErrorSendEmail
)
