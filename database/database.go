package database

import (
	"github.com/jinzhu/gorm"
	"machine-geek.cn/recruit-server/model"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:xy942698.@tcp(localhost:3306)/geek_recruit?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&model.Admin{})
	DB.SingularTable(true)
}
