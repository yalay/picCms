package api

import (
	"controllers"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	cacheKey := controllers.MakeCacheKey(controllers.KcachePrefixHome)
	if cacheData, err := controllers.CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		c.Data["cid"] = int32(0)
		c.Data["webName"] = controllers.GetGconfig("web_name")
		c.Data["webKeywords"] = controllers.GetGconfig("web_keywords")
		c.Data["webDesc"] = controllers.GetGconfig("web_description")
		c.Data["tongji"] = controllers.GetGconfig("web_tongji")
		c.Data["copyright"] = controllers.GetGconfig("web_copyright")
		controllers.CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "index.tpl"
}
