package main

import "fmt"

const AGE = 30

func main() {
	const name = "bharath"
	fmt.Println(AGE)

	var (
		port = 5000
		host = "localhost"
	)

	fmt.Print(port, host, " true")
}
