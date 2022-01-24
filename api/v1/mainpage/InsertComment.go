package mainpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strconv"
)
//插入评论
func InsertComment(c *gin.Context) {
	postId,_ := strconv.Atoi(c.PostForm("PostId"))
	postType,_ := strconv.Atoi(c.PostForm("PostType"))
	userId,_ := strconv.Atoi(c.PostForm("UserId"))
	content := c.PostForm("Content")
	parentId,_ := strconv.Atoi(c.PostForm("ParentId"))
	db := respository.GetDB()
	//publisherId := c.Query("publisherId")
	var comment model.Comment
	//db.AutoMigrate(&comment)
	comment.PostId = postId
	comment.PostType = postType
	comment.UserId = userId
	comment.Content = content
	comment.ParentId = parentId
	//设置直接评论帖子的parentId为-1，否则为对应的commentId
	if parentId != -1 {
		var parentComment model.Comment
		parentComment.CommentId = parentId
		//查找对应commentId的parentComment
		db.First(&parentComment)
		parentComment.ChildNumber++
		//childNumber递增后存入comment表中
		db.Save(&parentComment)
	}
	db.Save(&comment)
	response.Response(c,200)
	//c.JSON(http.StatusOK, gin.H{
	//	"Comment": comment,
	//})
}
