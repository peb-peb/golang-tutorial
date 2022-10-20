package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=31"

func main() {
	fmt.Println("URLs")

	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	// Better way to store RawQuery these info
	qparams := result.Query()
	fmt.Printf("%T\n", qparams)
	fmt.Println(qparams["list"])

	partsOfURL := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/",
	}

	toStringURL := partsOfURL.String()
	fmt.Println(toStringURL)

}
