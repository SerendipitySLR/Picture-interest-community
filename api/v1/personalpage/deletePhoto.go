package personalpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strconv"
)

func DeletePhoto(c *gin.Context) {
	//参数
	userId := c.Query("UserId")
	db := respository.GetDB()

	var uDetails model.UserDetails
	//对于userdetails进行修改
	uDetails.UserId, _ = strconv.Atoi(userId)
	db.Find(&uDetails)
	uDetails.ProfileUrl = "nil"
	db.Save(&uDetails)
	//返回修改成功信息
	response.Response(c, 200)

}
