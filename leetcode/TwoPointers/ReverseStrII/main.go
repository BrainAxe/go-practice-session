package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverseStr(s string, k int) string {

	reversePart := []rune(s)
	n := len(s)

	for i := 0; i < n; i++ {

		j := min(i+k-1, n-1)
		fmt.Println(j)
		start, end := i, j

		for start < end {
			start++
			end--
			reversePart[start], reversePart[end] = reversePart[end], reversePart[start]
		}
		i += 2 * k

	}

	return string(reversePart)
}

func main() {
	s := "abcdefg" // "bacdfeg"
	k := 2
	// s := "abcd" // "bacd"
	// k := 2
	// s := "abcdef" // "cbadef"
	// k := 3
	result := reverseStr(s, k)
	fmt.Println(result)
}
