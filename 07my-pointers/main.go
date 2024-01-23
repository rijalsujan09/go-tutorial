package main

import "fmt"

func main() {

	fmt.Println("welcome to pointer in Go Lang..")

	var ptr *int
	fmt.Println("ptr = ", ptr)

	myNum := 25

	pointer := &myNum

	fmt.Println("Pointer = ", pointer)
	fmt.Println("Pointer = ", *pointer)

	*pointer = *pointer + 1
	fmt.Println("New Pointer = ", *pointer)

	var myNewNum = *pointer + 1
	fmt.Println("My New Num = ", myNewNum)
}
