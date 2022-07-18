package main

import "fmt"

func main() {

	// golang中的进制
	// 二进制输出
	var i int = 5
	fmt.Printf("i = %b\n", i)
	// 八进制
	var j int = 013
	fmt.Println("j = ", j)
	// 十六进制
	var k int = 0x11
	fmt.Println("k = ", k)
}
