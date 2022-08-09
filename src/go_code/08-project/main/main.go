package main

import (
	"fmt"
	myUtils "learngo/src/go_code/08-project/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'
	result := myUtils.Cal(n1, n2, operator)
	fmt.Println("result = ", result)
}
