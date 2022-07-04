package main

import "fmt"

func main() {

	// Go的字符串是由字节组成的
	// 单个字符使用byte存储(0-255)

	// 直接输出字符，会输出对应字符的码值
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1= ", c1, "c2= ", c2)
	// 如果需要输出对应的字符，则需要使用格式化输出
	fmt.Printf("c1= %c c2= %c\n", c1, c2)

	var c3 int = '张'
	fmt.Println("\"张\"对应的码值为: ", c3)
	fmt.Printf("c3= %c\n", c3)

	var c4 byte = '\\'
	fmt.Printf("c4= %c\n", c4)

	// 可以直接给变量赋一个数字，然后格式化输出时%c，会输出该数字对应Unicode字符
	var c5 int = 120
	fmt.Printf("c5= %c\n", c5)

	// 字符类型可以进行运算，相当于一个整数，运算时是按照码值进行
	var n1 = 10 + 'a' // 10 + 97 = 107
	fmt.Println("n1 = ", n1)
	fmt.Printf("n1 = %c", n1)

}
