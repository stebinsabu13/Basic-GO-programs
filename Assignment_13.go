package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter a string")
	fmt.Scan(&input)
	value := func(input string) bool {
		for i := 0; i < len(input)/2; i++ {
			if input[i] != input[len(input)-1-i] {
				return false
			}
		}
		return true
	}
	if value(strings.ToUpper(input)) == true {
		fmt.Print("its a palindrome")
	} else {
		fmt.Print("Not a palindrome")
	}
}
