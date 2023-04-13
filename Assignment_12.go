package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		arr1      []int
		size, num int
	)
	fmt.Print("Enter the size of array")
	fmt.Scan(&size)
	fmt.Print("Enter the array")
	for i := 0; i < size; i++ {
		fmt.Scan(&num)
		arr1 = append(arr1, num)
	}
	fmt.Println("Array before sorting:")
	fmt.Println(arr1)
	sort.SliceStable(arr1, func(i, j int) bool { return arr1[i] > arr1[j] })
	fmt.Println("Array after sorting:")
	fmt.Print(arr1)
}
