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

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	topicEngName := c.Ctx.Input.Param(":engname")
	topic := controllers.GetTopicByEngName(topicEngName)
	if topic == nil {
		c.Ctx.WriteString("topic id non-exist")
		return
	}

	pageStrId := c.Ctx.Input.Param(":page")
	pageId, _ := strconv.Atoi(pageStrId)
	if pageId == 0 {
		pageId = 1
	}

	curLang := GetLang(c.Ctx.Input.Header("Accept-Language"))
	cacheKey := controllers.MakeCacheKey(controllers.KcachePrefixTopic, topicEngName, pageStrId, curLang)
	if cacheData, err := controllers.CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		topicUrl := conf.GetTopicUrl(topicEngName)
		pathExt := path.Ext(topicUrl)
		page := &models.Page{
			TotalNum:  ((controllers.GetTopicArticleNum(topic.Tid) - 1) / conf.KpageArticleNum) + 1,
			CurNum:    pageId,
			SizeNum:   10,
			UrlPrefix: strings.TrimSuffix(topicUrl, pathExt),
			UrlSuffix: pathExt,
		}
		c.Data["webName"] = controllers.Translate(controllers.GetGconfig("web_name"), curLang)
		c.Data["webKeywords"] = controllers.Translate(controllers.GetGconfig("web_keywords"), curLang)
		c.Data["webDesc"] = controllers.Translate(controllers.GetGconfig("web_description"), curLang)
		c.Data["tongji"] = controllers.GetGconfig("web_tongji")
		c.Data["copyright"] = controllers.GetGconfig("web_copyright")
		c.Data["cid"] = 90
		c.Data["tid"] = topic.Tid
		c.Data["pageId"] = pageId
		c.Data["tName"] = controllers.Translate(topic.Name, curLang)
		c.Data["tTitle"] = controllers.Translate(topic.Stitle, curLang)
		c.Data["tCover"] = topic.Cover
		c.Data["tKeywords"] = controllers.Translate(topic.Skeywords, curLang)
		c.Data["tDesc"] = controllers.Translate(topic.Sdescription, curLang)
		c.Data["tContent"] = controllers.Translate(topic.Content, curLang)
		c.Data["tArticles"] = controllers.GetTopicPageArticles(topic.Tid, pageId)
		c.Data["pagination"] = page.Html()
		c.Data["lang"] = curLang
		controllers.CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "topic.tpl"
}
