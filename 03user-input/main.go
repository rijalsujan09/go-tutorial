package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("welcome to user input in golang..")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you input:")

	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	} else {
		fmt.Println(input)
	}

	fmt.Println("---also---")

	var name string

	fmt.Print("Enter you name: ")
	fmt.Scan(&name)

	fmt.Printf("you name is %s \n", name)
}
