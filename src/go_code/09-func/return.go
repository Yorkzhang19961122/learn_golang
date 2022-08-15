package main

import "fmt"

// 支持对函数返回值命名，函数返回时直接return即可

func getSum3(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	res1, res2 := getSum3(10, 20)
	fmt.Println("res1 = ", res1, " res2 = ", res2)
}
