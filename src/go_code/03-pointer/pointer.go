package main

import "fmt"

func main() {

	// &: 取地址符
	var i int = 10
	fmt.Println("i的地址 = ", &i)

	// ptr: 指针变量，类型为*int，值为&i
	var ptr *int = &i
	fmt.Printf("ptr = %v\n", ptr)
	// 指针变量本身也有一个地址
	fmt.Printf("ptr的地址 = %v\n", &ptr)

	// 获取指针指向的值，使用星号*
	fmt.Printf("ptr指向的值 = %v\n", *ptr)

}
