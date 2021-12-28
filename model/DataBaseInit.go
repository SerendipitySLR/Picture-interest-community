package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Create()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/fileTest?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	db.AutoMigrate(&User_detail{})
	db.AutoMigrate(&User_register{})
}

func GetDataBase() *gorm.DB{
	dsn := "root:123456@tcp(127.0.0.1:3306)/fileTest?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	return db
}