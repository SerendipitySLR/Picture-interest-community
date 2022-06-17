package mainpage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ptc/internal/application/command"
	errmsg "ptc/internal/ermsg"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"strings"
	"path"
	"encoding/base64"
	"github.com/spf13/viper"
	"strconv"
	"os"
	"time"
)

func Send(c *gin.Context) {
	telephone := c.GetString("telephone")
	uid := respository.GetUID(telephone)

	var creatPost command.CreatPost

	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&creatPost); err != nil {
		response.Response(c, errmsg.POST_FORMAT_ERROR)
		return
	}
	
	
	// 保存图片
	var pictureUrl string
	if len(creatPost.ImgList) == 0 {
		response.Response(c, errmsg.HAVING_NO_PICTURE)
		return
	} else {
		var (
			enc      = base64.StdEncoding
			filepath string
		)
		for index, img := range creatPost.ImgList {
			fmt.Println("准备图片：。。。。。。。。")
			if img[11] == 'j' {
				img = img[23:]
				timeUnix := strconv.FormatInt(time.Now().Unix(), 10)
				filepath = path.Join(viper.GetString("server.PostImages"), timeUnix+strconv.Itoa(index)+".jpg")
			} else if img[11] == 'p' {
				img = img[22:]
				timeUnix := strconv.FormatInt(time.Now().Unix(), 10)
				filepath = path.Join(viper.GetString("server.PostImages"), timeUnix+strconv.Itoa(index)+".png")
			} else if img[11] == 'g' {
				img = img[22:]
				timeUnix := strconv.FormatInt(time.Now().Unix(), 10)
				filepath = path.Join(viper.GetString("server.PostImages"), timeUnix+strconv.Itoa(index)+".gif")
			} else {
				response.Response(c, errmsg.PICTURE_FORM_ERROR)
				return
			}
			//图片解码
			data, err := enc.DecodeString(img)
			if err != nil {
				response.Response(c, errmsg.PICTURE_DNCODE_ERROR)
				return
			}
			//图片写入文件
			f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				response.Response(c, errmsg.POST_PICTURE_WRITE_ERROR)
				return
			}
			defer f.Close()
			f.Write(data)
			filepath = "/" + filepath
			fmt.Println("当前图片地址为：" + filepath)
			//记录保存的地址
			if len(pictureUrl) == 0 {
				pictureUrl += filepath
			} else {
				pictureUrl += ";" + filepath
			}
		}
	}
	
	// 新建 post 行准备写入数据库
	post := model.Post{
		PublisherId:      uid,
		PhotoNumber:      0,
		Content:          creatPost.Contents,
		CommentNumber:    0,
		ForwardNumber:    0,
		LikeNumber:       0,
		CollectionNumber: 0,
		PhotoPathUrl:     pictureUrl,
		Location:         creatPost.Location,
	}

	// 写入数据库
	err := respository.GetDB().Create(&post).Error
	if err != nil {
		response.Response(c, errmsg.POST_WRITE_ERROR)
		return
	}

	// 写入 PostPhoto 表
	photos := strings.Split(pictureUrl, ";")
	for _, photo := range photos {
		respository.GetDB().Create(&model.PostPhoto{
			PostId:   post.PostId,
			PhotoUrl: photo,
		})
	}

	// 查询所有的粉丝 uid
	followers := make([]model.Follow, 0)
	err = respository.GetDB().Where("user_id = ?", uid).Find(&followers).Error
	if err != nil {
		response.Response(c, errmsg.UID_IS_NOT_EXIST)
		return
	}

	// 写入 feeds 表
	for _, follower := range followers {
		followerID := follower.FollowID
		feed := model.Feeds{
			UserId:   followerID,
			PostId:   post.PostId,
			SendId:   uid,
			PostType: model.POST_TYPE,
		}
		err = respository.GetDB().Create(&feed).Error
		if err != nil {
			response.Response(c, errmsg.FEED_WRITE_ERROR)
			return
		}
	}

	// 修改客户的 relateData 信息
	err = respository.GetDB().Model(&model.UserRelatedData{}).Where("user_id = ?", uid).Update("post_number", gorm.Expr("post_number + ?", 1)).Error
	if err != nil {
		fmt.Println(err.Error())
		response.Response(c, errmsg.RELATEDATA_WRITE_ERROR)
		return
	}

	data := make(map[string]interface{})
	data["postId"] = post.PostId
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}
