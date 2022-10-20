package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("USER INPUT")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter string: ")

	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered string is", input)
	fmt.Printf("Type is %T\n", input)
}
