package main

import "fmt"

func main() {
	var (
		written, lab, assign float32
	)
	fmt.Print("Enter the mark for written test:")
	fmt.Scan(&written)
	fmt.Print("Enter the mark for lab exam:")
	fmt.Scan(&lab)
	fmt.Print("Enter the mark for assignment:")
	fmt.Scan(&assign)
	grade := (written*70)/100 + (lab*20)/100 + (assign*10)/100
	fmt.Println("Overall Grade=", grade)
}
