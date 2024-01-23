package main

import "fmt"

func main() {
	fmt.Println("welcome to function in golang")
	greeter()
	// greeterTwo()

	result := adder(3, 5)
	var proResult, myMessage = proAdder(2, 5, 8, 7)

	fmt.Println("Result is : ", result)
	fmt.Println("pro message is : ", myMessage)
	fmt.Println("Pro Result is : ", proResult)

}

func greeter() {
	fmt.Println("Namastey from golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi pro result function"
}

// func greeterTwo() {
// 	fmt.Println("Another Method")
// }
