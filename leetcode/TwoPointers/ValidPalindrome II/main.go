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

func validPalindrome(s string) bool {
	start, end := 0, len(s)-1

	for start < end {
		if s[start] != s[end] {
			fmt.Println(s[start], s[end])
			return checkPalindrome(s, start+1, end) || checkPalindrome(s, start, end-1)
		}
		start++
		end--
	}
	return true
}

func main() {
	s := "abc"
	result := validPalindrome(s)
	fmt.Println(result)
}
