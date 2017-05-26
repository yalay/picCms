package server

import (
	"conf"
	"controllers"
	"controllers/api"

	"github.com/astaxie/beego"
	"strings"
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
	beego.Router("/social/:id/:action", &api.SocialController{})

	beego.SetViewsPath("views/v3")
	beego.SetStaticPath("/css", "views/v3/css")
	beego.SetStaticPath("/js", "views/v3/js")
	beego.SetStaticPath("/img", "views/v3/img")
	beego.SetStaticPath("/favicon.ico", "views/v3/img/favicon.ico")
	beego.SetStaticPath("/robots.txt", "robots.txt")

	beego.AddFuncMap("func_articles", controllers.GetArticles)
	beego.AddFuncMap("func_cates", controllers.GetCates)
	beego.AddFuncMap("func_topics", controllers.GetTopics)
	beego.AddFuncMap("func_time", api.TimeFormat)
	beego.AddFuncMap("func_time2", api.TimeFormat2)
	beego.AddFuncMap("func_articleurl", conf.GetArticleUrl)
	beego.AddFuncMap("func_cateurl", conf.GetCateUrl)
	beego.AddFuncMap("func_tagurl", conf.GetTagUrl)
	beego.AddFuncMap("func_topicurl", conf.GetTopicUrl)

	// 后台面板路由注册
	beego.Get(adminPath, controllers.AdminHandler)
	beego.Post(adminPath, controllers.AdminHandler)
	beego.Get(adminPath+"/*sub", controllers.AdminHandler)
	beego.Post(adminPath+"/*sub", controllers.AdminHandler)

	// 兼容旧文章页链接
	beego.Router("/index.:suffix(php|html|htm)", &api.HomeController{})
	beego.Router("/html/:engname("+cateEngNamesExp+")/:id([0-9]+).html", &api.ArticleController{})
	beego.Router("/html/:engname("+cateEngNamesExp+")/:id([0-9]+)-:page([0-9]+).html", &api.ArticleController{})
}

func Run(adminPath string) {
	router(adminPath)
	beego.Run()
}
