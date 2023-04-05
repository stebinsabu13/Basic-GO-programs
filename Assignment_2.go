package main

import "fmt"

func main() {
	var (
		num1, sum float32
	)
	var num2 int
	fmt.Println("Enter two numbers")
	fmt.Scan(&num1, &num2)
	sum = float32(num2) + num1
	fmt.Printf("Sum of %f and %d is %v", num1, num2, sum)
}
