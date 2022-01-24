package personalpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strconv"
)

func ShowPersonalInfo(c *gin.Context) {

	//先通过Uid找到用户的nickname，
	userId := c.Query("UserId")
	db := respository.GetDB()
	var userDetails model.UserDetails
	//var nickName string		//用于接收用户的nickname
	userDetails.UserId, _ = strconv.Atoi(userId)
	db.Find(&userDetails)
	//nickName = userDetails.NickName

	//再利用Uid找到用户的相关信息比如profile、postsNumber······
	var userRelatedData model.UserRelatedData
	userRelatedData.UserId, _ = strconv.Atoi(userId)
	db.Find(&userRelatedData)

	//拿到基本信息
	data := make(map[string]interface{})
	data["Nickname"] = userDetails.NickName
	data["Profile"] = userRelatedData.ProfileUrl
	data["PostsNumber"] = userRelatedData.PostNumber
	data["FansNumber"] = userRelatedData.FansNumber
	data["FollowsNumber"] = userRelatedData.FollowerNumber

	response.ResponseWithData(c, 200, data)
}
