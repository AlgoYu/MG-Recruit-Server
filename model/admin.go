package model

import "machine-geek.cn/recruit-server/database"

type Admin struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	Name      string `gorm:"unique_index" json:"name"`
	Password  string `gorm:"not null;char(60)" json:"password"`
	Introduce string `gorm:"varchar(50)" json:"introduce"`
	Picture   string `gorm:"not null;varchar(255)" json:"picture"`
	CreatedAt string `gorm:"not null;datetime" json:"picture"`
}

func (admin *Admin) TableName() string {
	return "admin"
}

func GetAllAdmin(name string) *Admin {
	var admin Admin
	database.DB.Where("name = ?", name).First(&admin)
	return &admin
}

func GetAdminByName(name string) *Admin {
	var admin Admin
	database.DB.Where("name = ?", name).First(&admin)
	return &admin
}
