package main

import "fmt"

func counter() func() int {
	var num int = 0 // variable context persists between the calls

	// this function is a closure because it accesses variable in it's surrounding scope
	return func() int {
		num++
		return num
	}
}

func main() {
	next := counter()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
