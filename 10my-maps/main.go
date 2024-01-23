package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to Map in Go Lang")

	language := make(map[string]string)
	language["js"] = "javascript"
	language["jv"] = "java"
	language["py"] = "python"
	language["go"] = "GO"
	language["sc"] = "scalla"

	fmt.Println("languages = > ", language)
	fmt.Println("jv", language["jv"])
	fmt.Println("length -> ", len(language))

	fmt.Println("")

	delete(language, "sc")
	fmt.Println("languages = > ", language)
	fmt.Println("length -> ", len(language))

	fmt.Println("")

	for key, value := range language {
		fmt.Printf("%v : %v \n", key, value)
	}
}

// func main() {
// 	fmt.Println("maps in golang")

// 	language := make(map[string]string)
// 	language["js"] = "javascript"
// 	language["rb"] = "ruby"
// 	language["py"] = "python"

// 	fmt.Println("list of all languages : ", language)
// 	fmt.Println("JS sorts for: ", language["js"])

// 	delete(language, "rb")

// 	fmt.Println("list of all languages : ", language)
// 	fmt.Println("---")

// 	// loops are intersting in golang

// 	for key, value := range language {
// 		fmt.Printf("for key %v, value is %v\n", key, value)
// 	}
// }
