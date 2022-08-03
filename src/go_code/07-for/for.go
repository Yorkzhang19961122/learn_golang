package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	var i int = 0
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello ", i)
		i++
	}
}
