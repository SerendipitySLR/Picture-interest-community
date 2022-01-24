package model

import "time"

type Like struct {
	PostId    int `gorm:"primary_key"`
	UserId    int `gorm:"primary_key"`
	State     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Like) TableName() string {
	return "like"
}
