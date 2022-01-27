package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	newString := re.ReplaceAllLiteralString(s, "")
	newString = strings.ReplaceAll(newString, " ", "")
	newString = strings.ToLower(newString)

	start, end := 0, len(newString)-1
	isValid := true

	for start < end {
		if newString[start] == newString[end] {
			start++
			end--
		} else {
			isValid = false
			break
		}
	}

	return isValid
}

func main() {
	// s := "A man, a plan, a canal: Panama"
	s := "race a car"
	result := isPalindrome(s)
	fmt.Println(result)
}
