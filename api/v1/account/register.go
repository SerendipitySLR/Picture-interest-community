package account

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/application/command"
	errmsg "ptc/internal/ermsg"
	"ptc/internal/response"
	"ptc/internal/respository"
	"ptc/pkg/jwt"
)

//注册接口
func Register(c *gin.Context) {
	var account command.Account
	_ = c.ShouldBindJSON(&account)

	// 检查用户是否存在
	isExist := respository.CheckAccount(account.Telephone)
	if isExist {
		response.Response(c, errmsg.ERROR_ACCOUNT_HAS_EXISTED)
		return
	}

	//添加新账号
	repoErr := respository.AddAccount(&account)
	if repoErr != nil {
		fmt.Println(repoErr.Error())
		response.Response(c, errmsg.ERROR)
		return
	}

	//成功注册颁发token
	data := make(map[string]interface{})
	token, err := jwt.GetToken(account.Telephone)
	if err != nil {
		fmt.Println(err.Error())
		response.Response(c, errmsg.ERROR)
		return
	}
	data["token"] = token
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}
