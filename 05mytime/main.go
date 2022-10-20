package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("TIME")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdTime := time.Date(2024, time.June, 1, 1, 1, 0, 0, time.UTC)
	fmt.Println(createdTime)
	fmt.Println(createdTime.Format("01-02-2006"))
}
