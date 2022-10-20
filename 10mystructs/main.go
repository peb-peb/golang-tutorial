package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	s1 := Student{"Peb", "192454", 2}
	fmt.Printf("%+v\n", s1)
	fmt.Println(s1.Name, s1.Year)
}

type Student struct {
	Name string
	USN  string
	Year int
}
