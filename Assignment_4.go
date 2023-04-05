package main

import "fmt"

func main() {
	var mark float32
	fmt.Println("Enter the mark of the student out of 100")
	fmt.Scan(&mark)
	if mark >= 50 {
		fmt.Print("Passed")
	} else {
		fmt.Print("Failed")
	}
}
