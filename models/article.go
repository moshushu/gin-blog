package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	// gorm会通过类名+ID的方式去找到这两个类之间的关联关系
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (a *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (a *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if article.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	// Preload就是一个预加载器，它会执行两条SQL
	// 分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
	// 在 gorm 中，查找不到记录也算一种 “错误”
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func EditArticle(id int, data interface{}) error {
	return db.Model(&Article{}).Where("id = ?", id).Update(data).Error
}

func AddArticle(data map[string]interface{}) error {
	article := &Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}
	return db.Create(&article).Error
}

func DeleteArticle(id int) error {
	return db.Where("id = ?", id).Delete(Article{}).Error
}

// 用cron启动定时清楚
// 硬删除要使用 Unscoped()，这是 GORM 的约定
func CleanAllArticle() error {
	return db.Unscoped().Where("deleted_on != ?", 0).Delete(&Article{}).Error
}
