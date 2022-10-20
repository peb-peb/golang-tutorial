package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	greeter()

	result := adder(3, 5)
	fmt.Println("3 + 5 =", result)

	proRes, _ := proAdder(1, 2, 3, 4, 5)
	fmt.Println("sum of 1, 2, 3, 4 and 5 =", proRes)
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Result from ProAdder"
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Greeting")
}
