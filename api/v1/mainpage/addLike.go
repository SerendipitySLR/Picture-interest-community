package mainpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strconv"
)

//点赞
func AddLike(c *gin.Context) {
	//参数
	postId := c.PostForm("PostId")
	userId := c.PostForm("UserId")
	//publisherId := c.Query("publisherId")
	db := respository.GetDB()
	var post model.Post
	var like model.Like
	//转换成字符串
	post.PostId,_ = strconv.Atoi(postId)

	//根据postId查找post表，存入post变量中  假设为变量与数据库绑定
	db.Find(&post)
	post.LikeNumber++
	//likeNumber增加1后，保存入post表项中
	db.Save(&post)

	like.PostId,_ = strconv.Atoi(postId)
	like.UserId,_ = strconv.Atoi(userId)
	//新建like表项
	db.Create(&like)
	response.Response(c,200)
	//c.JSON(http.StatusOK, gin.H{
	//	"post": post,
	//	"like": like,
	//})
}
