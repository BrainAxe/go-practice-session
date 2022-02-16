package main

import "fmt"

func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}
	cnt := 0
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i+1] != s[i+2] && s[i+2] != s[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	result := countGoodSubstrings("xyzzaz")
	fmt.Println(result)
}
