package controllers

import (
	"time"
	"strconv"
	"models"
	"conf"
	"path"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

const (
	kTimeLayout  = "2006-01-02 15:04:05"
	kTimeLayout2 = "2006-01-02"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["cid"] = int32(0)
	c.Data["webName"] = GetGconfig("web_name")
	c.Data["webKeywords"] = GetGconfig("web_keywords")
	c.Data["webDesc"] = GetGconfig("web_description")
	c.TplName = "index.tpl"
}

type CateController struct {
	beego.Controller
}

func (c *CateController) Get() {
	cateOriId := c.Ctx.Input.Param(":id")
	pageOriId := c.Ctx.Input.Param(":page")

	cateIntId, _ := strconv.Atoi(cateOriId)
	pageId, _ := strconv.Atoi(pageOriId)
	if cateIntId == 0 {
		c.Ctx.WriteString("invalid cate id")
		return
	}

	cateId := int32(cateIntId)
	cate := GetCate(cateId)
	if cate == nil {
		c.Ctx.WriteString("cate id non-exist")
		return
	}

	if pageId == 0 {
		pageId = 1
	}

	cateUrl := conf.GetCateUrl(cateId)
	pathExt := path.Ext(cateUrl)
	page := &models.Page{
		TotalNum:((GetCateArticleNum(cateId)-1) / conf.KpageArticleNum) + 1,
		CurNum: pageId,
		SizeNum: 10,
		UrlPrefix:strings.TrimSuffix(cateUrl, pathExt),
		UrlSuffix:pathExt,
	}
	c.Data["webName"] = GetGconfig("web_name")
	c.Data["webKeywords"] = GetGconfig("web_keywords")
	c.Data["webDesc"] = GetGconfig("web_description")
	c.Data["cid"] = cateId
	c.Data["pageId"] = pageId
	c.Data["cName"] = cate.Name
	c.Data["cKeywords"] = cate.Ckeywords
	c.Data["cDesc"] = cate.Cdescription
	c.Data["cArticles"] = GetCatePageArticles(cateId, pageId)
	c.Data["pagination"] = page.Html()
	c.TplName = "cate.tpl"
}
type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	articleOriId := c.Ctx.Input.Param(":id")
	pageOriId := c.Ctx.Input.Param(":page")

	articleIntId, _ := strconv.Atoi(articleOriId)
	pageId, _ := strconv.Atoi(pageOriId)
	if articleIntId == 0 {
		c.Ctx.WriteString("invalid article id")
		return
	}

	articleId := int32(articleIntId)
	article := GetArticle(articleId)
	if article == nil {
		c.Ctx.WriteString("article id non-exist")
		return
	}

	if pageId == 0 {
		pageId = 1
	}
	cate := GetArtilceCate(articleId)
	if cate == nil {
		c.Ctx.WriteString("article no cate")
		return
	}

	attach := GetArticleAttach(articleId, pageId)
	if attach == nil {
		c.Ctx.WriteString("article no attach")
		return
	}

	attachNum := GetArticleAttachNum(articleId)
	articleUrl := conf.GetArticleUrl(articleId)
	pathExt := path.Ext(articleUrl)
	page := &models.Page{
		TotalNum:attachNum,
		CurNum: pageId,
		SizeNum: 10,
		UrlPrefix:strings.TrimSuffix(articleUrl, pathExt),
		UrlSuffix:pathExt,
	}

	c.Data["webName"] = GetGconfig("web_name")
	c.Data["webKeywords"] = GetGconfig("web_keywords")
	c.Data["webDesc"] = GetGconfig("web_description")
	c.Data["title"] = article.Title
	c.Data["id"] = articleId
	c.Data["pubDate"] = TimeFormat(article.Addtime)
	c.Data["attachNum"] = attachNum
	c.Data["pageId"] = pageId
	c.Data["cName"] = cate.Name
	c.Data["cid"] = cate.Cid
	c.Data["cKeywords"] = cate.Ckeywords
	c.Data["cDesc"] = cate.Cdescription
	c.Data["cUrl"] = conf.GetCateUrl(cate.Cid)
	c.Data["file"] = attach.File
	c.Data["hits"] = article.Hits
	c.Data["preUrl"] = page.PreUrl()
	c.Data["nextUrl"] = page.NextUrl()
	c.Data["pagination"] = page.Html()
	c.Data["relates"] = func() []*models.Article {
		relatedArticles := GetRelatedArticles(articleId)
		if len(relatedArticles) <= 9 {
			return relatedArticles
		}
		return relatedArticles[:10]
	}
	c.Data["tags"] = GetArticleTags(articleId)
	c.TplName = "article.tpl"
}

func AdminHandler(c *context.Context) {
	AdminServer.ServeHTTP(c.ResponseWriter, c.Request)
}

func TimeFormat(ts int64) string {
	curTime := time.Unix(ts, 0)
	return curTime.Format(kTimeLayout)
}

func TimeFormat2(ts int64) string {
	curTime := time.Unix(ts, 0)
	return curTime.Format(kTimeLayout2)
}

