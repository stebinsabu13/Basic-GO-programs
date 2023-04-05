package main

import "fmt"

func main() {
	var (
		R, n float32
	)
	var p int
	fmt.Println("enter the rate and number of years and the principal amount respectively")
	fmt.Scan(&R, &n, &p)
	SI := (float32(p) * n * R) / 100
	fmt.Println("The Simple Interest is:", SI)
}
