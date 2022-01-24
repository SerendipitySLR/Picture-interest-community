package model

import "time"

type Collection struct {
	CollectionId   int `gorm:"primary_key"`
	UserId         int
	CollectionName string
	PostsId        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func (Collection) TableName() string {
	return "collection"
}
