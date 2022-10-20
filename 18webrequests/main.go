package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("WEB REQUESTS")

	const url = "https://lco.dev"

	response, err := http.Get(url)
	checkNilErr(err)

	// fmt.Printf("%T", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content := string(databytes)
	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
