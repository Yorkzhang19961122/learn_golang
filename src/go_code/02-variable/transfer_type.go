package main

import "fmt"

func main() {
	// Go在不同变量之间赋值时需要显示转换，也就是说Golang中数据类型不能自动转换!
	var i int32 = 100
	// 将i转为float
	var n1 float32 = float32(i)

	// var n2 int8 = int8(i)  // 高精度转低精度

	fmt.Printf("i = %v, n1 = %v\n", i, n1)

	// 被转换的是变量存储的数据（即值），变量本身的数据类型并没有发生变化
	fmt.Printf("the type of i = %T\n", i)

	// 将int64的大值转换为int8，编译时不会报错，只是转换的结果是按照溢出处理，和我们希望的结果不一样
	var num1 int64 = 10000
	var num2 int8 = int8(num1)
	fmt.Println("num2 = ", num2)
}
