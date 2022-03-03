package mainpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"sort"
	"strconv"
)

//显示帖子页面，这块的逻辑可以参照接口文档，基本保持一致
func ShowPage(c *gin.Context) {
	userId := c.Query("UserId")
	userId_int,_ := strconv.Atoi(userId)
	db := respository.GetDB()
	//db.AutoMigrate(&comments)
	//var followIds []int
	//// select followId from follow where user_id = userId;
	//db.Where("user_id = ?", userId).Model(&model.Follow{}).Pluck("follow_id", &followIds)
	//var feeds []model.Feeds
	////遍历followIds，从feeds表中取出相应的，因为要使用对应的userId
	//for _, followId := range followIds {
	//	var feedsByFollowId []model.Feeds
	//	db.Where("user_id = ?", followId).Find(&feedsByFollowId)
	//	for _, feedByFollowId := range feedsByFollowId {
	//		feeds = append(feeds, feedByFollowId)
	//	}
	//}

	var feeds []model.Feeds
	db.Where("user_id = ?", userId).Find(&feeds)

	var likes []model.Like
	db.Where("user_id = ?", userId).Find(&likes)
	likeMap := make(map[model.Like]bool)
	for _, like := range likes {
		temlike := model.NewLike(like.UserId,like.PostId,like.PostType)
		likeMap[*temlike] = true
	}
	//var posts []model.Post
	//var forwards []model.Forward

	//这里定义了新的数据类型，为了保证发送的次序按时间排序，前端不需要再做类似逻辑，详见model/SendPost
	var sendPage []model.SendPost
	for _, feed := range feeds {
		//var temUserDetails model.UserDetails
		////通过userId查找UserDetails的表项，因为要使用到其中的用户名，头像信息
		//db.Where("user_id = ?", feed.UserId).Find(&temUserDetails)
		////生成头像的url
		//temUserDetails.ProfileUrl = "http://" + c.Request.Host + temUserDetails.ProfileUrl
		//postType为0是post，为1是forward，从相应的表中取出sendPost所需要的信息，装填进去
		if feed.PostType == 0 {
			var temPost model.Post
			var photoUrls []string
			db.Where("post_id = ?", feed.PostId).Find(&temPost)

			var temUserDetails model.UserDetails
			//通过userId查找UserDetails的表项，因为要使用到其中的用户名，头像信息
			db.Where("user_id = ?", temPost.PublisherId).Find(&temUserDetails)
			//生成头像的url
			temUserDetails.ProfileUrl = "http://" + c.Request.Host + temUserDetails.ProfileUrl


			db.Where("post_id = ?", feed.PostId).Model(&model.PostPhoto{}).Pluck("photo_url", &photoUrls)

			for i, _ := range photoUrls {
				photoUrls[i] = "http://" + c.Request.Host + photoUrls[i]
			}
			temLike := model.NewLike(userId_int,temPost.PostId,0)
			isliked := false
			if likeMap[*temLike] == true{
				isliked = true
			}

			temSendPage := model.NewSendPost(temPost, temUserDetails, photoUrls,isliked)
			sendPage = append(sendPage, *temSendPage)
		} else {
			var temForwards model.Forward
			var sender model.UserDetails
			var photoUrls []string
			db.Where("forward_id = ?", feed.PostId).Find(&temForwards)

			var temUserDetails model.UserDetails
			//通过userId查找UserDetails的表项，因为要使用到其中的用户名，头像信息
			db.Where("user_id = ?", temForwards.UserId).Find(&temUserDetails)
			//生成头像的url
			temUserDetails.ProfileUrl = "http://" + c.Request.Host + temUserDetails.ProfileUrl


			db.Where("post_id = ?", temForwards.PostId).Model(&model.PostPhoto{}).Pluck("photo_url", &photoUrls)
			for i, _ := range photoUrls {
				photoUrls[i] = "http://" + c.Request.Host + photoUrls[i]
			}
			db.Where("user_id = ?", feed.SendId).Find(&sender)
			sender.ProfileUrl = "http://" + c.Request.Host + sender.ProfileUrl

			temLike := model.NewLike(userId_int,temForwards.ForwardId,1)
			isliked := false
			if likeMap[*temLike] == true{
				isliked = true
			}
			temSendPage := model.NewSendForward(temForwards, temUserDetails, sender, photoUrls,isliked)
			sendPage = append(sendPage, *temSendPage)
		}
	}
	//根据更新时间倒序排序
	sort.SliceStable(sendPage, func(i, j int) bool {
		if sendPage[i].PostType == 0 && sendPage[j].PostType == 0 {
			return sendPage[i].Post.CreatedAt.After(sendPage[j].Post.CreatedAt)
		} else if sendPage[i].PostType == 0 && sendPage[j].PostType == 1 {
			return sendPage[i].Post.CreatedAt.After(sendPage[j].Forward.CreatedAt)
		} else if sendPage[i].PostType == 1 && sendPage[j].PostType == 0 {
			return sendPage[i].Forward.CreatedAt.After(sendPage[j].Post.CreatedAt)
		}
		return sendPage[i].Forward.CreatedAt.After(sendPage[j].Forward.CreatedAt)
	})
	data := make(map[string]interface{})
	data["PostList"] = sendPage
	response.ResponseWithData(c, 200, data)
	//filePath := c.Query("file1")
	//c.JSON(http.StatusOK, gin.H{"url": "http://" + c.Request.Host + filePath})
}
