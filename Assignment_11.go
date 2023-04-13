package main

import "fmt"

func main() {
	var (
		arr1      []int
		size, num int
		count     int = 0
	)
	fmt.Print("Enter the size of array")
	fmt.Scan(&size)
	fmt.Print("Enter the array")
	for i := 0; i < size; i++ {
		fmt.Scan(&num)
		arr1 = append(arr1, num)
	}
	for i := 0; i < size; i++ {
		if arr1[i]%2 == 0 {
			count++
		}
	}
	fmt.Println("Number of even number is:", count)
}
