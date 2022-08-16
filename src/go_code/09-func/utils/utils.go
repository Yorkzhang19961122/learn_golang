package utils

import "fmt"

var Age int
var Name string

func init() {
	fmt.Println("utils.go中的init()的函数")
	Age = 82
	Name = "张三"
}
