package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to my pizza Shop..")

	fmt.Println("Enter rating between  1 to 5..")

	// creating reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating : ")
	input, err1 := reader.ReadString('\n')

	if err1 != nil {
		panic(err1)
	} else {
		fmt.Printf("You Entered : %v", input)
	}

	// converting String to  float
	convInput := strings.TrimSpace(input)
	rating, err2 := strconv.ParseFloat(convInput, 64)

	if err2 != nil {
		fmt.Println("some errro")
		panic(err2)
	} else {
		fmt.Println("Rating is : ", rating)
	}
}
