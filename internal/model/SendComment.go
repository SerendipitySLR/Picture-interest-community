package model

type SendComment struct {
	Comment    Comment
	UserName   string
	ProfileUrl string
}

func NewSendComent(comment Comment, userName string, profileUrl string) *SendComment {
	return &SendComment{
		Comment:    comment,
		UserName:   userName,
		ProfileUrl: profileUrl,
	}
}