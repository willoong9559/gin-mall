package service

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/willoong9559/gin-mall/conf"
	"github.com/willoong9559/gin-mall/pkg/utils"
)

func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !utils.DirExistOrNot(basePath) {
		utils.CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg" // todo: 把file的后缀提取出来
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

func UploadProductToLocalStatic(file multipart.File, userId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "boss" + bId + "/"
	if !utils.DirExistOrNot(basePath) {
		utils.CreateDir(basePath)
	}
	avatarPath := basePath + productName + ".jpg" // todo: 把file的后缀提取出来
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "boss" + bId + "/" + productName + ".jpg", nil
}
