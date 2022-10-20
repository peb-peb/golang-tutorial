package main

import "fmt"

func main() {
	fmt.Println("IF ELSE")

	n := 10

	if n == 5 {
		fmt.Println("n is 5")
	} else if n < 5 {
		fmt.Println("n is less than 5")
	} else {
		fmt.Println("n is greater than 5")
	}

	// if err != nil {
	// 	// body
	// }

	if num := 4; num < 5 {
		fmt.Println("VALID!")
	}
}
