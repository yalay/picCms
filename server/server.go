package server

import (
	"conf"
	"controllers"
	"controllers/api"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	kAdminPath = "/admin"
	VERSION    = "v4"
)

func router(adminPath string) {
	// 控制器注册
	beego.Router("/", &api.HomeController{})

	cateEngNames := controllers.GetCateEngNames()
	cateEngNamesExp := strings.Join(cateEngNames, "|")
	beego.Router("/:engname("+cateEngNamesExp+").html", &api.CateController{})
	beego.Router("/:engname("+cateEngNamesExp+")-:page([0-9]+).html", &api.CateController{})

	beego.Router("/article-:id([0-9]+).html", &api.ArticleController{})
	beego.Router("/article-:id([0-9]+)-:page([0-9]+).html", &api.ArticleController{})
	beego.Router(`/tags-:tag([^-\s]+).html`, &api.TagController{})
	beego.Router(`/tags-:tag([^-\s]+)-:page([0-9]+).html`, &api.TagController{})
	beego.Router(`/topic-:engname([^-\s]+).html`, &api.TopicController{})
	beego.Router(`/topic-:engname([^-\s]+)-:page([0-9]+).html`, &api.TopicController{})
	beego.Router("/social/:action/:id", &api.SocialController{})

	beego.SetViewsPath("views/" + VERSION)
	beego.SetStaticPath("/css", "views/"+VERSION+"/css")
	beego.SetStaticPath("/js", "views/"+VERSION+"/js")
	beego.SetStaticPath("/img", "views/"+VERSION+"/img")
	beego.SetStaticPath("/favicon.ico", "views/"+VERSION+"/img/favicon.ico")
	beego.SetStaticPath("/robots.txt", "robots.txt")

	beego.AddFuncMap("func_articles", controllers.GetArticles)
	beego.AddFuncMap("func_cates", controllers.GetAllCates)
	beego.AddFuncMap("func_topics", controllers.GetAllTopics)
	beego.AddFuncMap("func_hottags", controllers.GetHotTags)
	beego.AddFuncMap("func_adsense", controllers.GetAdsense)
	beego.AddFuncMap("func_time", api.TimeFormat)
	beego.AddFuncMap("func_time2", api.TimeFormat2)
	beego.AddFuncMap("func_articleurl", conf.GetArticleUrl)
	beego.AddFuncMap("func_cateurl", conf.GetCateUrl)
	beego.AddFuncMap("func_tagurl", conf.GetTagUrl)
	beego.AddFuncMap("func_topicurl", conf.GetTopicUrl)
	beego.AddFuncMap("func_lang", controllers.Translate)
	beego.AddFuncMap("func_add", Add)

	// 后台面板路由注册
	beego.Get(adminPath, controllers.AdminHandler)
	beego.Post(adminPath, controllers.AdminHandler)
	beego.Get(adminPath+"/*sub", controllers.AdminHandler)
	beego.Post(adminPath+"/*sub", controllers.AdminHandler)
	beego.Get("/403", func(c *context.Context) {
		c.Output.SetStatus(http.StatusForbidden)
	})

	// 快捷工具
	beego.Post("/tool/:type/:id", api.ToolHandler)

	// 兼容旧文章页链接，采用301跳转，seo友好
	beego.Get("/index.:suffix(php|html|htm)", func(c *context.Context) {
		c.Redirect(http.StatusMovedPermanently, c.Input.Site())
	})
	beego.Get("/html/:engname("+cateEngNamesExp+")/:id([\\d-]+).html", func(c *context.Context) {
		articleStrId := c.Input.Param(":id")
		c.Redirect(http.StatusMovedPermanently, c.Input.Site()+conf.GetArticlePageUrl(articleStrId))
	})
}

func Run() {
	controllers.InitAdmin(kAdminPath)
	router(kAdminPath)
	beego.Run()
}

func Add(args ...int) int {
	var total int
	for _, arg := range args {
		total += arg
	}
	return total
}
