package main

import (
	"WebWork/core"
	"WebWork/demo"
	"WebWork/global"
	"WebWork/initialize"
)

func main() {
	//testPointer()
	//testSlice()
	//demo.TestReflect()
	demo.TestChannel()
	return

	initialize.Mysql()
	defer global.G_DB.Close()
	initialize.DBTables()

	core.RunServer()
}

func testPointer() {
	demo.PointTest()
	demo.FlagPara()
}

func testSlice() {
	demo.Slice1()
}
