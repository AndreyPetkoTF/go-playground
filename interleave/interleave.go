package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(newIsInterleave(
		"abababababababababababababababababababababababababababababababababababababababababababababababababbb",
		"abababababababababababababababababababababababababababababababababababababababababababababababababab",
		"abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
	),
	)

	fmt.Println(time.Since(now).Seconds())
}

var wg = sync.WaitGroup{}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if s1 == "" && s2 == "" && s3 == "" {
		return true
	}

	result := false
	wg.Add(1)
	check(s1, s2, s3, &result)

	wg.Wait()

	return result
}

func check(s1 string, s2 string, s3 string, result *bool) {
	if s1 == "" {
		if s3 == s2 {
			*result = true
		}
		wg.Done()
		return
	}

	if s2 == "" {
		if s3 == s1 {
			*result = true
		}
		wg.Done()
		return
	}

	if s3[0] == s1[0] {
		wg.Add(1)
		go check(s1[1:], s2, s3[1:], result)
	}

	if s3[0] == s2[0] {
		wg.Add(1)
		go check(s1, s2[1:], s3[1:], result)
	}

	wg.Done()
}

func newIsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make(map[int]bool)

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if i == 0 && j == 0 {
				dp[j] = true
			} else if i == 0 {
				dp[j] = dp[j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[j] = dp[j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) || (dp[j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return dp[len(s2)]
}
