package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("GET request from website")
	myurl := "http://localhost:8000/postform"
	PerformPostFormRequest(myurl)
}

func PerformGetRequest(myurl string) {
	response, err := http.Get(myurl)
	checkNilError(err)

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJSONRequest(myurl string) {
	myRequest := strings.NewReader(`
		{
			"country":"india",
			"state":"Bihar",
			"capital":"Patna"
		}
	`)

	response, err := http.Post(myurl, "application/json", myRequest)
	checkNilError(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest(myurl string) {
	data := url.Values{}
	data.Add("name", "peb")
	data.Add("ID", "9999")

	response, err := http.PostForm(myurl, data)
	checkNilError(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
