package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("SWITCH CASE")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Number =", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and You can Open")
		// fallthrough
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and Roll dice again")
	default:
		fmt.Println("This message is not supposed to be printed!")
	}
}
