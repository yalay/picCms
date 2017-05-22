package main

import (
	"controllers"
	"server"
)

const (
	kAdminPath = "/admin"
)

func init() {
	controllers.InitDb()
	controllers.InitAdmin(kAdminPath)
	controllers.InitCache(240)
}

func main() {
	server.Run(kAdminPath)
}

