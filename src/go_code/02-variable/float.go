package main

import "fmt"

func main() {

	// 单精度float32，双精度float64（默认），浮点型的存储分为三部分：符号位+指数位+尾数位
	var price float32 = 98.123
	fmt.Println("Price =", price)

	// 浮点数精度可能损失
	var a float32 = 9.0000901
	var b float64 = 9.0000901
	fmt.Println("a= ", a, "b= ", b)

	// 默认是float64，尽量使用64位
	var c = 1.1
	fmt.Printf("the type of c is: %T", c)

}
