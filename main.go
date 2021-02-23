package main

import (
	"WebWork/core"
	"WebWork/demo"
	"WebWork/foundation"
	"WebWork/global"
	"WebWork/initialize"
)

func main() {
	//testPointer()
	//testSlice()
	//demo.TestReflect()
	// demo.TestChannel()
	foundation.NQueen()
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
