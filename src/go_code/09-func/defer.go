package main

import "fmt"

// 函数中的defer用于延时执行：
// 比如打开文件后关闭文件，连接数据库后关闭连接

func sum(n1 int, n2 int) int {
	// 当执行到defer时，会将defer后面的语句压入到独立的defer栈中，暂时不执行
	// 当函数执行完毕后，再从defer栈中，依次从栈顶取出语句执行
	// 同时也将n1，n2的值拷贝入栈
	defer fmt.Println("ok1 n1= ", n1)
	defer fmt.Println("ok2 n2= ", n2)

	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res", res)
}
