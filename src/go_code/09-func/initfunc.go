package main

import (
	"fmt"
	"learngo/src/go_code/09-func/utils"
)

var temp = test()

func test() int {
	fmt.Println("test()函数被执行 - 全局变量初始化")
	return 88
}

func init() {
	fmt.Println("initfunc.go中的init()函数被执行")
}
func main() {
	fmt.Println("temp = ", temp)
	fmt.Println("utils.Name = ", utils.Name, "utils.Age = ", utils.Age)
}
