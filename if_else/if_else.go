package main

import "fmt"

func main() {
	admin := true
	hasPermission := false

	if admin && hasPermission {
		fmt.Println("access granted")
	}

	// can declare a variable inside the if construct
	if age := 20; age >= 18 {
		fmt.Println("teenager")
	} else if age < 18 {
		fmt.Println("minor")
	}
}
