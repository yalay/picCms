package api

import (
	"conf"
	"controllers"
	"models"
	"path"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

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
	article := controllers.GetArticle(articleId)
	if article == nil {
		c.Ctx.WriteString("article id non-exist")
		return
	}

	pageId, _ := strconv.Atoi(pageStrId)
	if pageId == 0 {
		pageId = 1
	}
	cate := controllers.GetArtilceCate(articleId)
	if cate == nil {
		c.Ctx.WriteString("article no cate")
		return
	}

	attach := controllers.GetArticleAttach(articleId, pageId)
	if attach == nil {
		c.Ctx.WriteString("article no attach")
		return
	}

	cacheKey := controllers.MakeCacheKey(controllers.KcachePrefixArticle, articleStrId, pageStrId)
	if cacheData, err := controllers.CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		attachNum := controllers.GetArticleAttachNum(articleId)
		articleUrl := conf.GetArticleUrl(articleId)
		pathExt := path.Ext(articleUrl)
		page := &models.Page{
			TotalNum:  attachNum,
			CurNum:    pageId,
			SizeNum:   10,
			UrlPrefix: strings.TrimSuffix(articleUrl, pathExt),
			UrlSuffix: pathExt,
		}

		c.Data["webName"] = controllers.GetGconfig("web_name")
		c.Data["webKeywords"] = controllers.GetGconfig("web_keywords")
		c.Data["webDesc"] = controllers.GetGconfig("web_description")
		c.Data["tongji"] = controllers.GetGconfig("web_tongji")
		c.Data["copyright"] = controllers.GetGconfig("web_copyright")
		c.Data["title"] = article.Title
		c.Data["id"] = articleId
		c.Data["pubDate"] = TimeFormat(article.Addtime)
		c.Data["attachNum"] = attachNum
		c.Data["pageId"] = pageId
		c.Data["cName"] = cate.Name
		c.Data["cid"] = cate.Cid
		c.Data["cKeywords"] = cate.Ckeywords
		c.Data["cDesc"] = cate.Cdescription
		c.Data["cUrl"] = conf.GetCateUrl(cate.EngName)
		c.Data["file"] = attach.File
		c.Data["hits"] = article.Hits
		c.Data["preUrl"] = page.PreUrl()
		c.Data["nextUrl"] = page.NextUrl()
		c.Data["pagination"] = page.Html()
		c.Data["relates"] = controllers.GetRelatedArticles(articleId , 9)
		c.Data["tags"] = controllers.GetArticleTags(articleId)
		controllers.CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "article.tpl"
}
