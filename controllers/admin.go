package controllers

import (
	"models"
	"net/http"

	"github.com/qor/qor"
	"github.com/qor/admin"
)

var Admin *admin.Admin
var AdminServer  http.Handler

func InitAdmin(adminPath string) {
	Admin = admin.New(&qor.Config{DB: DB})
	Admin.AddResource(&models.Article{})
	Admin.AddResource(&models.Attach{})
	Admin.AddResource(&models.Cate{})
	Admin.AddResource(&models.Tags{})

	AdminServer = Admin.NewServeMux(adminPath)
}
