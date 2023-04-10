package utils_test

import (
	"testing"

	"github.com/willoong9559/gin-mall/pkg/utils"
)

func TestTokenValid(t *testing.T) {
	token, err := utils.GenerateToken(1, "tom", 0)
	if err != nil {
		t.Error(err)
	}
	claim, err := utils.ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	if claim.ID != 1 || claim.Username != "tom" {
		t.Error("jwt验证失败")
	}
}
