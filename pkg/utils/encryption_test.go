package utils_test

import (
	"testing"

	"github.com/willoong9559/gin-mall/pkg/utils"
)

func TestEncrypt(t *testing.T) {
	plaintext := "hello world"
	key := "abcde"
	utils.Encrypt.SetKey(key)
	ciphertext := utils.Encrypt.AesEncoding("hello world")
	if utils.Encrypt.AesDecoding(ciphertext) != plaintext {
		t.Error(`aes加密验证失败`)
	}
}
