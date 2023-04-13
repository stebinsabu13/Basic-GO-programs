package main

import "fmt"

func main() {
	var arr1 []int
	input := func(arr *[]int) {
		var size, num int
		fmt.Print("Enter the size of array")
		fmt.Scan(&size)
		fmt.Print("Enter the array elements:")
		for i := 0; i < size; i++ {
			fmt.Scan(&num)
			*arr = append(*arr, num)
		}
	}
	input(&arr1)
	output := func(arr []int) {
		fmt.Print("The array is:", arr)
	}
	output(arr1)
}
