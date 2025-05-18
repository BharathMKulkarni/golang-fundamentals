package main

import "fmt"

func main() {

	// arrays of fixed size that are declared at the time of defining/declaring the variable
	// memory optimized and constant time access in golang.

	// zeroed values:
	// int => 0, string => "", bool => false
	var names [4]float64
	names[1] = 12.334
	fmt.Println(names)

	// to declare an array in a single line
	nums := [2]int{1, 2}
	fmt.Println(nums)

	// 2D arrays
	nums2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums2)
}
