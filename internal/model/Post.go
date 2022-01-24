package model

import (
	"time"
)

type Post struct {
	PostId           int `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
	PublisherId      int
	PhotoNumber      int
	Content          string `gorm:"type:varchar(500)"`
	CommentNumber    int
	ForwardNumber    int
	LikeNumber       int
	CollectionNumber int
	PhotoPathUrl     string
	Location         string
}
