package main

import "fmt"

func main() {
	fmt.Println("POINTERS")

	// var ptr *int
	// fmt.Println(ptr)

	n := 5
	var ptr *int = &n
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr += 5
	fmt.Println(n)
}
