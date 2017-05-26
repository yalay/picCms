package controllers

import (
	"models"
	"net/http"

	"github.com/qor/qor"
	"github.com/qor/admin"
	"github.com/astaxie/beego/context"
)

var Admin *admin.Admin
var AdminServer  http.Handler

func InitAdmin(adminPath string) {
	Admin = admin.New(&qor.Config{DB: DB})
	Admin.RegisterViewPath("views/admin")
	Admin.AddResource(&models.Config{}, &admin.Config{Menu: []string{"系统配置"}})

	Admin.AddResource(&models.Article{}, &admin.Config{Menu: []string{"资源管理"}})
	Admin.AddResource(&models.Attach{}, &admin.Config{Menu: []string{"资源管理"}})
	Admin.AddResource(&models.Cate{}, &admin.Config{Menu: []string{"资源管理"}})
	Admin.AddResource(&models.Tags{}, &admin.Config{Menu: []string{"资源管理"}})
	Admin.AddResource(&models.Topic{}, &admin.Config{Menu: []string{"资源管理"}})
	Admin.AddResource(&models.Ad{}, &admin.Config{Menu: []string{"资源管理"}})

	AdminServer = Admin.NewServeMux(adminPath)
}

func AdminHandler(c *context.Context) {
	AdminServer.ServeHTTP(c.ResponseWriter, c.Request)
}
