package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	result := ValidParentheses("(())((()())())")
	fmt.Println(result)
}

func ValidParentheses(parens string) bool {
	sum := 0
	for _, i := range parens {
		char := string(i)

		if char == "(" {
			sum++
		} else {
			sum--
		}

		if sum < 0 {
			return false
		}
	}

	return sum == 0
}

func check() {
	fmt.Println(haveInBoth([]int{1, 5, 2, 2, 2, 6, 6}, []int{6, 2, 2, 4, 2}))
	fmt.Println(intersection([]int{1, 5, 2, 2, 2, 6, 6}, []int{6, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2}))

	for i := range randNumsGenerator(10) {
		fmt.Println(i)
	}
}

func intersection(first, second []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, v := range first {
		counter[v]++
	}

	for _, value := range second {
		if count, ok := counter[value]; ok && count > 0 {
			counter[value]--
			result = append(result, value)
		}
	}

	return result
}

func haveInBoth(first, second []int) []int {
	var result []int

	firstCounts := toCountsMap(first)
	secondCounts := toCountsMap(second)

	for k, v := range firstCounts {
		for ks, vs := range secondCounts {
			if k == ks {
				minCount := getMinCount(v, vs)

				for i := 0; i < minCount; i++ {
					result = append(result, k)
				}
			}
		}
	}

	return result
}

func toCountsMap(numbers []int) map[int]int {
	counts := map[int]int{}

	for _, v := range numbers {
		counts[v]++
	}

	return counts
}

func getMinCount(value, anotherValue int) int {
	if value < anotherValue {
		return value
	}

	return anotherValue
}

func randNumsGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(n)
		}
		close(out)
	}()

	return out
}
