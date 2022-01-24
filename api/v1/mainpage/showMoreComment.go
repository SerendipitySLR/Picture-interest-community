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

	data := make(map[string]interface{})
	data["CommentList"] = comments
	response.ResponseWithData(c,200,data)
}
