package main

import "fmt"

func main() {
	var (
		num1, num2 int
	)
	fmt.Print("Enter first number:")
	fmt.Scan(&num1)
	fmt.Print("Enter second number:")
	fmt.Scan(&num2)
	fmt.Println("Before swapping")
	fmt.Print("First number:")
	fmt.Println(num1)
	fmt.Print(("Second number:"))
	fmt.Println(num2)
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println("After swapping")
	fmt.Print("First number:")
	fmt.Println(num1)
	fmt.Print(("Second number:"))
	fmt.Println(num2)
}
