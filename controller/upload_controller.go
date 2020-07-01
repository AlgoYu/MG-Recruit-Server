package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"machine-geek.cn/recruit-server/common"
	"math/rand"
	"os"
	"path"
	"time"
)

const (
	UPLOADPATH = "/static/upload"
)

// 允许图片类型
var allowPictureType = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

// 允许文件类型
var allowFileType = map[string]bool{
	".zip": true,
}

// 初始化目录
func init() {
	if _, err := os.Stat(UPLOADPATH); err != nil {
		os.MkdirAll(UPLOADPATH, os.ModePerm)
	}
}

// 图片上传
func UploadPicture(c *gin.Context) {
	file, _ := c.FormFile("picture")
	postfix := path.Ext(file.Filename)
	if allowPictureType[postfix] {
		c.JSON(200, common.Fail("文件格式错误！"))
	}
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2020_07_01_15_04_05_") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + postfix
	filePath := UPLOADPATH + fileName
	c.SaveUploadedFile(file, filePath)
	c.JSON(200,common.Success(filePath))
}

// 文件上传
func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	postfix := path.Ext(file.Filename)
	if allowFileType[postfix] {
		c.JSON(200, common.Fail("文件格式错误！"))
	}
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2020_07_01_15_04_05_") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + postfix
	filePath := UPLOADPATH + fileName
	c.SaveUploadedFile(file, filePath)
	c.JSON(200,common.Success(filePath))
}
