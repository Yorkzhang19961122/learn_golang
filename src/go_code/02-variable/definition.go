package main

import "fmt"

/*
 变量申明
*/

// 全局变量的申明
var p1 = 100
var p2 = 200
var name4 = "jack"

// 也可以一起申明全局变量
var (
	x1    = 10
	x2    = 20
	name5 = "mary"
)

func main() {

	// 一、单变量申明
	// 1、指定变量类型，若不赋予初值，则使用默认值
	var i int
	fmt.Println("i = ", i)
	// 2、根据值自动推导
	var num = 10.11
	fmt.Println("num = ", num)
	// 3、省略var，使用 :=
	name1 := "Tom"
	fmt.Println("my name1 = ", name1)

	// 二、多变量申明
	// 1、多个同类型变量申明
	var n1, n2, n3 int
	n1 = 10
	n2 = 20
	n3 = 30
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)
	// 2、多个不同类型变量申明
	var n4, name2, n5 = 100, "Tom", 888
	fmt.Println("n4 =", n4, "name2 =", name2, "n5 =", n5)
	// 3、多变量申明之使用类型推导
	n6, name3, n7 := 12, "Tom", 11
	fmt.Println("n6 =", n6, "name3 =", name3, "n7 =", n7)

	// 三、全局变量
	fmt.Println("p1 =", p1, "p2 =", p2, "name4 =", name4) // 单独申明
	fmt.Println("x1 =", x1, "x2 =", x2, "name5 =", name5) // 一起申明

}
