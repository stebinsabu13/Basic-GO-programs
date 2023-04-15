package main

import "fmt"

type arr struct {
	value       [][]int
	row, cloumn int
}

func main() {
	var array1 arr
	fmt.Print("Enter the row and column size:")
	fmt.Scan(&array1.row, &array1.cloumn)
	getArray(&array1)
	displayArray(array1)
}
func getArray(array *arr) {
	var num int
	fmt.Print("Enter the array elements:")
	for i := 0; i < array.row; i++ {
		row := []int{}
		for j := 0; j < array.cloumn; j++ {
			fmt.Scan(&num)
			row = append(row, num)
		}
		array.value = append(array.value, row)
	}
}
func displayArray(array arr) {
	fmt.Print("The array is:")
	fmt.Println(array.value)
}
