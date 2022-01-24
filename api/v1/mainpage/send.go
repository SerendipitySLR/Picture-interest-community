package mainpage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ptc/internal/application/command"
	errmsg "ptc/internal/ermsg"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strings"
)

func Send(c *gin.Context) {
	telephone := c.GetString("telephone")
	uid := respository.GetUID(telephone)

	var creatPost command.CreatPost

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&creatPost); err != nil {
		response.Response(c, errmsg.POST_FORMAT_ERROR)
		return
	}

	// 新建 post 行准备写入数据库
	post := model.Post{
		PublisherId:      creatPost.UID,
		PhotoNumber:      0,
		Content:          creatPost.Contents,
		CommentNumber:    0,
		ForwardNumber:    0,
		LikeNumber:       0,
		CollectionNumber: 0,
		PhotoPathUrl:     creatPost.PictureUrl,
		Location:         creatPost.Location,
	}

	// 写入数据库
	err := respository.GetDB().Create(&post).Error
	if err != nil {
		response.Response(c, errmsg.POST_WRITE_ERROR)
		return
	}

	// 写入 PostPhoto 表
	photos := strings.Split(creatPost.PictureUrl, ";")
	for _, photo := range photos {
		respository.GetDB().Create(&model.PostPhoto{
			PostId:   post.PostId,
			PhotoUrl: photo,
		})
	}

	// 查询所有的粉丝 uid
	followers := make([]model.Follow, 0)
	err = respository.GetDB().Where("user_id = ?", uid).Find(&followers).Error
	if err != nil {
		response.Response(c, errmsg.UID_IS_NOT_EXIST)
		return
	}

	// 写入 feeds 表
	for _, follower := range followers {
		followerID := follower.FollowID
		feed := model.Feeds{
			UserId:   followerID,
			PostId:   post.PostId,
			SendId:   uid,
			PostType: model.POST_TYPE,
		}
		err = respository.GetDB().Create(&feed).Error
		if err != nil {
			response.Response(c, errmsg.FEED_WRITE_ERROR)
			return
		}
	}

	// 修改客户的 relateData 信息
	err = respository.GetDB().Model(&model.UserRelatedData{}).Where("user_id = ?", uid).Update("post_number", gorm.Expr("post_number + ?", 1)).Error
	if err != nil {
		fmt.Println(err.Error())
		response.Response(c, errmsg.RELATEDATA_WRITE_ERROR)
		return
	}

	data := make(map[string]interface{})
	data["postId"] = post.PostId
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}
