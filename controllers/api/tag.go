package api

import (
	"conf"
	"models"
	"path"
	"strconv"
	"strings"

	"controllers"
	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

func (c *TagController) Get() {
	tag := c.Ctx.Input.Param(":tag")
	pageStrId := c.Ctx.Input.Param(":page")

	curLang := GetLang(c.Ctx.Input.Header("Accept-Language"))
	curLang = "en"
	cacheKey := controllers.MakeCacheKey(controllers.KcachePrefixTag, tag, pageStrId, curLang)
	if cacheData, err := controllers.CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		pageId, _ := strconv.Atoi(pageStrId)
		if pageId == 0 {
			pageId = 1
		}

		tagUrl := conf.GetTagUrl(tag)
		pathExt := path.Ext(tagUrl)
		page := &models.Page{
			TotalNum:  ((controllers.GetTagArticleNum(tag) - 1) / conf.KpageArticleNum) + 1,
			CurNum:    pageId,
			SizeNum:   10,
			UrlPrefix: strings.TrimSuffix(tagUrl, pathExt),
			UrlSuffix: pathExt,
		}
		c.Data["webName"] = controllers.Translate(controllers.GetGconfig("web_name"), curLang)
		c.Data["webKeywords"] = controllers.Translate(controllers.GetGconfig("web_keywords"), curLang)
		c.Data["webDesc"] = controllers.Translate(controllers.GetGconfig("web_description"), curLang)
		c.Data["tongji"] = controllers.GetGconfig("web_tongji")
		c.Data["copyright"] = controllers.GetGconfig("web_copyright")
		c.Data["cid"] = int32(0)
		c.Data["tag"] = controllers.Translate(tag, curLang)
		c.Data["pageId"] = pageId
		c.Data["tArticles"] = controllers.GetTagPageArticles(tag, pageId)
		c.Data["pagination"] = page.Html()
		c.Data["lang"] = curLang
		controllers.CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "tag.tpl"
}
