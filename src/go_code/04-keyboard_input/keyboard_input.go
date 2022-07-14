package main

import "fmt"

func main() {

	// 键盘输入函数
	// 1. scanln()
	var name string
	var age byte
	//var sal float32
	//var isPass bool
	//fmt.Println("请输入姓名：")
	//// 程序执行到	fmt.Scanln(&name)时，程序会停止在这里，等待用户输入，并回车
	//fmt.Scanln(&name)
	//fmt.Println("请输入年龄：")
	//fmt.Scanln(&age)
	//fmt.Println("请输入薪水：")
	//fmt.Scanln(&sal)
	//fmt.Println("请输入是否通过考试：")
	//fmt.Scanln(&isPass)
	//
	//fmt.Printf("名字是 %v\n", name)
	//fmt.Printf("年龄是 %v\n", age)
	//fmt.Printf("薪水是 %v\n", sal)
	//fmt.Printf("是否通过考试 %v\n", isPass)

	// 2. scanf()，可以按指定的格式输入
	fmt.Println("请输入你的姓名、年龄，使用空格断开")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("名字是 %v\n 年龄是 %v\n", name, age)

}
