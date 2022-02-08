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
	signature := c.Query("Signature")
	email := c.Query("Email")
	telephone := c.Query("Telephone")
	username := c.Query("UserName")
	profileUrl := c.Query("ProfileUrl")
	db := respository.GetDB()

	var uDetails model.UserDetails
	var uRegister model.UserRegister

	//对于userdetails进行修改
	uDetails.UserId, _ = strconv.Atoi(userId)
	db.Find(&uDetails)
	uDetails.ProfileUrl = profileUrl
	uDetails.NickName = nickname
	uDetails.Sex = sex
	uDetails.Signature = signature
	db.Save(&uDetails)
	//对于UserRegister进行修改
	uRegister.UID, _ = strconv.Atoi(userId)
	db.Find(&uRegister)
	uRegister.Telephone = telephone
	uRegister.Email = email
	uRegister.UserName = username
	db.Save(&uRegister)

	//select语句,返回修改后的内容
	db.Where("user_id = ?", userId).Find(&uDetails)
	db.Where("uid = ?", userId).Find(&uRegister)

	data := make(map[string]interface{})
	data["accountInfo"] = uDetails
	data["UserName"] = uRegister.UserName
	data["Telephone"] = uRegister.Telephone
	data["Email"] = uRegister.Email
	//fmt.Println(data)
	response.ResponseWithData(c, 200, data)
	//response.Response(c, 200)
	//c.JSON(http.StatusOK, gin.H{
	//	"post": post,
	//	"like": like,
	//})
}
