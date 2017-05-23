package controllers

import (
	"conf"
	"log"
	"models"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"

	"bytes"
	"database/sql"
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

func GetRelatedArticles(id int32) []*models.Article {
	tags := GetArticleTags(id)
	if len(tags) == 0 {
		return nil
	}

	tagsSql := bytes.NewBufferString("")
	for i, tag := range tags {
		if i == 0 {
			tagsSql.WriteString(`"` + tag + `"`)
		} else {
			tagsSql.WriteString(`,"` + tag + `"`)
		}
	}
	rows, err := DB.Model(&models.Tags{}).Select("article_id, count(*) as cnt").Where("tag in (" + tagsSql.String() + ")").Order("cnt desc").Group("article_id").Rows()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	defer rows.Close()
	var relatedArticleIds = make([]string, 0)
	for rows.Next() {
		var articleId, count int
		rows.Scan(&articleId, &count)
		if articleId == int(id) {
			continue
		}
		relatedArticleIds = append(relatedArticleIds, strconv.Itoa(articleId))
	}

	if len(relatedArticleIds) == 0 {
		return nil
	}

	relatedArticles := make([]*models.Article, 0, len(relatedArticleIds))
	DB.Where("id in (" + strings.Join(relatedArticleIds, ",") + ")").Find(&relatedArticles)
	return relatedArticles
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

func GetCateEngNames() []string {
	rows, err := DB.Model(&models.Cate{}).Select("eng_name").Rows()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	engNames := make([]string, 0)
	for rows.Next() {
		var engName string
		rows.Scan(&engName)
		if engName == "" {
			continue
		}
		engNames = append(engNames, engName)
	}
	return engNames
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

func GetCateByEngName(cateEngName string) *models.Cate {
	var cate = &models.Cate{}
	DB.Where("eng_name = ?", cateEngName).Find(cate)
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

func IncArticleView(articleId int32)  {
	DB.Model(&models.Article{}).Where("id = ?", articleId).Update("hits", gorm.Expr("hits + ?", 1))
}

func IncArticleUp(articleId int32)  {
	DB.Model(&models.Article{}).Where("id = ?", articleId).Update("up", gorm.Expr("up + ?", 1))
}

func IncArticleDown(articleId int32)  {
	DB.Model(&models.Article{}).Where("id = ?", articleId).Update("down", gorm.Expr("down + ?", 1))
}

func GetTagArticleNum(tag string) int {
	var count int
	DB.Model(&models.Tags{}).Where("tag = ?", tag).Count(&count)
	return count
}

func GetTagPageArticles(tag string, pageNum int) []*models.Article {
	if pageNum <= 0 {
		return nil
	}

	var rows *sql.Rows
	var err error
	if pageNum == 1 {
		rows, err = DB.Model(&models.Tags{}).Limit(conf.KpageArticleNum).Order("article_id desc").Select("article_id").Where("tag = ?", tag).Rows()
	} else {
		rows, err = DB.Model(&models.Tags{}).Offset((pageNum-1)*conf.KpageArticleNum).Limit(conf.KpageArticleNum).Order("article_id desc").Where("tag = ?", tag).Rows()
	}
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	tagArticleStrIds := make([]string, 0)
	for rows.Next() {
		var articleId int
		rows.Scan(&articleId)
		tagArticleStrIds = append(tagArticleStrIds, strconv.Itoa(articleId))
	}

	if len(tagArticleStrIds) == 0 {
		return nil
	}

	tagArticles := make([]*models.Article, 0, len(tagArticleStrIds))
	DB.Where("id in (" + strings.Join(tagArticleStrIds, ",") + ")").Find(&tagArticles)
	return tagArticles
}

func GetGconfig(key string) string {
	config := &models.Config{
		Name: key,
	}
	DB.First(config)
	return config.Value
}
