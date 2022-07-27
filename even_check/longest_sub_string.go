package main

//func main() {
//	now := time.Now()
//
//	fmt.Println(lengthOfLongestSubstring("au")) // 3
//
//	fmt.Println("seconds")
//	fmt.Println(time.Since(now).Seconds())
//}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	var longestSubstring int
	for startIndex := 0; startIndex < len(s); startIndex++ {
		chars := make(map[uint8]struct{})

		for i := startIndex; i < len(s); i++ {
			if _, ok := chars[s[i]]; ok {
				break
			}
			chars[s[i]] = struct{}{}
		}

		lenChars := len(chars)
		if lenChars > longestSubstring {
			longestSubstring = lenChars
		}
	}

	return longestSubstring
}
