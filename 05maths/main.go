package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to Maths in go lang")

	// var myNumOne int = 2
	// var myNumTwo float64 = 4.5
	// fmt.Println("The Sum is : ", myNumOne+int(myNumTwo))

	//random number

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandNum)

}
