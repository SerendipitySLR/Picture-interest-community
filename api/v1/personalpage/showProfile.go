package personalpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
)

func ShowProfile(c *gin.Context) {
	userId := c.Query("UserId")
	//fmt.Println(userId)
	db := respository.GetDB()
	//publisherId := c.Query("publisherId")
	var account model.UserDetails
	//db.AutoMigrate(&comments)

	//select语句
	db.Where("user_id = ?", userId).Find(&account)

	//db.Where("post_id = ? AND post_type = ?",postId,postType).Find(&comments)
	//c.JSON(http.StatusOK, gin.H{
	//	"CommentList": comments,
	//})

	data := make(map[string]interface{})
	data["accountInfo"] = account
	//fmt.Println(data)
	response.ResponseWithData(c, 200, data)
}
