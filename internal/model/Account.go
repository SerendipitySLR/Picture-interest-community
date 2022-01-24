package model

import (
	"gorm.io/gorm"
	"time"
)

//用户相关信息
type UserRegister struct {
	UID       int `gorm:"primaryKey;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserName  string
	PassWord  string
	Telephone string
	Email     string
}

type UserDetails struct {
	UserId     int `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	NickName   string
	Sex        string
	Birthday   string
	Location   string
	Signature  string
	ProfileUrl string
}

type UserRelatedData struct {
	UserId         int `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	FollowerNumber int
	FansNumber     int
	PostNumber     int
	CollectNumber  int
	ForwardNumber  int
	ProfileUrl     string
}

func (UserDetails) TableName() string {
	return "user_details"
}
func (UserRelatedData) TableName() string {
	return "user_related_data"
}
