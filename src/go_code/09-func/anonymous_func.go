package main

import "fmt"

var (
	// 3、func1这个全局变量就是一个全局匿名函数
	Func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 【匿名函数的用处：在一个函数中定义另外一个函数】
	// 1、在定义匿名函数的时候就直接调用，这种方式匿名函数只能调用一次

	// 例1 求两数之和
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1 = ", res1)

	// 2、将匿名函数赋值给一个变量
	// a的数据类型就是函数变量,此时我们可以通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 * n2
	}

	res2 := a(10, 30)
	fmt.Println("res2 = ", res2)

	// 3、全局匿名函数的使用
	res4 := Func1(4, 5)
	fmt.Println("res4 = ", res4)
}
