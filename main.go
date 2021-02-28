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
	nums := foundation.CreateNums(50, 200, 10)
	foundation.PrintNum("  ", nums)
	foundation.ChooseSort(nums)
	foundation.PrintNum("  ", nums)

	return

}

func runServer() {
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
