package main

import "fmt"

func main() {
	fmt.Println("struct in golang")
	// No inheritance in golang; No super no parent

	user := User{"sujan", "rijalsujan09@gmail.com", true, 25}
	fmt.Println(user)
	fmt.Println("---")
	fmt.Printf("hitesh  details are : %+v\n", user)
	fmt.Printf("Name id %v and email is %v.\n", user.Name, user.Email)

	user.GetStatus()
	user.NewMail()

	fmt.Printf("Name id %v and email is %v.\n", user.Name, user.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int is not exporatable
}

func (u User) GetStatus() {
	fmt.Println("---")
	fmt.Println("Is user Active : ", u.Status)
}

func (u User) NewMail() {
	fmt.Println("---")
	u.Email = "test@go.dev"
	fmt.Println("New Email is : ", u.Email)
}
