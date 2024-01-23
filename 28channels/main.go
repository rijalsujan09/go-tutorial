package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in goLang - SUJAN RIJAL")

	myCh := make(chan int, 2)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(myCh)
		val, isOpen := <-myCh
		fmt.Println(isOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 0
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
