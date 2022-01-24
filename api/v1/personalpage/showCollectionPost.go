package personalpage

import (
	"github.com/gin-gonic/gin"
	"ptc/internal/model"
	"ptc/internal/response"
	"ptc/internal/respository"
	"sort"
	"strconv"
)

func ShowCollectionPost(c *gin.Context) {
	//参数
	//userId := c.PostForm("UserId")
	collectionId := c.Query("CollectionId")
	db := respository.GetDB()
	var collection model.Collection
	var post model.Post
	var posts []model.Post
	var postsid string //postsid放字符串切片
	var substr1 []byte //用于接收每一个postid
	var substr2 string //用于接收中间值
	var postid int
	var postids []int //postid放每一个postid

	//从先利用collectionId找到collection表,拿出里面的postsid进行切割
	collection.CollectionId, _ = strconv.Atoi(collectionId)
	db.Find(&collection)
	postsid = collection.PostsId
	for i := 0; i < len(postsid); i++ {
		if postsid[i] != ',' {
			substr1 = append(substr1, postsid[i])
		} else {
			substr2 = string(substr1)
			postid, _ = strconv.Atoi(substr2)
			//fmt.Println(postid)
			postids = append(postids, postid)
			//fmt.Println(postids)
			substr1 = append(substr1[:len(substr1)-1], substr1[len(substr1):]...) //将内容置为空
			//fmt.Println(substr1)
		}
	}

	//切割完毕放入postid[]
	//循环
	for i := 0; i < len(postids); i++ {
		post.PostId = postids[i]
		db.Find(&post)
		posts = append(posts, post)
	}

	//帖子按照更新顺序摆放
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].UpdatedAt.After(posts[j].UpdatedAt)
	})
	data := make(map[string]interface{})
	data["PostList"] = posts
	response.ResponseWithData(c, 200, data)

	//c.JSON(http.StatusOK, gin.H{
	//	"post": post,
	//	"like": like,
	//})
}
