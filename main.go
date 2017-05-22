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
}

func main() {
	server.Run(kAdminPath)
}

