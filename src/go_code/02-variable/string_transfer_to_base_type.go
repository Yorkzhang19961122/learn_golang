package main

import (
	"fmt"
	"strconv"
)

func main() {

	// string转换成基本数据类型
	// 使用strconv.ParseBool()函数

	// string -> bool
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T, b = %v\n", b, b)

	// string -> int
	var str1 string = "123456"
	var n1 int64
	n1, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("n1 type is %T, n1 = %v\n", n1, n1)

	// string -> float
	var str2 string = "123.456"
	var n2 float64
	n2, _ = strconv.ParseFloat(str2, 64)
	fmt.Printf("n2 type is %T, n2 = %v\n", n2, n2)

}
