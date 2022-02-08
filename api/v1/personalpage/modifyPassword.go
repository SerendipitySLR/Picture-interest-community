package personalpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strconv"
)

func ModifyPassword(c *gin.Context) {
	//参数
	userId := c.Query("UserId")
	password := c.Query("Password")
	newPassword := c.Query("NewPassword")
	db := respository.GetDB()

	var uRegister model.UserRegister
	//对于userdetails进行修改
	uRegister.UID, _ = strconv.Atoi(userId)
	db.Find(&uRegister)
	if uRegister.PassWord != password {
		response.Response(c, 901)
	} else {
		uRegister.PassWord = newPassword
		db.Save(&uRegister)
		//返回修改成功信息
		response.Response(c, 200)
	}
}
