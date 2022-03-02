package mainpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
)

//打印评论的评论
func ShowMoreComment(c *gin.Context) {
	parentId:= c.Query("CommentId")
	//publisherId := c.Query("publisherId")
	db := respository.GetDB()
	var comments []model.Comment
	//db.AutoMigrate(&comments)

	//select语句
	db.Where("parent_id = ?",parentId).Find(&comments)

	//db.Where("post_id = ? AND post_type = ?",postId,postType).Find(&comments)
	//c.JSON(http.StatusOK, gin.H{
	//	"CommentList": comments,
	//})
	//var temUserDetails model.UserDetails
	var sendComment []model.SendComment
	for _,comment := range comments{
		//通过userId查找UserDetails的表项，因为要使用到其中的用户名，头像信息
		var temUserDetails model.UserDetails
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
