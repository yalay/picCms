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

	cacheKey := controllers.MakeCacheKey(controllers.KcachePrefixTopic, topicEngName, pageStrId)
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
		c.Data["webName"] = controllers.GetGconfig("web_name")
		c.Data["webKeywords"] = controllers.GetGconfig("web_keywords")
		c.Data["webDesc"] = controllers.GetGconfig("web_description")
		c.Data["tongji"] = controllers.GetGconfig("web_tongji")
		c.Data["copyright"] = controllers.GetGconfig("web_copyright")
		c.Data["cid"] = 90
		c.Data["tid"] = topic.Tid
		c.Data["pageId"] = pageId
		c.Data["tName"] = topic.Name
		c.Data["tTitle"] = topic.Stitle
		c.Data["tCover"] = topic.Cover
		c.Data["tKeywords"] = topic.Skeywords
		c.Data["tDesc"] = topic.Sdescription
		c.Data["tContent"] = topic.Content
		c.Data["tArticles"] = controllers.GetTopicPageArticles(topic.Tid, pageId)
		c.Data["pagination"] = page.Html()
		controllers.CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "topic.tpl"
}
