package controller

import (
	"Picture-interest-community/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func build_Random_Account()(account string){
	//生成字符串集,A-Z,a-z,0-9,_
	var random_char [26 + 26 + 10 + 1]byte
	index := 0
	for i := 0;i<26;i++{
		random_char[index] = 'A' + byte(i)
		index++
	}
	for i := 0;i<26;i++{
		random_char[index] = 'a' + byte(i)
		index++
	}
	for i := 0;i<10;i++{
		random_char[index] = '0' + byte(i)
		index++
	}
	random_char[index]='_'

	//随机选取10个字符
	count := 26 + 26 + 10 + 1
	rand.Seed(time.Now().Unix())

	var chs [10]byte
	for i := 0;i<10;i++{
		chs[i]=random_char[rand.Intn(count)]
	}
	account = string(chs[:])
	return
}


func parse_RegisterPost(c *gin.Context)(user_detail model.User_detail, user_register model.User_register){

	//生成datail表需要的数据结构
	nickname := c.PostForm("nickname")
	var sex bool
	if c.PostForm("sex") == "male"{
		sex = true
	}else {
		sex = false
	}
	fmt.Println("-------------------------------------------------")
	birthday,_ := time.Parse("20060102",c.PostForm("birthday"))
	print("11111111111111111111111111111111111111111111111111111111111")
	location := c.PostForm("location")
	signature := c.PostForm("signature")
	profile_url := c.PostForm("profile_url")

	user_detail = model.User_detail{
		Nickname:nickname,
		Sex:sex,
		Register_data:time.Now(),
		Update_date:time.Now(),
		Birthday:birthday,
		Location:location,
		Signature: signature,
		Profile_url: profile_url,
	}


	//生成User_register数据结构
	//生成随机的userName也就是账户名
	userName := build_Random_Account()


	//用户密码加密后存入数据库，加密算法---MD5
	passWord := c.PostForm("password")
	h := md5.New()
	h.Write([]byte(passWord))
	passWord = hex.EncodeToString(h.Sum(nil))

	telephone := c.PostForm("telephone")
	email := c.PostForm("email")

	user_register = model.User_register{
		UserName: userName,
		PassWord: passWord,
		Telephone: telephone,
		Email: email,
	}
	return
}

func Register(c *gin.Context)  {
	db := model.GetDataBase()
	user_detail,user_register:= parse_RegisterPost(c)
	r1 := db.Create(&user_detail)
	r2 := db.Create(&user_register)
	if( r1.Error == nil && r2.Error == nil ){
		c.JSON(http.StatusOK, gin.H{
			"status":1,
			"account":user_register.UserName,
			"msg":"success",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"status":0,
			"account":nil,
			"msg":"fail",
		})
	}
}
