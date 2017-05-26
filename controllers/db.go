package controllers

import (
	"bytes"
	"conf"
	"database/sql"
	"log"
	"models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"

	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	retryCount = 3
)

var DB *gorm.DB

func InitDb() {
	var err error
	dbUser := beego.AppConfig.String("mysqluser")
	dbPass := beego.AppConfig.String("mysqlpass")
	dbName := beego.AppConfig.String("mysqldb")
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))
	if err != nil {
		for i := 0; i < retryCount; i++ {
			log.Printf("retry times:%d\n", i+1)
			time.Sleep(time.Minute)
			DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))
			if err == nil {
				break
			}
		}
		if err != nil {
			log.Panic(err)
		}
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

func GetRelatedArticles(id int32, limit int) []*models.Article {
	tagArticles := GetArticlesByTags(GetArticleTags(id), limit, 0)
	if len(tagArticles) <= 1 {
		return nil
	}

	// 删除自己
	var matchedIndex int
	for i, tagArticle := range tagArticles {
		if tagArticle.Id == id {
			matchedIndex = i
			break
		}
	}

	relatedArticles := make([]*models.Article, 0, len(tagArticles)-1)
	if matchedIndex == 0 {
		relatedArticles = tagArticles[1:]
	} else if matchedIndex < len(tagArticles) {
		relatedArticles = append(tagArticles[:matchedIndex], tagArticles[matchedIndex+1:]...)
	} else {
		relatedArticles = tagArticles[:matchedIndex]
	}
	return relatedArticles
}

func GetArticlesByTags(tags []string, limit, offset int) []*models.Article {
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

	var rows *sql.Rows
	var err error
	if limit == 0 {
		if offset == 0 {
			rows, err = DB.Model(&models.Tags{}).Select("article_id, count(*) as cnt").Where("tag in (" + tagsSql.String() + ")").Order("cnt desc").Group("article_id").Rows()
		} else {
			rows, err = DB.Model(&models.Tags{}).Offset(offset).Limit(-1).Select("article_id, count(*) as cnt").Where("tag in (" + tagsSql.String() + ")").Order("cnt desc").Group("article_id").Rows()
		}
	} else {
		rows, err = DB.Model(&models.Tags{}).Offset(offset).Limit(limit).Select("article_id, count(*) as cnt").Where("tag in (" + tagsSql.String() + ")").Order("cnt desc").Group("article_id").Rows()
	}
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	defer rows.Close()
	var tagArticleIds = make([]string, 0)
	for rows.Next() {
		var articleId, count int
		rows.Scan(&articleId, &count)
		tagArticleIds = append(tagArticleIds, strconv.Itoa(articleId))
	}

	if len(tagArticleIds) == 0 {
		return nil
	}

	tagArticles := make([]*models.Article, 0, len(tagArticleIds))
	DB.Where("id in (" + strings.Join(tagArticleIds, ",") + ")").Find(&tagArticles)
	return tagArticles
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

func GetTopHitsArticles(count int) []*models.Article {
	if count <= 0 {
		return nil
	}

	articles := make([]*models.Article, 0, count)
	DB.Limit(count).Order("hits desc").Where("status = 1").Find(&articles)
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

func IncArticleView(articleId int32) {
	DB.Model(&models.Article{}).Where("id = ?", articleId).Update("hits", gorm.Expr("hits + ?", 1))
}

func IncArticleUp(articleId int32) {
	DB.Model(&models.Article{}).Where("id = ?", articleId).Update("up", gorm.Expr("up + ?", 1))
}

func IncArticleDown(articleId int32) {
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

func GetHotTags(count int) []string {
	topArticles := GetTopHitsArticles(count)
	if len(topArticles) == 0 {
		return nil
	}

	articleIds := make([]string, 0, len(topArticles))
	for _, topArticle := range topArticles {
		articleIds = append(articleIds, strconv.Itoa(int(topArticle.Id)))
	}

	rows, err := DB.Limit(count).Select("tag, count(*) as cnt").Where("id in (" + strings.Join(articleIds, ",") + ")").Order("cnt desc").Group("tag").Rows()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	defer rows.Close()
	var tags = make([]string, 0)
	for rows.Next() {
		var tag string
		var count int
		rows.Scan(&tag, &count)
		tags = append(tags, tag)
	}
	return tags
}

func GetTopics() []*models.Topic {
	var topics = make([]*models.Topic, 0)
	DB.Where("status = 1").Find(&topics)
	return topics
}

func GetTopicByEngName(topicEngName string) *models.Topic {
	var topic = &models.Topic{}
	DB.Where("eng_name = ?", topicEngName).Find(topic)
	if topic.Name == "" {
		return nil
	}
	return topic
}

func GetTopicArticleNum(id int32) int {
	var topic = &models.Topic{
		Tid: id,
	}
	DB.First(topic)
	if topic.Skeywords == "" {
		return 0
	}

	topicTags := strings.Split(topic.Skeywords, ",")
	tagsSql := bytes.NewBufferString("")
	for i, tag := range topicTags {
		if i == 0 {
			tagsSql.WriteString(`"` + tag + `"`)
		} else {
			tagsSql.WriteString(`,"` + tag + `"`)
		}
	}

	var count int
	DB.Model(&models.Tags{}).Select("article_id, count(*)").Where("tag in (" + tagsSql.String() + ")").Group("article_id").Count(&count)
	return count
}

func GetTopicPageArticles(topicId int32, pageNum int) []*models.Article {
	if pageNum <= 0 {
		return nil
	}

	var topic = &models.Topic{
		Tid: topicId,
	}
	DB.First(topic)
	if topic.Name == "" {
		return nil
	}

	topicTags := []string{topic.Name, topic.EngName, topic.Stitle}
	return GetArticlesByTags(topicTags, conf.KpageArticleNum, (pageNum-1)*conf.KpageArticleNum)
}

func GetGconfig(key string) string {
	config := &models.Config{
		Name: key,
	}
	DB.First(config)
	return config.Value
}
