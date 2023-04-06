package main

import "fmt"

func main() {
	var (
		sum, num int
	)
	sum = 0
	fmt.Println("Enter the limt:")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println("Sum of odd numbers is:", sum)
}
