package main

import (
	"fmt"
	"time"
)

func main() {
	// use the time library for time related operations
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	// type switch
	// interface{} denotes that the type of the arguement can be anything
	// variable_name.(type) gives the data type of the variable. This can only be used in a type switch.
	whoami := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("other", t)
		}
	}
	whoami(true)
}
