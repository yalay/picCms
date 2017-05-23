package api

import (
	"controllers"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

const (
	kActionTypeView = "view"
	kActionTypeUp   = "up"
	kActionTypeDown = "down"
)

type SocialController struct {
	beego.Controller
}

func (c *SocialController) Post() {
	articleStrId := c.Ctx.Input.Param(":id")
	articleIntId, _ := strconv.Atoi(articleStrId)
	if articleIntId == 0 {
		c.Ctx.WriteString("invalid article id")
		return
	}

	action := c.Ctx.Input.Param(":action")
	switch action {
	case kActionTypeView, kActionTypeUp, kActionTypeDown:
		viewData := c.Ctx.GetCookie(action)
		if viewData == "" {
			controllers.IncArticleView(int32(articleIntId))
			c.Ctx.SetCookie(action, articleStrId, beego.BConfig.WebConfig.Session.SessionCookieLifeTime)
		} else {
			var hasViewed bool
			dataFields := strings.Split(viewData, ",")
			for _, viewArticleId := range dataFields {
				if viewArticleId == articleStrId {
					hasViewed = true
					break
				}
			}
			if !hasViewed {
				controllers.IncArticleView(int32(articleIntId))
				viewData += "," + articleStrId
				c.Ctx.SetCookie(action, viewData, beego.BConfig.WebConfig.Session.SessionCookieLifeTime)
			}
		}
	default:
		beego.Warn("unknown action type:" + action)
	}
	c.Ctx.WriteString("OK")
}
