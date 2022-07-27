package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

func main() {
	//chan1 := make(chan int)
	//chan2 := make(chan int)

	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := asChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)

	for v := range mergeChannels(a, b, c) {
		fmt.Println(v)
	}
}

func mergeChannels(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		c := c
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}

			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func mergeChannelsFull(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(<-chan int) {
			for v := range c {
				out <- v
			}

			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func mergeChannelsWithReflect(...<-chan int) <-chan int {
	out := make(chan int)
	var cases []reflect.SelectCase

	for _, c := range cases {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		})
	}

	for {
		i, v, ok := reflect.Select(cases)
		if !ok {
			cases = append(cases[:i], cases[i+1:]...)
			continue
		}

		out <- v.Interface().(int)
	}

	return out
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()

	return c
}
