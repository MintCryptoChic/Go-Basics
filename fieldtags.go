package main

import (
	"encoding/json"
	"fmt"
)

type salary struct {
	Basic int `json: "basic"`
	HRA   int `json: "hra"`
	TA    int `json: "ta"`
}

type employees struct {
	Name  string   `json: "name"`
	Age   int      `json: "age"`
	City  string   `json: "city"`
	Wages []salary `json: "wages"`
}

func main() {
	json_string := `
	{
		"name": "Meow",
		"age": 5,
		"city": "Miami",
		"wages":[{
				"basic" : 10000,
				"hra" : 2000,
				"ta" : 500
			}]
	}`

	emp1 := new(employees)
	err := json.Unmarshal([]byte(json_string), emp1)
	if err != nil {
		fmt.Print(err, "\n")
	} else {
		fmt.Print(emp1, "\n")
	}
}
