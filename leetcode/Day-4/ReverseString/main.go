package main

import "fmt"

func reverseString(s []byte) {

	fmt.Println(string(s))
	start := 0
	end := len(s) - 1

	for start < end {
		temp := s[start]
		s[start] = s[end]
		s[end] = temp
		start++
		end--
	}
	fmt.Println(string(s))
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
}
