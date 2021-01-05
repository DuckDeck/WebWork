package main

import (
	"WebWork/core"
	"WebWork/demo"
	"WebWork/global"
	"WebWork/initialize"
	"fmt"
	"reflect"
)

func main() {
	//testPointer()
	//testSlice()
	testReflect()
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

func testReflect() {
	typeOfHero := reflect.TypeOf(demo.Hero{})
	fmt.Printf("Hero's type is %s, kind is %s \n", typeOfHero, typeOfHero.Kind())
	typeOfHeroPtr := reflect.TypeOf(&demo.Hero{})
	fmt.Printf("Hero's type is %s, kind is %s \n", typeOfHeroPtr, typeOfHeroPtr.Kind())
	newTypeOfHero := typeOfHeroPtr.Elem()
	fmt.Printf("typeOfHeroPtr elem to TypeOfHero ,hero's type is %s, kind is %s", newTypeOfHero, newTypeOfHero.Kind())

}
