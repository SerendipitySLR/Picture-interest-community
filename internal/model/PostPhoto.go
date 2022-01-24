package model

import "time"

type PostPhoto struct {
	PostId    int    `gorm:"primary_key"`
	PhotoUrl  string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (PostPhoto) TableName() string {
	return "post_photo"
}
