package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the number:")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "*", num, "=", i*num)
	}
}
