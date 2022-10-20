package main

import "fmt"

const TempVar bool = true // Public

func main() {
	var name string = "peb"
	fmt.Println(name)

	var age uint8 = 44
	fmt.Printf("%T\n", age)

	// default values
	var a int
	fmt.Println(a)
	// fmt.Printf("%T\n", a)

	// implicit type
	var s = "some string"
	fmt.Println(s)

	// no var style (walrus operator) (only allowed in methods)
	number := 1000
	fmt.Println(number)

	fmt.Println(TempVar)
}
