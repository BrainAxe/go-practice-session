package main

import "fmt"

func strStr(haystack string, needle string) int {
	slow := 0
	fast := len(needle)

	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}

	for fast <= len(haystack) {
		if haystack[slow:fast] == needle {
			return slow
		}
		slow++
		fast++
	}

	return -1

}

func main() {
	haystack := "hello"
	needle := "ll"
	result := strStr(haystack, needle)
	fmt.Println(result)
}
