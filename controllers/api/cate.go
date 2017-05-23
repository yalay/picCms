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
	cateOriId := c.Ctx.Input.Param(":id")
	pageStrId := c.Ctx.Input.Param(":page")

	cateIntId, _ := strconv.Atoi(cateOriId)
	pageId, _ := strconv.Atoi(pageStrId)
	if cateIntId == 0 {
		c.Ctx.WriteString("invalid cate id")
		return
	}

	cateId := int32(cateIntId)
	cate := controllers.GetCate(cateId)
	if cate == nil {
		c.Ctx.WriteString("cate id non-exist")
		return
	}

	if pageId == 0 {
		pageId = 1
	}

	cacheKey := controllers.MakeCacheKey(controllers.KcachePrefixCate, cateOriId, pageStrId)
	if cacheData, err := controllers.CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		cateUrl := conf.GetCateUrl(cateId)
		pathExt := path.Ext(cateUrl)
		page := &models.Page{
			TotalNum:  ((controllers.GetCateArticleNum(cateId) - 1) / conf.KpageArticleNum) + 1,
			CurNum:    pageId,
			SizeNum:   10,
			UrlPrefix: strings.TrimSuffix(cateUrl, pathExt),
			UrlSuffix: pathExt,
		}
		c.Data["webName"] = controllers.GetGconfig("web_name")
		c.Data["webKeywords"] = controllers.GetGconfig("web_keywords")
		c.Data["webDesc"] = controllers.GetGconfig("web_description")
		c.Data["cid"] = cateId
		c.Data["pageId"] = pageId
		c.Data["cName"] = cate.Name
		c.Data["cKeywords"] = cate.Ckeywords
		c.Data["cDesc"] = cate.Cdescription
		c.Data["cArticles"] = controllers.GetCatePageArticles(cateId, pageId)
		c.Data["pagination"] = page.Html()
		controllers.CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "cate.tpl"
}
