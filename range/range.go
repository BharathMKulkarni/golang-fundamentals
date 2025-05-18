package main

import (
	"fmt"
	"slices"
)

func main() {
	var slice []int = []int{10, 20, 30}
	for i, num := range slice {
		fmt.Println(i, num)
	}

	for i, v := range slices.All(slice) {
		fmt.Println(i, v)
	}

	var store map[string]string = map[string]string{
		"name": "bharath",
		"age":  "25",
	}

	for k, value := range store {
		fmt.Println(k, value)
	}

	var s string = "bhƏrath"
	for i, c := range s {
		// i -> staring byte of the character
		// c -> decimal value of the ASCII character
		// string(c) -> gives the string of the decimal value
		fmt.Println(i, string(c))
	}

}
