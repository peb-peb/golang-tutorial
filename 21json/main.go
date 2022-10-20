package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string   `json:"name"`
	ID    string   `json:"id"` // `json:"-"`
	Extra []string `json:"cce,omitempty"`
}

func main() {
	fmt.Println("JSON")
	// EndcodeJSON()
	DecodeJSON()
}

// Creating JSON
func EndcodeJSON() {
	studentData := []Student{
		{"student1", "1234", []string{"football", "basketball"}},
		{"student2", "3214", []string{"chess", "debate"}},
		{"student3", "5321", nil},
	}

	response, err := json.MarshalIndent(studentData, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", response)
}

// Consuming JSON
func DecodeJSON() {
	myData := []byte(`
		{
			"name": "student2",
			"id": "3214",
			"cce": [
					"chess",
					"debate"
			]
		}
	`)

	checkValid := json.Valid(myData)

	var stud Student

	if checkValid {
		fmt.Println("JSON is Valid")
		json.Unmarshal(myData, &stud)
		fmt.Printf("%#v\n", stud)
	} else {
		fmt.Println("JSON is Invalid")
	}

	var myKeyVal map[string]interface{}
	json.Unmarshal(myData, &myKeyVal)
	fmt.Printf("%#v\n", myKeyVal)

	for k, v := range myKeyVal {
		fmt.Printf("Key: %v -> Val: %v\n", k, v)
	}

}
