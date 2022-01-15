package main

import (
	"fmt"
	"strings"
)

func reverseString(s []byte) string {
	start := 0
	end := len(s) - 1

	for start < end {
		temp := s[start]
		s[start] = s[end]
		s[end] = temp
		start++
		end--
	}
	return string(s)
}

func reverseWords(s string) string {
	new_string := strings.Split(s, " ")
	for i, str := range new_string {
		new_string[i] = reverseString([]byte(str))
	}
	return strings.Join(new_string, " ")
}

func main() {
	s := "Let's take LeetCode contest"
	result := reverseWords(s)
	fmt.Println(result)
}
