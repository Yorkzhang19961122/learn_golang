package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str1 string = "北京长城"
	str1 = "杭州西湖"

	fmt.Println("str1 = ", str1)
	fmt.Printf("size of str1 = %d\n", unsafe.Sizeof(str1))

	// 字符串一旦赋值，就不能再修改了，在go中字符串是不可变的
	// var str = "hello"
	// str[0] = 'a'  // 不能修改str的内容，在go中字符串是不可变的

	// 字符串的两种表示形式
	// 1、双引号，会识别转义字符
	var str2 = "abc\nabc"
	fmt.Println(str2)
	// 2、反引号(``)，以字符串的原生形式输出。包括换行和特殊字符，可实现输出源码的效果
	var str3 = `package main
				import (
					"fmt"
					"unsafe"
				)
				func main() {
					var str1 string = "北京长城"
					str1 = "杭州西湖"
				
					fmt.Println("str1 = ", str1)
					fmt.Printf("size of str1 = %d\n", unsafe.Sizeof(str1))
				
					// 字符串一旦赋值，就不能再修改了，在go中字符串是不可变的
					// var str = "hello"
					// str[0] = 'a'  // 不能修改str的内容，在go中字符串是不可变的
				
					// 字符串的两种表示形式
					// 1、双引号，会识别转义字符
					var str2 = "abc\nabc"
					fmt.Println(str2)
					// 2、反引号，以字符串的原生形式输出。包括换行和特殊字符，可实现输出源码的效果
				}`
	fmt.Println(str3)

	// 字符串的拼接方式：用“+”号
	var str4 = "Hello"
	var str5 = str4 + " World!"
	fmt.Println(str5)
	// 当一个字符串的拼接操作很长时，可以换行写，但是加号要留在上面!
	var str6 = "hello" + "hello" + "hello" + "hello" +
		"hello" + "hello" + "hello" + "hello"
	fmt.Println(str6)
}
