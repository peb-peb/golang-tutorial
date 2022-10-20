package main

import "fmt"

func main() {
	fmt.Println("LOOPS")

	// var list1 = []string{"MON", "TUE", "WED", "THR", "FRI", "SAT", "SUN"}

	// for i := 0; i < len(list1); i++ {
	// 	fmt.Println(list1[i])
	// }

	// for i := range list1 {
	// 	fmt.Println(list1[i])
	// }

	// for index, day := range list1 {
	// 	fmt.Printf("Index = %v and Day = %v\n", index, day)
	// }

	n := 1
	for n < 10 {
		if n == 5 {
			// break
			// n++
			// continue
			goto label1
		}
		fmt.Println(n)
		n++
	}

label1:
	fmt.Println("GOT TO LABEL1")

}
