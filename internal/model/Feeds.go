package model

import "time"

const (
	FOLLOW_TYPE = 1
	POST_TYPE   = 0
)

type Feeds struct {
	UserId    int `gorm:"primaryKey"`
	PostId    int `gorm:"primaryKey"`
	PostType  int `gorm:"primary_key"`
	SendId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Feeds) TableName() string {
	return "feeds"
}
