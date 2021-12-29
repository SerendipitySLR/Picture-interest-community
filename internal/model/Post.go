package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	PublicherId      int
	PhotoNumber      int
	Content          string
	CommentNumber    int
	ForwardNumber    int
	LikeNumber       int
	CollectionNumber int
	PublishTime      time.Time
	PhotoPathhUrl    string
	Location         string
}
