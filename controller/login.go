package controller

import (
	"Picture-interest-community/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

func check(passwd string, after_md5 string) bool{
	h := md5.New()
	h.Write([]byte(passwd))
	m := hex.EncodeToString(h.Sum(nil))
	return m == after_md5
}

func Login(c *gin.Context){
	userName := c.PostForm("username")
	password := c.PostForm("password")

	db := model.GetDataBase()
	var result []model.User_register
	r := db.Where("user_name = ?",userName).Find(&model.User_register{}).Scan(&result)
	if (r.RowsAffected == 0){
		c.JSON(http.StatusOK,gin.H{
			"status":0,
			"errmsg":"账号不存在",
		})
	}else {
		flag := false
		for i:= 0;i< len(result);i++{
			if(check(password,result[i].PassWord)){
				flag = true
				break
			}
		}
		if flag{
			c.JSON(http.StatusOK,gin.H{
				"status":1,
				"errmsg":nil,
			})
		}else {
			c.JSON(http.StatusOK,gin.H{
				"staus":0,
				"errmsg":"密码错误",
			})
		}

	}
}