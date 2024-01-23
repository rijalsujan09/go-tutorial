package main

import "fmt"

func main() {
	fmt.Println("pointers in Go-lang")

	num := 24
	fmt.Printf("num value = %v\n", num)
	fmt.Printf("num address = %v\n", &num)

	var numPtr *int = &num
	fmt.Println("Address = ", numPtr)
	fmt.Println("Value = ", *numPtr)

	// lets change the value of num now
	fmt.Println("-------------")

	num = num + 1

	fmt.Printf("num value = %v\n", num)
	fmt.Printf("num address = %v\n", &num)

	fmt.Println("Address = ", numPtr)
	fmt.Println("Value = ", *numPtr)

	fmt.Println("-------------")

	// lets change the value of numPointer
	*numPtr = 26

	fmt.Printf("num value = %v\n", num)
	fmt.Printf("num address = %v\n", &num)

	fmt.Println("Address = ", numPtr)
	fmt.Println("Value = ", *numPtr)

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	cars := []string{"bmw", "toyota", "tata", "marcedes"}

	for index, car := range cars {
		fmt.Printf("index = %v :: value = %v\n", index, car)
	}
}
