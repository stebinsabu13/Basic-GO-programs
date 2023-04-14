package main

import "fmt"

func main() {
	var input int
	value := func() bool {
		fmt.Print("Enter a number:")
		fmt.Scan(&input)
		for i := 2; i <= input/2; i++ {
			if input%i == 0 {
				return false
			}
		}
		return true
	}
	if value() == true {
		fmt.Print("prime number")
	} else {
		fmt.Print("not a prime")
	}
}
