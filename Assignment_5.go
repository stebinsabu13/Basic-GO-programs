package main

import "fmt"

func main() {
	var total_mark float32
	fmt.Print("Enter the total mark out of 100:")
	fmt.Scan(&total_mark)
	if total_mark >= 90 {
		fmt.Println("Grade: A")
	} else if total_mark >= 80 && total_mark < 90 {
		fmt.Println("Grade: B")
	} else if total_mark >= 70 && total_mark < 80 {
		fmt.Println("Grade: C")
	} else if total_mark >= 60 && total_mark < 70 {
		fmt.Println("Grade: D")
	} else if total_mark >= 50 && total_mark < 60 {
		fmt.Println("Grade: E")
	} else {
		fmt.Println("Failed")
	}
}
