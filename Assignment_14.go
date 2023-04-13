package main

import "fmt"

func main() {
	var (
		arr1, arr2         [][]int
		rows, columns, num int
	)
	fmt.Print("Enter the row and column size")
	fmt.Scan(&rows, &columns)
	fmt.Print("Enter the first array")
	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < columns; j++ {
			fmt.Scan(&num)
			row = append(row, num)
		}
		arr1 = append(arr1, row)
	}
	fmt.Print("First array")
	fmt.Println(arr1)
	fmt.Print("Enter the second array")
	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < columns; j++ {
			fmt.Scan(&num)
			row = append(row, num)
		}
		arr2 = append(arr2, row)
	}
	fmt.Print("Second array")
	fmt.Println(arr2)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			arr1[i][j] += arr2[i][j]
		}
	}
	fmt.Print("Sum of two array")
	fmt.Println(arr1)
}
