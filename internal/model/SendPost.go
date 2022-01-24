package model

// 每个发送的帖子都会包括post和forward，若是帖子，则forward为空，若是转发，则post为空，通过postType辨别
type SendPost struct {
	Post                Post
	Forward             Forward
	PostType            int
	PublisherName       string
	PublisherProfileUrl string
	SenderName          string
	SenderProfileUrl    string
	PhotoUrl            []string
}

//构造函数，构造一个post
func NewSendPost(post Post, user UserDetails, photoUrl []string) *SendPost {
	return &SendPost{
		Post:                post,
		PostType:            0,
		PublisherName:       user.NickName,
		PublisherProfileUrl: user.ProfileUrl,
		PhotoUrl:            photoUrl,
	}
}

//构造函数，构造一个forward
func NewSendForward(forward Forward, publisher UserDetails, sender UserDetails, photoUrl []string) *SendPost {
	return &SendPost{
		Forward:             forward,
		PostType:            1,
		PublisherName:       publisher.NickName,
		PublisherProfileUrl: publisher.ProfileUrl,
		SenderName:          sender.NickName,
		SenderProfileUrl:    sender.ProfileUrl,
		PhotoUrl:            photoUrl,
	}
}
