package main

import "fmt"

// 累加器
func AddUpper() func(int) int {
	var n int = 10
	// return匿名函数，引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包
	// 理解成：闭包是类，函数是操作，n是字段，函数和它使用到的n形成了闭包
	// 返回的匿名函数和匿名函数在外部使用到的变量形成了闭包
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {

	// 闭包

	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13

}
