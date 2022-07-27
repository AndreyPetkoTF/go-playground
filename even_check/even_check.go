package main

import (
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	s = strings.Replace(s, "-", "", -1)
	var parts []string
	firstItemLength := len(s) % k

	if firstItemLength != 0 {
		parts = append(parts, s[0:firstItemLength])
	}

	iterations := (len(s) - firstItemLength) / k

	for i := 0; i < iterations; i++ {
		parts = append(parts, s[i*k+firstItemLength:(i+1)*k+firstItemLength])
	}

	return strings.ToUpper(strings.Join(parts, "-"))
}
