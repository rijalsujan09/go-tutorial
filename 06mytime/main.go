package main

import (
	"fmt"
	"time"
)

// 01-02-2006 15:04:05 Monday

func main() {
	fmt.Println("welcome to Time in GO lang..")

	timeNow := time.Now()
	fmt.Println(timeNow)
	formatedTimeNow := timeNow.Format("01-02-2006 15:04:05 Monday")
	fmt.Println(formatedTimeNow)

	fmt.Println("---")

	// Now lets create time
	createdTime := time.Date(1998, time.August, 01, 9, 9, 9, 9, time.UTC)
	fmt.Println(createdTime)
	formatedTime := createdTime.Local().Format("01-02-2006 15:04:05 Monday")
	fmt.Println(formatedTime)
}
