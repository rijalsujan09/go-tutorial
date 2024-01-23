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
	fmt.Println("welcome to JSON in GO lang")
	fmt.Println("---")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	course1 := []course{
		{"Go lang", 999, "youtube", "pass1", []string{"web-dev", "json"}},
		{"Java", 199, "youtube", "pass1", []string{"full-stack", "backend"}},
		{"Python", 299, "youtube", "pass1", nil},
	}

	// finalJson, err := json.Marshal(course1)
	finalJson, err := json.MarshalIndent(course1, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	{
		"coursename": "Go lang",
		"Price": 999,
		"website": "youtube",
		"_": "pass1",
		"tags": [
				"web-dev",
				"json"
		]
	}
	`)

	var myCourse course
	isvalid := json.Valid(jsonDataFromWeb)

	if isvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("the json was not valid")
	}

	// some cases where you just want to  add data to a key  value

	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	fmt.Printf("%#v\n", myData)

	for key, value := range myData {
		fmt.Printf("key is %v : value is %v : type is %T\n", key, value, value)
	}
}
