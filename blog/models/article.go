package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"` //gorm是索引
	Tag   Tag `json:"tag"`  //作关联用
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.InstanceSet("createOn",time.Now().Unix())  //返回当前时间戳
	return nil
}
func (article *Article)BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn",time.Now().Unix())
	return nil
}


func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?",id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}


//获取文章总数
func GetArticleTotal(maps interface{})(count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int,pageSize int,maps interface{})(articles []Article){
	db.Preload("Tag").Where(maps) //预加载
	return
}

//获取文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?").First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

//修改文章
func EditArticle(id int,data interface{}) bool {
	db.Model(&Article{}).Where("id = ?").Updates(data)
	return true
}

//添加文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID : data["tag_id"].(int),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	})
	return true
}

//删除文章
func DeleteArticle(id int) bool {
	db.Where("id = ?").Delete(Article{})
	return true
}


