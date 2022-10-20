package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	m := make(map[string]string)
	m["py"] = "python"
	m["js"] = "javascript"
	fmt.Println(m)

	delete(m, "js")
	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, "->", value)
	}

}
