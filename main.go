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
	// demo.TestChannel()
	// nums := foundation.CreateNums(50, 200, 10)
	// nums = append(nums, 111)
	// foundation.PrintNum("  ", nums)
	// foundation.ChooseSort(nums)
	// foundation.PrintNum("  ", nums)

	//index := foundation.BiranySearch(nums, 111)
	// fmt.Print(index)

	//foundation.TestFrog()

	// foundation.TestFindRepeatNum()

	// isMirrorTree := foundation.TestMirrorTree()
	// fmt.Printf("isMirrorTree: %t \n", isMirrorTree)

	// fmt.Printf("以内的丑数据%d \n ", foundation.FindUglyNum(2000))

	//foundation.TestTwoDimensionArrayFind()

	//Xiuren()

	//runServer()

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
