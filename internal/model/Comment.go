package model

import "time"

type Comment struct {
	CommentId   int `gorm:"primary_key"`
	ParentId    int `gorm:"index:parentId_postId_postType"` //使用索引，注意顺序保持一致，否则无效
	PostId      int `gorm:"index:parentId_postId_postType"`
	PostType    int `gorm:"index:parentId_postId_postType"`
	UserId      int
	ChildNumber int
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (Comment) TableName() string {
	return "comment"
}
