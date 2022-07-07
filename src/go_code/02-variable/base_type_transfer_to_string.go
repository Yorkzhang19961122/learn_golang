package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 基本数据类型转String的两种方法
	var n1 int = 100
	var n2 float64 = 1.1
	var n3 bool = true
	var n4 byte = 'h'

	// 1、使用fmt.Sprintf("%参数", 表达式) 来将基本数据类型转为字符串
	var str1 string
	str1 = fmt.Sprintf("%d", n1)
	fmt.Printf("str1 type = %T, str1 = %q, n1 type = %T\n", str1, str1, n1)

	var str2 string
	str2 = fmt.Sprintf("%f", n2)
	fmt.Printf("str2 type = %T, str2 = %q, n2 type = %T\n", str2, str2, n2)

	var str3 string
	str3 = fmt.Sprintf("%t", n3)
	fmt.Printf("str3 type = %T, str3 = %q, n3 type = %T\n", str3, str3, n3)

	var str4 string
	str4 = fmt.Sprintf("%c", n4)
	fmt.Printf("str4 type = %T, str4 = %q, n4 type = %T\n", str4, str4, n4)

	// 2、使用strconv包
	var str5 string
	str5 = strconv.FormatInt(int64(n1), 10)
	fmt.Printf("str5 type = %T, str5 = %q\n", str5, str5)

	var str6 string
	str6 = strconv.FormatFloat(n2, 'f', 10, 64)
	fmt.Printf("str6 type = %T, str6 = %q\n", str6, str6)

	var str7 string
	str7 = strconv.FormatBool(n3)
	fmt.Printf("str7 type = %T, str7 = %q\n", str7, str7)

	//var str8 string
	//str8 = strconv.FormatInt(int64(n4), 10)
	//fmt.Printf("str8 type = %T, str8 = %q\n", str8, str8)

}
