package service

import (
	"machine-geek.cn/recruit-server/database"
	"machine-geek.cn/recruit-server/model"
)

func GetAllAdmin(page, size int, keyWord string) []model.Admin {
	var admins []model.Admin
	database.DB.Offset((page - 1) * size).Limit(size).Find(&admins)
	return admins
}

func GetAdminByName(name string) *model.Admin {
	var admin model.Admin
	database.DB.Where("name = ?", name).First(&admin)
	return &admin
}

func Count(keyWord string) int {
	keyWord = "%" + keyWord + "%"
	var count int
	database.DB.Model(&model.Admin{}).Where("id like ?", keyWord).Where("name like ?", keyWord).Where("introduce like ?", keyWord).Count(&count)
	return count
}

func AddAdmin(admin *model.Admin) {
	database.DB.Create(admin)
}

func DeleteAdmin(id uint) {
	database.DB.Delete(&model.Admin{Id: id})
}

func ModifyAdmin(admin *model.Admin) {
	database.DB.Update(admin)
}
