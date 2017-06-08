package api

import (
	"controllers"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	curLang := GetLang(c.Ctx.Input.Header("Accept-Language"))
	cacheKey := controllers.MakeCacheKey(controllers.KcachePrefixHome, curLang)
	if cacheData, err := controllers.CACHE.Get(cacheKey); err == nil {
		c.Data = cacheData.(map[interface{}]interface{})
	} else {
		c.Data["cid"] = int32(0)
		c.Data["webName"] = controllers.Translate(controllers.GetGconfig("web_name"), curLang)
		c.Data["webKeywords"] = controllers.Translate(controllers.GetGconfig("web_keywords"), curLang)
		c.Data["webDesc"] = controllers.TranslateLongText(controllers.GetGconfig("web_description"), curLang)
		c.Data["tongji"] = controllers.GetGconfig("web_tongji")
		c.Data["copyright"] = controllers.GetGconfig("web_copyright")
		c.Data["lang"] = curLang
		controllers.CACHE.Set(cacheKey, c.Data)
	}

	c.TplName = "index.tpl"
}
