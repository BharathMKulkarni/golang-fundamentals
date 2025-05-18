package main

import "fmt"

func main() {
	// while kinda loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	fmt.Println("hello")
	// }

	// classic for loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for {
		if j == 1 {
			j++
			continue
		}
		if j == 4 {
			break
		} else {
			fmt.Print(j)
			j++
		}
	}

	for i := range 7 {
		fmt.Println(i)
	}
}
