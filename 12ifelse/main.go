package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to if else in golang")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Num : ")
	input, _ := reader.ReadString('\n')

	num, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if num+1 < 10 {
		fmt.Println("Num is smaller than 10 ", input)
	} else if num+1 > 10 {
		fmt.Println("Num is greater than 1o ", input)
	} else if num+1 == 10 {
		fmt.Println("Num is equal to 10 ", num)
	} else {
		fmt.Println("Num is Not a number ", num)
	}

	//----------------------------------------------------------
	num1 := int(num)

	if num1%2 == 0 {
		fmt.Println("It is even")
	} else {
		fmt.Println("It is odd")
	}

	//-----------------------------------------------------------

	if num2 := num1; num2 < 10 {
		fmt.Println("Num is less than 10..")
	} else {
		fmt.Println("Num is greater than  10..")
	}

}
