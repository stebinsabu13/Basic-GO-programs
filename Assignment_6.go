package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter an input:")
	fmt.Scan(&input)
	fmt.Print(("Corresponding day is:"))
	switch input {
	case 1:
		fmt.Print("Sunday")
	case 2:
		fmt.Print("Monday")
	case 3:
		fmt.Print("Tuesday")
	case 4:
		fmt.Print("Wednesday")
	case 5:
		fmt.Print("Thursday")
	case 6:
		fmt.Print("Friday")
	case 7:
		fmt.Print("Saturday")
	default:
		fmt.Print("Invalid Number")
	}
}
