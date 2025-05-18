package main

import (
	"fmt"
	"maps"
	"slices"
)

// maps -> hash, object, dict
func main() {
	// key, value store

	m := make(map[string]string)

	// setting an element
	m["name"] = "bharath"

	fmt.Println(m)

	for value := range maps.Values(m) {
		fmt.Printf(" value: %v\n", value)

	}

	var s []int = []int{1, 2, 3}
	var tomap map[int]int = maps.Collect(slices.All(s))
	fmt.Printf("m: %v\n", tomap)
	fmt.Println(len(tomap))

	var m1 map[string]string = map[string]string{"name": "bharath", "age": "20"}
	var m2 map[string]string = map[string]string{"name": "bharath", "age": "20"}
	fmt.Println(maps.Equal(m1, m2)) //compares equality of both key and value
}
