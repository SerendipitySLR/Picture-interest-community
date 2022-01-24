package model

import (
	"time"
)

type Forward struct {
	ForwardId     int `gorm:"primaryKey"`
	PostId        int
	UserId        int
	State         int
	CommentNumber int
	LikeNumber    int
	Content       string `gorm:"type:varchar(100)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

func (Forward) TableName() string {
	return "forward"
}
