package server

import (
	"conf"
	"controllers"
	"controllers/api"

	"github.com/astaxie/beego"
)

func router(adminPath string) {
	// 控制器注册
	beego.Router("/", &api.HomeController{})
	beego.Router("/list-:id([0-9]+).html", &api.CateController{})
	beego.Router("/list-:id([0-9]+)-:page([0-9]+).html", &api.CateController{})
	beego.Router("/post-:id([0-9]+).html", &api.ArticleController{})
	beego.Router("/post-:id([0-9]+)-:page([0-9]+).html", &api.ArticleController{})
	beego.Router("/tag-:tag.html", &api.TagController{})
	beego.Router("/tag-:tag-:page([0-9]+).html", &api.TagController{})
	beego.Router("/social/:id/:action", &api.SocialController{})

	beego.SetViewsPath("views/v3")
	beego.SetStaticPath("/css", "views/v3/css")
	beego.SetStaticPath("/js", "views/v3/js")
	beego.SetStaticPath("/img", "views/v3/img")

	beego.AddFuncMap("func_articles", controllers.GetArticles)
	beego.AddFuncMap("func_cates", controllers.GetCates)
	beego.AddFuncMap("func_time", api.TimeFormat)
	beego.AddFuncMap("func_time2", api.TimeFormat2)
	beego.AddFuncMap("func_articleurl", conf.GetArticleUrl)
	beego.AddFuncMap("func_cateurl", conf.GetCateUrl)
	beego.AddFuncMap("func_tagurl", conf.GetTagUrl)

	// 后台面板路由注册
	beego.Get(adminPath, controllers.AdminHandler)
	beego.Post(adminPath, controllers.AdminHandler)
	beego.Get(adminPath+"/*sub", controllers.AdminHandler)
	beego.Post(adminPath+"/*sub", controllers.AdminHandler)
}

func Run(adminPath string) {
	router(adminPath)
	beego.Run()
}
