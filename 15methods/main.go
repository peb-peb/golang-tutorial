package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	s1 := Student{"Peb", "192454", 2}
	fmt.Printf("%+v\n", s1)

	s1.GetUSN()
	s1.NewUSN()
	fmt.Printf("%+v\n", s1) // USN doesn't change because copy is sent

}

type Student struct {
	Name string
	USN  string
	Year int
}

func (s Student) GetUSN() {
	fmt.Println("USN =", s.USN)
}

func (s Student) NewUSN() {
	s.USN = "111222"
	fmt.Println("New USN =", s.USN)
}
