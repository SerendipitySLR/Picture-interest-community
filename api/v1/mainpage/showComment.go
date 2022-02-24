package mainpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
)

//打印评论
func ShowComment(c *gin.Context) {
	postId:= c.Query("PostId")
	postType:= c.Query("PostType")
	db := respository.GetDB()
	//publisherId := c.Query("publisherId")
	var comments []model.Comment
	//db.AutoMigrate(&comments)

	//select语句
	db.Where("parent_id = -1 AND post_id = ?" +
		" AND post_type = ?",postId,postType).Find(&comments)

	//db.Where("post_id = ? AND post_type = ?",postId,postType).Find(&comments)
	//c.JSON(http.StatusOK, gin.H{
	//	"CommentList": comments,
	//})
	var temUserDetails model.UserDetails
	var sendComment []model.SendComment
	for _,comment := range comments{
		//通过userId查找UserDetails的表项，因为要使用到其中的用户名，头像信息
		db.Where("user_id = ?", comment.UserId).Find(&temUserDetails)
		//生成头像的url
		temUserDetails.ProfileUrl = "http://" + c.Request.Host + temUserDetails.ProfileUrl
		temSendComment := model.NewSendComent(comment,temUserDetails.NickName,temUserDetails.ProfileUrl)
		sendComment = append(sendComment, *temSendComment)
	}


	data := make(map[string]interface{})
	data["CommentList"] = sendComment
	response.ResponseWithData(c,200,data)
}
