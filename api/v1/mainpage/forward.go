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
)

func Forward(c *gin.Context) {
	telephone := c.GetString("telephone")
	uid := respository.GetUID(telephone)

	var creatForward command.CreatForward
	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&creatForward); err != nil {
		response.Response(c, errmsg.FORWARD_FORMAT_ERROR)
		return
	}

	// 新建 forward 行准备写入数据库
	forward := model.Forward{
		PostId:        creatForward.PostID,
		UserId:        creatForward.UID,
		State:         0, // 这是啥??
		CommentNumber: 0,
		LikeNumber:    0,
		Content:       creatForward.Contents,
	}

	// 写入数据库
	err := respository.GetDB().Create(&forward).Error
	if err != nil {
		response.Response(c, errmsg.FORWARD_WRITE_ERROR)
		return
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
			PostId:   forward.ForwardId,
			SendId:   uid,
			PostType: model.FOLLOW_TYPE,
		}
		err = respository.GetDB().Create(&feed).Error
		if err != nil {
			response.Response(c, errmsg.FEED_WRITE_ERROR)
			return
		}
	}

	// 修改用户的 relateData 信息
	err = respository.GetDB().Model(&model.UserRelatedData{}).Where("user_id = ?", uid).Update("forward_number", gorm.Expr("forward_number + ?", 1)).Error
	if err != nil {
		response.Response(c, errmsg.RELATEDATA_WRITE_ERROR)
		return
	}

	// 修该被转发的 post 的转发数
	err = respository.GetDB().Model(&model.Post{}).Where("post_id = ?", creatForward.PostID).Update("forward_number", gorm.Expr("forward_number + ?", 1)).Error
	if err != nil {
		fmt.Println(err.Error())
		response.Response(c, errmsg.POST_WRITE_ERROR)
		return
	}

	data := make(map[string]interface{})
	data["forwardId"] = forward.PostId
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}
