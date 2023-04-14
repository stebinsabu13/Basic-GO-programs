package main

import "fmt"

func main() {
	var num int = 0
	for i := 0; i < 4; i++ {
		for j := 0; j <= i; j++ {
			num++
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}
