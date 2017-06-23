package controllers

import (
	"models"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

var Admin *admin.Admin
var AdminServer http.Handler

type MyAuth struct{}
type MyUser struct {
	Name string
}

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
	Admin.AddResource(&models.Lang{}, &admin.Config{Menu: []string{"资源管理"}})

	Admin.AddResource(&models.Download{}, &admin.Config{Menu: []string{"下载管理"}})

	AdminServer = Admin.NewServeMux(adminPath)
	Admin.SetAuth(&MyAuth{})
}

func AdminHandler(c *context.Context) {
	AdminServer.ServeHTTP(c.ResponseWriter, c.Request)
}

func (*MyAuth) LoginURL(c *admin.Context) string {
	return "/403"
}

func (*MyAuth) LogoutURL(c *admin.Context) string {
	return "/403"
}

func (*MyAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	if userCookie, err := c.Request.Cookie("adminuser"); err == nil && userCookie != nil {
		if userCookie.Value == beego.AppConfig.String("adminuser") {
			return &MyUser{Name: "admin"}
		}
	}
	return nil
}

func (u *MyUser) DisplayName() string {
	return u.Name
}
