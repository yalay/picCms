package main

import (
	"controllers"

	"github.com/astaxie/beego"

	_ "routers"
)

func init() {
	controllers.InitDb()
	controllers.InitAdmin("/admin")
}

func main() {
	beego.Run()
}

