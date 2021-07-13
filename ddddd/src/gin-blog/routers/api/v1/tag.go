package v1

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/cryptobyte/asn1"
	"github.com/dearllu/ddddd/src/gin-blog/models"
)



//获取多个文件的标签
func GetTags(c *gin.Context){
}

//新增文章标签
func AddTag(c *gin.Context){
}

//修改文章标签
func EditTag(c *gin.Context){
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?",name)
}