package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	fmt.Println(days)

	for d := len(days) - 1; d >= 0; d-- {
		fmt.Println(days[d])
	}
	fmt.Println("---")

	for i := range days {
		fmt.Println(days[i])
	}

	for index, value := range days {
		fmt.Printf("index [%v] == value [%v]\n", index, value)
	}
	fmt.Println("---")

	rougevalue := 1

	for rougevalue < 10 {

		// if rougevalue == 2 {
		// 	goto myMsg
		// }

		if rougevalue == 5 {
			rougevalue++
			continue
		}
		fmt.Println("value is : ", rougevalue)
		rougevalue++
	}

	// myMsg:
	// 	fmt.Println("hello from sujan")

}
