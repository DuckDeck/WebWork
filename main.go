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
	nums := foundation.CreateNums(50, 200, 5)
	foundation.PrintNum("_", nums)
	foundation.BubbleSort(nums)
	foundation.PrintNum("_", nums)

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
