package main

import "fmt"

// receiving reference to the argument
func countUp(num *int) {
	*num++
	fmt.Println("countup: ", *num)
}

func main() {
	num := 1
	countUp(&num) // pass by reference
	countUp(&num) // pass by reference
	fmt.Println("num after countup: ", num)
}
