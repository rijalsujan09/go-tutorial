package main

import "fmt"

// jwtToken1 := "xxxx-xxxx-xxxx"   // --> Not allowed outside of the Function
// num := 1234   // --> Not allowed outside of the Function

var jwtToken2 string = "xxxx-xxxx-xxxx"
var num int = 123

// creating CONSTANT
const LOGINTOKEN string = "abcd-efgh-ijkl"

func main() {

	//initilizing variables
	var username string = "sujan"
	fmt.Printf("variable is of type %T and value i :  %s \n", username, username)
	fmt.Println("---")

	var isLoggedIn bool = false
	fmt.Println("Is sujan logged In :", isLoggedIn)
	fmt.Println("---\n")

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable of type : %T \n", smallVal)
	fmt.Println("---\n")

	var smallFloat float32 = 255.455444412555512523525
	fmt.Println(smallFloat)
	fmt.Printf("Variable of type : %T \n", smallFloat)
	fmt.Println("---\n")

	// default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type : %T \n", anothervariable)
	fmt.Println("---\n")

	// implicit type
	var website = "enrichmentfinancialgroup.com"
	fmt.Println(website)
	fmt.Printf("value is of type : %T\n", website)
	fmt.Println("---\n")

	// no var style
	age := 25
	fmt.Println(age)
	fmt.Printf("value is of type : %T\n", age)
	fmt.Println("---\n")

	fmt.Println(LOGINTOKEN)
	// LOGINTOKEN = "1234-5678-9101" // --> CONSTANT is not changable
	fmt.Printf("VALUE is of TYPE : %T\n", LOGINTOKEN)
	fmt.Println("---\n")
}
