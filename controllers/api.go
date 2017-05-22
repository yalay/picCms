package controllers

import (
	"conf"
	"models"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	kTimeLayout  = "2006-01-02 15:04:05"
	kTimeLayout2 = "2006-01-02"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	cacheKey := MakeCacheKey(KcachePrefixHome)
	if cacheData, err := CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		c.Data["cid"] = int32(0)
		c.Data["webName"] = GetGconfig("web_name")
		c.Data["webKeywords"] = GetGconfig("web_keywords")
		c.Data["webDesc"] = GetGconfig("web_description")
		CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "index.tpl"
}

type CateController struct {
	beego.Controller
}

func (c *CateController) Get() {
	cateOriId := c.Ctx.Input.Param(":id")
	pageStrId := c.Ctx.Input.Param(":page")

	cateIntId, _ := strconv.Atoi(cateOriId)
	pageId, _ := strconv.Atoi(pageStrId)
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

	cacheKey := MakeCacheKey(KcachePrefixCate, cateOriId, pageStrId)
	if cacheData, err := CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		cateUrl := conf.GetCateUrl(cateId)
		pathExt := path.Ext(cateUrl)
		page := &models.Page{
			TotalNum:  ((GetCateArticleNum(cateId) - 1) / conf.KpageArticleNum) + 1,
			CurNum:    pageId,
			SizeNum:   10,
			UrlPrefix: strings.TrimSuffix(cateUrl, pathExt),
			UrlSuffix: pathExt,
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
		CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "cate.tpl"
}

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	articleStrId := c.Ctx.Input.Param(":id")
	pageStrId := c.Ctx.Input.Param(":page")

	articleIntId, _ := strconv.Atoi(articleStrId)
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

	pageId, _ := strconv.Atoi(pageStrId)
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

	cacheKey := MakeCacheKey(KcachePrefixArticle, articleStrId, pageStrId)
	if cacheData, err := CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		attachNum := GetArticleAttachNum(articleId)
		articleUrl := conf.GetArticleUrl(articleId)
		pathExt := path.Ext(articleUrl)
		page := &models.Page{
			TotalNum:  attachNum,
			CurNum:    pageId,
			SizeNum:   10,
			UrlPrefix: strings.TrimSuffix(articleUrl, pathExt),
			UrlSuffix: pathExt,
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
			return relatedArticles[:9]
		}()
		c.Data["tags"] = GetArticleTags(articleId)
		CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "article.tpl"
}

type TagController struct {
	beego.Controller
}

func (c *TagController) Get() {
	tag := c.Ctx.Input.Param(":tag")
	pageStrId := c.Ctx.Input.Param(":page")

	cacheKey := MakeCacheKey(KcachePrefixTag, tag, pageStrId)
	if cacheData, err := CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		pageId, _ := strconv.Atoi(pageStrId)
		if pageId == 0 {
			pageId = 1
		}

		tagUrl := conf.GetTagUrl(tag)
		pathExt := path.Ext(tagUrl)
		page := &models.Page{
			TotalNum:  ((GetTagArticleNum(tag) - 1) / conf.KpageArticleNum) + 1,
			CurNum:    pageId,
			SizeNum:   10,
			UrlPrefix: strings.TrimSuffix(tagUrl, pathExt),
			UrlSuffix: pathExt,
		}
		c.Data["webName"] = GetGconfig("web_name")
		c.Data["webKeywords"] = GetGconfig("web_keywords")
		c.Data["webDesc"] = GetGconfig("web_description")
		c.Data["cid"] = int32(0)
		c.Data["tag"] = tag
		c.Data["pageId"] = pageId
		c.Data["tArticles"] = GetTagPageArticles(tag, pageId)
		c.Data["pagination"] = page.Html()
		CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "tag.tpl"
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
