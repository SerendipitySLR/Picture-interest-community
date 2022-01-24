package mainpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strconv"
)

func AddCollection(c *gin.Context) {
	//参数
	postId := c.Query("PostId")
	userId := c.Query("UserId")
	collectionId := c.Query("CollectionId")

	db := respository.GetDB()
	var post model.Post
	var userRelatedData model.UserRelatedData
	var collection model.Collection

	//对于collection进行修改
	collection.CollectionId, _ = strconv.Atoi(collectionId)
	db.Find(&collection)
	collection.PostsId += postId
	collection.PostsId += ","
	db.Model(&collection).Update("postsId", collection.PostsId)

	//根据UID查找userRelatedData表，然后对collectionNumber+1
	userRelatedData.UserId, _ = strconv.Atoi(userId)
	db.Find(&userRelatedData)
	userRelatedData.CollectNumber++
	db.Save(&userRelatedData)

	//根据postID查找post表，然后对CollectionNumber+1
	post.PostId, _ = strconv.Atoi(postId)
	db.Find(&post)
	post.CollectionNumber++
	db.Save(&post)

	response.Response(c, 200)
	//c.JSON(http.StatusOK, gin.H{
	//	"post": post,
	//	"like": like,
	//})
}
