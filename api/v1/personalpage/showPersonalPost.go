package personalpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"sort"
)

func ShowPersonalPost(c *gin.Context) {

	//通过Uid来找用户发布的帖子
	userId := c.Query("UserId")
	db := respository.GetDB()
	//fmt.Println(userId)
	var posts []model.Post
	db.Where("publisher_id=?", userId).Find(&posts)
	//fmt.Println(posts)
	//接着对于posts中结构体按照UpdatedAt来排序
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].UpdatedAt.After(posts[j].UpdatedAt)
	})
	data := make(map[string]interface{})
	data["PageList"] = posts
	response.ResponseWithData(c, 200, data)
}
