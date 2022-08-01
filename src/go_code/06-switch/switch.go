package main

import "fmt"

func main() {
	var score int = 90
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score >= 70:
		{
			fmt.Println("良好")
		}
	default:
		fmt.Println("不及格")
	}
}
