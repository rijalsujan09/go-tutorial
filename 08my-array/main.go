package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang..")

	// array  decleration stntax
	var fruitList [5]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	// fruitList[2] = "tomato"
	fruitList[3] = "lychee"
	fruitList[4] = "mango"
	// fruitList[5] = "grapes" index out of bound

	fmt.Println("fruitlist -> ", fruitList)

	var employeeList = [5]string{"Ram", "Shyam", "sita", "Gita", "Hari"}
	fmt.Println("empList -> ", employeeList)
}
