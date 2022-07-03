package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// int8 的范围(-128, 127)
	var j int8 = -128
	j = -128
	fmt.Println("j =", j)

	// uint8 的范围(0, 255)
	var k uint8 = 255
	fmt.Println("k =", k)

	// int(和操作系统有关), uint, rune, byte(0, 255)的使用
	var m int = 8900
	fmt.Println("m =", m)
	var n uint = 1
	fmt.Println("n =", n)
	var a byte = 1
	fmt.Println("a =", a)

	// 查看变量的类型和占用的大小, unsafe.Sizeof()可以返回变量占用的字节数
	var i = 100
	fmt.Printf("i的类型是: %T, i占用的字节数是: %d\n", i, unsafe.Sizeof(i))

}
