package main

import "fmt"

func main() {
	var (
		arr1, arr2    [][]int
		rows, columns int
	)
	fmt.Print("Enter the row and column size")
	fmt.Scan(&rows, &columns)
	fmt.Print("Enter the first array")
	getArray(&arr1, rows, columns)
	fmt.Print("First array")
	displayArray(arr1)
	fmt.Print("Enter the second array")
	getArray(&arr2, rows, columns)
	fmt.Print("Second array")
	displayArray(arr2)
	fmt.Print("Sum of two array")
	addArray(arr1, arr2, rows, columns)
	displayArray(arr1)
}
func getArray(arr *[][]int, r, c int) {
	var num int
	for i := 0; i < r; i++ {
		row := []int{}
		for j := 0; j < c; j++ {
			fmt.Scan(&num)
			row = append(row, num)
		}
		*arr = append(*arr, row)
	}
}
func displayArray(arr [][]int) {
	fmt.Println(arr)
}
func addArray(ar, arr [][]int, r, c int) {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ar[i][j] += arr[i][j]
		}
	}
}

// func addArray(ar, arr [][]int, arrr *[][]int, r, c int) {
// 	for i := 0; i < r; i++ {
// 		row := []int{}
// 		for j := 0; j < c; j++ {
// 			sum := ar[i][j] + arr[i][j]
// 			row = append(row, sum)
// 		}
// 		*arrr = append(*arrr, row)
// 	}
// }
