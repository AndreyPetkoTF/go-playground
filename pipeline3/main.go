package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}

	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for value := range in {
		out <- value * value
	}
	close(out)
}

func printer(in <-chan int) {
	for value := range in {
		fmt.Println(value)
	}
}
