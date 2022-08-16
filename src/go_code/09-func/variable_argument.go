package main

import "fmt"

// golang的可变参数，CalSum函数求i个整数的和
func CalSum(i int, args ...int) int {
	sum := 0
	for j := 0; j < i; j++ {
		sum += args[j]
	}
	return sum
}

// 通常可以在init()函数中进行一些初始化操作
func init() {
	fmt.Println("init() runing......")
}

func main() {
	res := CalSum(5, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println("res = ", res)
}
