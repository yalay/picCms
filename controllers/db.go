package controllers

import (
	"models"
	"log"
	"strings"
	"conf"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDb() {
	var err error
	DB, err = gorm.Open("mysql", "root:root@/tutu")
	if err != nil {
		log.Panic(err)
	}

	DB.AutoMigrate(
		&models.Article{},
		&models.Attach{},
		&models.Cate{},
		&models.Tags{},
		&models.Config{},
	)
}

func GetArticle(id int32) *models.Article {
	var article = &models.Article{
		Id: id,
	}
	DB.First(article)
	if article.Title == "" {
		return nil
	}
	return article
}

func GetArticleTags(id int32) []string {
	article := GetArticle(id)
	if article == nil || article.Keywords == "" {
		return nil
	}

	return strings.Split(article.Keywords, ",")
}

func GetArtilceCate(id int32) *models.Cate {
	article := GetArticle(id)
	if article == nil {
		return nil
	}
	return GetCate(article.Cid)
}

func GetArticles(cateId int32, star, count int) []*models.Article {
	if count <= 0 {
		return nil
	}

	// 0 表示不限制
	articles := make([]*models.Article, 0, count)
	if cateId == 0 {
		if star == 0 {
			DB.Limit(count).Order("id desc").Where("status = 1").Find(&articles)
		} else {
			DB.Limit(count).Order("id desc").Where("status = 1 AND star = ?", star).Find(&articles)
		}
	} else {
		if star == 0 {
			DB.Limit(count).Order("id desc").Where("status = 1 AND cid = ?", cateId).Find(&articles)
		} else {
			DB.Limit(count).Order("id desc").Where("status = 1 AND cid = ? AND star = ?", cateId, star).Find(&articles)
		}

	}

	return articles
}

func GetCates() []*models.Cate {
	var cates = make([]*models.Cate, 0)
	DB.Order("index").Where("status = 1").Find(&cates)
	return cates
}

func GetCate(cateId int32) *models.Cate {
	var cate = &models.Cate{
		Cid: cateId,
	}
	DB.First(cate)
	if cate.Ctitle == "" {
		return nil
	}
	return cate
}

func GetCateArticleNum(cateId int32) int {
	var count int
	DB.Model(&models.Article{}).Where("cid = ?", cateId).Count(&count)
	return count
}

func GetCatePageArticles(cateId int32, pageNum int) []*models.Article {
	articles := make([]*models.Article, 0)
	if pageNum <= 0 {
		return nil
	}

	if pageNum == 1 {
		DB.Limit(conf.KpageArticleNum).Order("id desc").Where("cid = ?", cateId).Find(&articles)
	} else {
		DB.Offset((pageNum-1)*conf.KpageArticleNum).Limit(conf.KpageArticleNum).Order("id desc").Where("cid = ?", cateId).Find(&articles)
	}

	if len(articles) == 0 {
		return nil
	}
	return articles
}

func GetArticleAttachNum(articleId int32) int {
	var count int
	DB.Model(&models.Attach{}).Where("article_id = ?", articleId).Count(&count)
	return count
}

func GetArticleAttach(articleId int32, pageId int) *models.Attach {
	var attach = &models.Attach{}
	if pageId <= 1 {
		DB.Where("article_id = ?", articleId).Find(attach)
	} else {
		DB.Offset(pageId-1).Limit(1).Where("article_id = ?", articleId).Find(attach)
	}
	if attach.Id == 0 {
		return nil
	}
	return attach
}

func GetArticleAttachs(articleId int32) []*models.Attach {
	attachs := make([]*models.Attach, 0)
	DB.Where("article_id = ?", articleId).Find(&attachs)
	if len(attachs) == 0 {
		return nil
	}
	return attachs
}

func GetArticleAttachUrls(articleId int32) []string {
	attachs := GetArticleAttachs(articleId)
	if len(attachs) == 0 {
		return nil
	}

	attachUrls := make([]string, len(attachs))
	for i, attach := range attachs {
		attachUrls[i] = attach.File
	}
	return attachUrls
}

func GetGconfig(key string) string {
	config := &models.Config{
		Name: key,
	}
	DB.First(config)
	return config.Value
}
