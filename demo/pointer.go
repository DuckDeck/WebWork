package demo

import (
	"flag"
	"fmt"
)

func PointTest() {
	str := "Golang is good"
	strPrt := &str
	fmt.Printf("str type is %T, value is \"%v\",address is %p \n", str, str, &str)
	fmt.Printf("strPtr type is %T,and value is \"%v\" \n", strPrt, strPrt)
	newStr := *strPrt
	fmt.Printf("newStr type is %T, value is \"%v\",address is %p \n", newStr, newStr, &newStr)
	*strPrt = "Java is not good!"
	fmt.Printf("newStr type is %T, value is \"%v\",address is %p \n", newStr, newStr, &newStr)
	fmt.Printf("str type is %T,and value is \"%v\" address is %p \n", str, str, &str)
}

//从命令行里里读取参数
//go run .\main.go -surname="垢" -personName="菲夜" -id=111
//结果是 I am 垢 菲夜,and my id is 111
func FlagPara() {
	surName := flag.String("surname", "王", "你的姓")
	var personName string
	flag.StringVar(&personName, "personName", "小二", "你的名")
	id := flag.Int("id", 0, "你的ID")
	flag.Parse()
	fmt.Printf("I am %v %v,and my id is %v \n", *surName, personName, *id)
}
