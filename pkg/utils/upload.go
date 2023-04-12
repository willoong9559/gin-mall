package utils

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/willoong9559/gin-mall/conf"
)

type UploadType int

const (
	Avatar UploadType = iota
	ProductImg
)

var TypeFileName = [...]string{Avatar: "user", ProductImg: "boss"}

func UploadToLocalStatic(file multipart.File, userId uint, Name string, utype UploadType) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	var confBasePath string
	switch utype {
	case Avatar:
		confBasePath = conf.AvatarPath
	case ProductImg:
		confBasePath = conf.ProductPath
	}
	basePath := "." + confBasePath + TypeFileName[utype] + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + Name + ".jpg" // todo: 把file的后缀提取出来
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return TypeFileName[utype] + bId + "/" + Name + ".jpg", nil
}
