package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slice in Go Lang...")

	// slice decleration

	var fruitList1 = []string{"apple", "banana", "mango"}
	fmt.Println(fruitList1, " ", len(fruitList1))

	var fruitList2 = []string{"grapes", "guva", "lychee"}
	fmt.Println(fruitList2, " ", len(fruitList2))

	fruitList := append(fruitList1, fruitList2...)
	fmt.Println(fruitList, " ", len(fruitList))
	// also
	fruitList = append(fruitList, "pomagranete")
	fmt.Println(fruitList, " ", len(fruitList))

	fmt.Println("----------------------------------------")
	//---------------------------------------------------

	nums := make([]int, 4)

	nums[0] = 5
	nums[1] = 1
	nums[2] = 7
	nums[3] = 3
	// nums[4] = 8  // index out of range
	fmt.Println("nums -> ", nums, " length => ", len(nums))
	fmt.Println("Is Sorted => ", sort.IntsAreSorted(nums))
	nums = append(nums, 1, 4, 2, 6)
	fmt.Println("nums -> ", nums, " length => ", len(nums))
	fmt.Println("Is Sorted => ", sort.IntsAreSorted(nums))
	sort.Ints(nums)
	fmt.Println("nums -> ", nums, " Length =>", len(nums))
	fmt.Println("Is Sorted => ", sort.IntsAreSorted(nums))

	fmt.Println("----------------------------------------")
	//---------------------------------------------------

	var courses = []string{"java", "java-script", "swift", "go-lang", "python", "scala"}
	fmt.Println("courses => ", courses, " length => ", len(courses))

	courses = append(courses[:2], courses[3:]...)
	fmt.Println("courses => ", courses, " length => ", len(courses))
}
