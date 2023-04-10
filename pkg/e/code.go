package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	// User Error
	ErrorExistUser = 30001 + iota
	ErrorFailEncryption
	ErrorExistUserNotFound
	ErrorNotCompare
	ErrorAuthToken
	ErrorAuthCheckTokenTimeout
	ErrorUploadFail
	ErrorSendEmail

	// Product Error
	ErrorProductUpload = 40001 + iota
)
