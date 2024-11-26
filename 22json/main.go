package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")

	// Perform the JSON Encoding
	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	lcocourses := []course{
		{"ReactJS BootCamp", 299, "learcodeonline", "dqqdq", []string{"Beginner", "js"}},
		{"MERN BootCamp", 199, "learcodeonline", "fff", []string{"Interm", "js"}},
		{"Angular BootCamp", 299, "learcodeonline", "dqwd", nil},
	}

	// package this data as JSON data

	finalJSON, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS BootCamp",
			"Price": 299,
			"website": "learcodeonline",
			"tags": ["Beginner","js"]
		}
	`)

	var lcocourse course
	isValid := json.Valid(jsonDataFromWeb)
	if isValid {
		fmt.Println("JSON is Valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON was not Valid")
	}

	// cases where data is needed to be added to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, val := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", key, val, val)
	}
}
