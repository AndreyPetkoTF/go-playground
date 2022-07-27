package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	res := double([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	fmt.Println(res)
}

func double(items []int) []int {
	values := make(chan int, len(items))
	var wg sync.WaitGroup

	wg.Add(len(items))
	for i, _ := range items {
		go func(i int) {
			defer wg.Done()
			value := items[i] * 2
			values <- value
		}(i)
	}

	go func() {
		wg.Wait()
		close(values)
	}()

	var result []int

	for {
		value, ok := <-values
		if !ok {
			break
		}

		result = append(result, value)
	}

	sort.Ints(result)

	return result
}
