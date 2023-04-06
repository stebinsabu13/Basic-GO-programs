package main

import "fmt"

func main() {
	var (
		arr1, arr2 []int
		size, num  int
	)
	fmt.Print("Enter the size of array")
	fmt.Scan(&size)
	fmt.Print("Enter the first array")
	for i := 0; i < size; i++ {
		fmt.Scan(&num)
		arr1 = append(arr1, num)
	}
	fmt.Print("Enter the second array")
	for i := 0; i < size; i++ {
		fmt.Scan(&num)
		arr2 = append(arr2, num)
	}
	fmt.Println("Before swapping")
	fmt.Print("First array:")
	fmt.Println(arr1)
	fmt.Print(("Second array:"))
	fmt.Println(arr2)
	for i := 0; i < size; i++ {
		temp := arr1[i]
		arr1[i] = arr2[i]
		arr2[i] = temp
	}
	fmt.Println("After swapping")
	fmt.Print("First array:")
	fmt.Println(arr1)
	fmt.Print(("Second array:"))
	fmt.Println(arr2)
}
