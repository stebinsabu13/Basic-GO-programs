package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		input, num1, num2 int
	)
	for {
		fmt.Println("1.Addition\n2.Subtraction\n3.Multiplication\n4.Division\n5.Exit")
		fmt.Println("Enter your choice:")
		fmt.Scan(&input)
		switch input {
		case 1:
			fmt.Print("Enter two numbers for Addition:")
			fmt.Scan(&num1, &num2)
			addition(num1, num2)
		case 2:
			fmt.Print("Enter two numbers for Subtraction:")
			fmt.Scan(&num1, &num2)
			subtraction(num1, num2)
		case 3:
			fmt.Print("Enter two numbers for multiplication:")
			fmt.Scan(&num1, &num2)
			multiplication(num1, num2)
		case 4:
			fmt.Print("Enter two numbers for Division:")
			fmt.Scan(&num1, &num2)
			division(num1, num2)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid number")
		}
	}
}
func addition(a, b int) {
	c := a + b
	fmt.Println("Sum is:", c)
}
func subtraction(a, b int) {
	c := a - b
	fmt.Println("Difference is:", c)
}
func multiplication(a, b int) {
	c := a * b
	fmt.Println("Product is:", c)
}
func division(a, b int) {
	c := a / b
	fmt.Println("Quotient is:", c)
}
