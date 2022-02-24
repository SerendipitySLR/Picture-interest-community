package model

import "time"

type Like struct {
	UserId    int `gorm:"primary_key"`
	PostId    int `gorm:"primary_key"`
	PostType  int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Like) TableName() string {
	return "like"
}

func NewLike(userId int, postId int, postType int) *Like {
	return &Like{
		UserId:   userId,
		PostId:   postId,
		PostType: postType,
	}
}
