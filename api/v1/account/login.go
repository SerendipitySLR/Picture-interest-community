package account

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/application/command"
	errmsg "ptc/internal/ermsg"
	"ptc/internal/response"
	"ptc/internal/respository"
	"ptc/pkg/jwt"
)

//登录接口
func Login(c *gin.Context) {
	var account command.Account
	_ = c.ShouldBindJSON(&account)
	isExist := respository.CheckAccount(account.Telephone)
	if !isExist {
		//账号不存在
		response.Response(c, errmsg.ERROR_ACCOUNT_NOT_EXIST)
		return
	}

	password, uid, err := respository.QueryAccountByTelephone(account.Telephone)

	if err != nil || password != account.Password {
		response.Response(c, errmsg.ERROR_PASSWORD_ERROR)
	} else {
		//成功登录颁发token
		data := make(map[string]interface{})
		token, err := jwt.GetToken(account.Telephone)
		if err == nil {
			data["token"] = token
			data["uid"] = uid
			response.ResponseWithData(c, errmsg.SUCCESS, data)
			return
		}
	}
}
