package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	// dynamic allocation of memory space if the number of elements in an array is not known prehand

	// uninitialized slices are nil
	var nums []int
	fmt.Println(len(nums))
	fmt.Println(nums == nil) // output: true

	// using make function to initialize slices
	var spaces []int = make([]int, 2)
	fmt.Println(spaces)
	fmt.Println(spaces == nil) // output: false

	// iterating over slices
	var names []string = []string{"bharath", "bhavana"}
	for i, name := range slices.All(names) {
		fmt.Println(i, ":", name)
	}

	// iterating backwards over slices
	for i, name := range slices.Backward(names) {
		fmt.Println(i, ":", name)
	}

	// binary search -> needs a sorted slice in increasing order only to be passed
	randomValues := []int{8, 7, 5, 4, 3, 2, 1}
	n, found := slices.BinarySearch(randomValues, 3)
	fmt.Println("value found or not:", found, n)

	// searching in slices with non primitive data types:
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{"bharath", 25},
		{"bhavana", 22},
		{"sinchana", 25},
	}

	// iterating over people
	for i, person := range slices.All(people) {
		fmt.Println(i, ":", person.name)
	}

	// binary search on people
	i, found := slices.BinarySearchFunc(people, Person{"bharath", 0}, func(a, b Person) int {
		return strings.Compare(a.name, b.name)
	})

	fmt.Println("person found or not: ", i, found)

	home := []int{}
	// for i := 0; i < 10; i++ {
	// 	home[i] = i
	// }
	home = append(home, 1)
	home = append(home, 1)
	home = append(home, 1)
	home = append(home, 1)
	home = append(home, 1)
	fmt.Println(home)

	home2 := make([]int, len(home))
	copy(home2, home)
	fmt.Println(home2)
	fmt.Println(slices.Equal(home, home2))
	home2[2] = 4
	fmt.Println(home2, len(home2), cap(home2))
	home2 = append(home2, 10)
	fmt.Println(home2, len(home2), cap(home2))
	fmt.Println(slices.Equal(home, home2))

	// to clear all elements in a slice
	home = nil

	fmt.Println(home, home2[:3])
	fmt.Println(slices.Contains(home2, 4))

	home2[3] = 5
	hasOdd := slices.ContainsFunc(home2, func(i int) bool {
		return i%2 != 0
	})
	fmt.Println("hasOdd in home2: ", hasOdd)

	// finding the index of an element in a slice
	fmt.Println("Index of 5 in home2: ", home2, slices.IndexFunc(home2, func(i int) bool {
		return i == 5
	}))

	// reversing a slice
	slices.Reverse(home2)
	fmt.Println(home2)

	// sorting a slice
	slices.Sort(home2)
	fmt.Println(home2)

}
