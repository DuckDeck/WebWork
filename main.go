package main

import (
	"WebWork/core"
	"WebWork/global"
	"WebWork/initialize"
)

func main() {
	runServer()
}

func runServer() {
	initialize.Mysql()
	defer global.G_DB.Close()
	initialize.DBTables()
	core.RunServer()
}
