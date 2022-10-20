package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter number between 1 to 5: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Number =", input)

	// increament number by 1 and print
	n, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Incremented Number =", n+1)
	}
}
