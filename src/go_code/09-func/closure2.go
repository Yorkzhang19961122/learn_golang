package main

import (
	"fmt"
	"strings"
)

// 闭包的最佳实践
// 编写一个函数，makeSuffix(suffix string) 可以接收一个文件后缀名（如.jpg），并返回一个闭包
// 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（比如.jpg），则返回文件名.jpg，如果有指定的后缀，则原样返回
// 要求使用闭包i的方式完成
// strings.HasSuffix() ，该函数可以判断某个字符串是否有指定的后缀

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//var suffix string = ".jpg"
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 传统的方法实现同样的功能
func makeSuffix2(name string, suffix string) string {
	if !strings.HasSuffix(name, suffix) {
		name += suffix
	}
	return name
}

func main() {

	// 测试makeSuffix的使用
	// 返回的匿名函数和makeSuffix(suffix string)的suffix变量组成了一个闭包，因为返回的函数引用到suffix这个变量
	// 区别于传统的方法，闭包的方式可以让我们不必要每次调用函数都传入后缀名，而闭包可以保留上次引用的值，所以传入一次就可以反复使用
	f := makeSuffix(".jpg")
	fmt.Println("修改后的文件名 = ", f("winter"))
	fmt.Println("修改后的文件名 = ", f("win.jpg"))

	fmt.Println("修改后的文件名 = ", makeSuffix2("cold", ".jpg"))
	fmt.Println("修改后的文件名 = ", makeSuffix2("cool.jpg", ".jpg"))

}
