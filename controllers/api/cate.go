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

type CateController struct {
	beego.Controller
}

func (c *CateController) Get() {
	cateEngName := c.Ctx.Input.Param(":engname")
	cate := controllers.GetCateByEngName(cateEngName)
	if cate == nil {
		c.Ctx.WriteString("cate id non-exist")
		return
	}

	pageStrId := c.Ctx.Input.Param(":page")
	pageId, _ := strconv.Atoi(pageStrId)
	if pageId == 0 {
		pageId = 1
	}

	curLang := GetLang(c.Ctx.Input.Header("Accept-Language"))
	cacheKey := controllers.MakeCacheKey(controllers.KcachePrefixCate, cateEngName, pageStrId, curLang)
	if cacheData, err := controllers.CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		cateUrl := conf.GetCateUrl(cateEngName)
		pathExt := path.Ext(cateUrl)
		page := &models.Page{
			TotalNum:  ((controllers.GetCateArticleNum(cate.Cid) - 1) / conf.KpageArticleNum) + 1,
			CurNum:    pageId,
			SizeNum:   10,
			UrlPrefix: strings.TrimSuffix(cateUrl, pathExt),
			UrlSuffix: pathExt,
		}
		c.Data["webName"] = controllers.Translate(controllers.GetGconfig("web_name"), curLang)
		c.Data["webKeywords"] = controllers.Translate(controllers.GetGconfig("web_keywords"), curLang)
		c.Data["webDesc"] = controllers.Translate(controllers.GetGconfig("web_description"), curLang)
		c.Data["tongji"] = controllers.GetGconfig("web_tongji")
		c.Data["copyright"] = controllers.GetGconfig("web_copyright")
		c.Data["cid"] = cate.Cid
		c.Data["pageId"] = pageId
		c.Data["cName"] = controllers.Translate(cate.Name, curLang)
		c.Data["cKeywords"] = controllers.Translate(cate.Ckeywords, curLang)
		c.Data["cDesc"] = controllers.Translate(cate.Cdescription, curLang)
		c.Data["cArticles"] = controllers.GetCatePageArticles(cate.Cid, pageId)
		c.Data["pagination"] = page.Html()
		c.Data["lang"] = curLang
		controllers.CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "cate.tpl"
}
