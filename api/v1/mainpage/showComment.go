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

	data := make(map[string]interface{})
	data["CommentList"] = comments
	response.ResponseWithData(c,200,data)
}
