package main

import "fmt"

func firstPalindrome(words []string) string {

	word := ""
	for _, s := range words {
		// fmt.Println(s)
		start, end := 0, len(s)-1
		isValid := true

		for start < end {
			if s[start] == s[end] {
				start++
				end--
			} else {
				isValid = false
				break
			}
		}

		if isValid == true {
			word = s
			break
		}
	}

	return word
}

func main() {
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	result := firstPalindrome(words)
	fmt.Println(result)
}
