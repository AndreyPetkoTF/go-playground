package main

import (
	"fmt"
	"sync"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	var wg sync.WaitGroup

	// Generation
	go func() {
		for x := 0; x < 1000; x++ {
			wg.Add(1)
			naturals <- x
		}

		wg.Wait()
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				close(squares)
				break
			}
			wg.Done()
			squares <- x * x
		}
	}()

	for {

		v, ok := <-squares
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}

	fmt.Println("bye")
}
