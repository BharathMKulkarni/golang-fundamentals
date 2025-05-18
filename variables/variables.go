package main

import "fmt"

func main() {
	var name string = "golang"
	fmt.Println(name)

	// type inference
	var name2 = "bharat"
	fmt.Println(name2 + "\n")

	var ishuman = true
	fmt.Println(ishuman)

	var age int16 = 30
	fmt.Println(age)

	// shorthand declaration. Infers type automatically
	city := "Bengaluru"
	fmt.Println(city)

	var state string
	state = "karnataka"
	fmt.Println(state)

	var price float32 = 50.5
	fmt.Println(price)

}
