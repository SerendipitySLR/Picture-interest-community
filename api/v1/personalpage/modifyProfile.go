package personalpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strconv"
)

func ModifyProfile(c *gin.Context) {
	//参数
	userId := c.Query("UserId")
	nickname := c.Query("NickName")
	sex := c.Query("Sex")
	birthday := c.Query("Birthday")
	location := c.Query("Location")
	signature := c.Query("Signature")
	profileUrl := c.Query("ProfileUrl")
	db := respository.GetDB()

	var uDetails model.UserDetails

	//对于userdetails进行修改
	uDetails.UserId, _ = strconv.Atoi(userId)
	db.Find(&uDetails)
	uDetails.ProfileUrl = profileUrl
	uDetails.Birthday = birthday
	uDetails.Location = location
	uDetails.NickName = nickname
	uDetails.Sex = sex
	uDetails.Signature = signature
	db.Save(&uDetails)

	//select语句
	db.Where("user_id = ?", userId).Find(&uDetails)
	data := make(map[string]interface{})
	data["accountInfo"] = uDetails
	//fmt.Println(data)
	response.ResponseWithData(c, 200, data)
	//response.Response(c, 200)
	//c.JSON(http.StatusOK, gin.H{
	//	"post": post,
	//	"like": like,
	//})
}
