package main

import "fmt"

// 自定义数据类型的使用：type

func getSum2(n1 int, n2 int) int {
	return n1 + n2
}

// 案例2、给函数取别名
type myFunType func(int, int) int // 此时myFunType就是func(int, int) int类型

func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func main() {
	// 案例1、给int取别名，虽然都是int，但是go认为他们是两个不同的类型
	type myInt int

	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1) // 此处还是需要显示转换
	fmt.Println("num1 = ", num1, "num2 = ", num2)

	// 案例2的case
	res := myFun2(getSum2, 100, 40)
	fmt.Println("res = ", res)

}
