package personalpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strconv"
)

func NewCollection(c *gin.Context) {
	//参数
	userId := c.PostForm("UserId")
	collectionName := c.PostForm("CollectionName")
	db := respository.GetDB()
	var collection model.Collection

	//对于collection进行修改
	collection.CollectionName = collectionName
	collection.UserId, _ = strconv.Atoi(userId)
	db.Create(&collection)

	response.Response(c, 200)
	//c.JSON(http.StatusOK, gin.H{
	//	"post": post,
	//	"like": like,
	//})
}
