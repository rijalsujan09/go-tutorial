package main

import "fmt"

func main() {
	defer fmt.Println("Hello Rijal--2")
	fmt.Println("Hello Sujan--1")
	defer fmt.Println("two")
	defer fmt.Println("one")
	myDefer()
}

func myDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
