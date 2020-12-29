package main

import (
	"WebWork/core"
	"WebWork/demo"
	"WebWork/global"
	"WebWork/initialize"
)

func main() {
	demo.PointTest()
	demo.FlagPara()
	return

	initialize.Mysql()
	defer global.G_DB.Close()
	initialize.DBTables()

	core.RunServer()
}
