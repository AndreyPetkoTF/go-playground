package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

/**
go run main.go // from console to create file in this directory
go tool trace trace.out
*/

func main() {
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	wg := sync.WaitGroup{}

	hellWorld()
	res := 0
	wg.Add(1)
	go func() {
		for i := 0; i < 10000000000; i++ {
			res += i * i
		}
		fmt.Println(res)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10000000000; i++ {
			res += i * i
		}
		fmt.Println(res)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10000000000; i++ {
			res += i * i
		}
		fmt.Println(res)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(res)

}

func hellWorld() {
	res := 0
	for i := 0; i < 100000; i++ {
		res += i * i
	}
}
