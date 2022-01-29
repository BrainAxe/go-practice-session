package main

import "fmt"

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	r := []rune(s)
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	start, end := 0, len(r)-1

	for start < end {
		if !contains(vowels, string(r[start])) {
			start++
		} else if !contains(vowels, string(r[end])) {
			end--
		} else {
			r[start], r[end] = r[end], r[start]
			start++
			end--
		}
	}

	return string(r)
}

func main() {
	// s := "hello"
	// s := "leetcode"
	s := "aA"
	result := reverseVowels(s)
	fmt.Println(result)
}
