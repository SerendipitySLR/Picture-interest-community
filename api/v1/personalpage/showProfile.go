package personalpage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
)

func ShowProfile(c *gin.Context) {
	userId := c.Query("UserId")
	//fmt.Println(userId)
	db := respository.GetDB()
	var account0 model.UserDetails
	var account1 model.UserRegister

	//select语句
	db.Where("user_id = ?", userId).Find(&account0)
	db.Where("uid = ?", userId).Find(&account1)

	data := make(map[string]interface{})
	data["accountInfo"] = account0
	data["UserName"] = account1.UserName
	data["Telephone"] = account1.Telephone
	data["Email"] = account1.Email
	fmt.Println(data)
	response.ResponseWithData(c, 200, data)
}
