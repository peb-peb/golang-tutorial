package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("SLICES")

	var list1 = []int{1, 2, 3}
	fmt.Printf("Type of list1: %T\n", list1)
	fmt.Println(list1)
	fmt.Println(len(list1))

	list1 = append(list1, 4, 5)
	fmt.Println(list1)
	fmt.Println(len(list1))

	fmt.Println(list1[1:3])
	fmt.Println(len(list1))

	// another way to initialise and declare
	list2 := make([]int, 3)
	list2[0] = 4
	list2[1] = 2
	list2[2] = 6
	// list2[3] = 1 // gives error
	fmt.Println((list2))

	list2 = append(list2, 8, 1, 7)
	fmt.Println(list2)

	sort.Ints(list2)
	fmt.Println(list2)
	fmt.Println(sort.IntsAreSorted(list2))

	// remove a value from the slice based on index
	var list3 = []string{"Apple", "Banana", "Guava", "Pineapple"}
	fmt.Println(list3)
	index := 2
	list3 = append(list3[:index], list3[index+1:]...)
	fmt.Println(list3)

}
