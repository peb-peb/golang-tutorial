package main

import "fmt"

func main() {
	fmt.Println("ARRAY")

	var list1 [4]int
	fmt.Println(list1)
	fmt.Println(len(list1))

	var list2 = [3]int{1, 2, 3}
	fmt.Println(list2)
	fmt.Println(len(list2))
}
