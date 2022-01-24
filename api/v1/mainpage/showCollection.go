package mainpage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
)

func ShowCollection(c *gin.Context) {
	userId := c.Query("UserId")
	fmt.Println(userId)
	db := respository.GetDB()
	//publisherId := c.Query("publisherId")
	var collections []model.Collection
	//db.AutoMigrate(&comments)

	//select语句
	db.Where("user_id = ?", userId).Find(&collections)

	//db.Where("post_id = ? AND post_type = ?",postId,postType).Find(&comments)
	//c.JSON(http.StatusOK, gin.H{
	//	"CommentList": comments,
	//})

	data := make(map[string]interface{})
	data["CollectionList"] = collections
	response.ResponseWithData(c, 200, data)
}
