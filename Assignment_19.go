package main

import "fmt"

func main() {
	var income float32
	fmt.Print("Enter your income:")
	fmt.Scanln(&income)
	switch {
	case income > 250000 && income <= 500000:
		tax := (5 * income) / 100
		fmt.Println("Tax=", tax)
	case income > 500000 && income <= 1000000:
		tax := (20 * income) / 100
		fmt.Println("Tax=", tax)
	case income > 1000000 && income <= 5000000:
		tax := (30 * income) / 100
		fmt.Println("Tax=", tax)
	default:
		fmt.Println("No Tax")
	}
}
