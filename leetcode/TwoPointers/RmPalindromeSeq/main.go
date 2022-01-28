package main

import "fmt"

func checkPalindrome(s string, start int, end int) bool {

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func removePalindromeSub(s string) int {
	start, end := 0, len(s)-1

	for start < end {
		if s[start] != s[end] {
			return 2
		}
		start++
		end--
	}

	return 1
}

func main() {
	s := "abb"
	result := removePalindromeSub(s)
	fmt.Println(result)
}
