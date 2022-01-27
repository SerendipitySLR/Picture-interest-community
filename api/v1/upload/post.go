package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path"
	errmsg "ptc/internal/ermsg"
	"ptc/internal/response"
	"strconv"
	"time"
)

func PostImage(c *gin.Context) {
	file, err := c.FormFile("postImage")
	if err != nil {
		response.Response(c, errmsg.FROM_FORMAT_ERROR)
		return
	}
	timeUnix := strconv.FormatInt(time.Now().Unix(), 10)
	//组成新文件名
	filepath := path.Join(viper.GetString("server.PostImages"), timeUnix+file.Filename)
	fmt.Println(filepath)
	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		response.Response(c, errmsg.FILE_SAVE_ERROR)
		return
	}
	data := make(map[string]interface{})
	data["url"] = "/" + filepath
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}
