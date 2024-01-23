package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct in go lang")

	user1 := User{1, "Sujan Rijal", "rijal@test.io", 25}
	fmt.Println(user1)

	// accessing specific element from struct

	fmt.Println(user1.Name, " = ", user1.Age)

}

// creating struct

type User struct {
	Id    int
	Name  string
	Email string
	Age   int
}
