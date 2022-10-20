package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("FILES")

	content := "This should be written in file"

	file, err := os.Create("./text.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length:", length)

	defer file.Close()

	ReadFile("./text.txt")
}

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename) // data is in bytes
	checkNilError(err)
	fmt.Println(string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
