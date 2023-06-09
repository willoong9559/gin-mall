package e

type CustomError int

const (
	SUCCESS       CustomError = 200
	ERROR         CustomError = 500
	InvalidParams CustomError = 400

	// User Error
	ErrorExistUser CustomError = 30001 + iota
	ErrorCaptcha
	ErrorDatabase
	ErrorFailEncryption
	ErrorExistUserNotFound
	ErrorNotCompare
	ErrorUnmarshalType

	ErrorAuthToken
	ErrorAuthCheckTokenTimeout
	ErrorUploadFail
	ErrorSendEmail

	// Product Error
	ErrorProductUpload CustomError = 40001 + iota
)
