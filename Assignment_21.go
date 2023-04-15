package main

import "fmt"

func main() {
	var (
		ar, arr     []int
		input, size int
	)
	fmt.Print("Enter the size of the array:")
	fmt.Scan(&size)
	fmt.Print("Enter the array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&input)
		ar = append(ar, input)
	}
	for i := 0; i < size-1; i++ {
		input = ar[i] * ar[i+1]
		arr = append(arr, input)
	}
	fmt.Print(arr)
}
