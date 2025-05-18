package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func divide(a, b int) (float64, error) {
	if b == 0 {
		// use the errors package to throw errors
		return 0, errors.New("can't divide by zero")
	}
	return float64(a) / float64(b), nil
}

func getOldestCalculation() func(m map[string]int) string {
	calculateOldest := func(m map[string]int) string {
		var oldest int = math.MinInt
		var result string
		for key, value := range m {
			if value > oldest {
				oldest = value
				result = key
			}
		}
		return result
	}
	return calculateOldest
}

func main() {
	var result, error = divide(10, 0)
	if error != nil {
		log.Println(error)
	} else {
		fmt.Println(result)
	}

	var m map[string]int = map[string]int{
		"bharath": 25,
		"bhavana": 22,
		"chetan":  27,
		"chintu":  19,
	}
	calcOldest := getOldestCalculation()
	fmt.Println("the oldest is: ", calcOldest(m))

	executeFunc(calcOldest, m)

	// calling a variadic function
	nums := []int{2, 4, 5, 6}
	fmt.Println(summation(nums...), summation(1, 2, 3))
}

func executeFunc(f func(m map[string]int) string, m map[string]int) {
	var result = f(m)
	fmt.Println("the oldest is: ", result)
}

func summation(nums ...int) int {
	var total int = 0
	for _, num := range nums {
		total = total + num
	}
	return total
}
