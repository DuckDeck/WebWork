package main

import (
	"WebWork/core"
	"WebWork/demo"
	"WebWork/foundation"
	"WebWork/global"
	"WebWork/initialize"
	"fmt"
)

func main() {
	//testPointer()
	//testSlice()
	//demo.TestReflect()
	// demo.TestChannel()
	nums := foundation.CreateNums(50, 200, 10)
	nums = append(nums, 111)
	foundation.PrintNum("  ", nums)
	foundation.ChooseSort(nums)
	foundation.PrintNum("  ", nums)

	index := foundation.BiranySearch(nums, 111)
	fmt.Print(index)
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
