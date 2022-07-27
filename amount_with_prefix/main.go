package main

import (
	"fmt"
	"strings"
)

func main() {
	res := amountWithPrefix([]string{"a", "b", "ba", "bax", "bac", "bc", "c"}, "ba")

	fmt.Println(res)
}

func amountWithPrefix(items []string, prefix string) int {
	l := findItem(items, prefix, true)
	r := findItem(items, prefix, false)

	return r - l + 1
}

/**
[l.........m..........r]
m > prefix
c >  ba // true

*/
func findItem(items []string, prefix string, isLeft bool) int {
	l := 0
	r := len(items) - 1

	for l <= r {
		m := (l + r) / 2

		if strings.HasPrefix(items[m], prefix) {
			if isLeft {
				if m == l || !strings.HasPrefix(items[m-1], prefix) {
					return m
				}

				r = m - 1
			} else {
				if m == r || !strings.HasPrefix(items[m+1], prefix) {
					return m
				}

				l = m + 1
			}
		} else {
			if items[m] > prefix {
				r = m - 1
			} else {
				l = m + 1
			}
		}

	}

	return -1
}
