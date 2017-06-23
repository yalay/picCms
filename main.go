package main

import (
	"controllers"
	"server"
)

func init() {
	controllers.InitDb()
	controllers.InitCache(240)
}

func main() {
	server.Run()
}
