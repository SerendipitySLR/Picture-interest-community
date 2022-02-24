package mainpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strconv"
)

//点赞
func CancelLike(c *gin.Context) {
	//参数
	postId,_ := strconv.Atoi(c.PostForm("PostId"))
	userId,_ := strconv.Atoi(c.PostForm("UserId"))
	postType,_ := strconv.Atoi(c.PostForm("PostType"))
	//publisherId := c.Query("publisherId")
	db := respository.GetDB()
	var like model.Like
	//转换成字符串
	if postType == 0{
		var post model.Post
		post.PostId = postId
		//根据postId查找post表，存入post变量中  假设为变量与数据库绑定
		db.Find(&post)
		post.LikeNumber--
		//likeNumber增加1后，保存入post表项中
		db.Save(&post)
	} else{
		var forward model.Forward
		forward.ForwardId = postId
		//根据postId查找post表，存入post变量中  假设为变量与数据库绑定
		db.Find(&forward)
		forward.LikeNumber--
		//likeNumber增加1后，保存入post表项中
		db.Save(&forward)
	}

	like.PostId = postId
	like.UserId = userId
	like.PostType = postType
	//新建like表项
	db.Delete(&like)
	response.Response(c,200)
	//c.JSON(http.StatusOK, gin.H{
	//	"post": post,
	//	"like": like,
	//})
}