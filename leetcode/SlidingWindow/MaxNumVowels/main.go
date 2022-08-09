package main

import (
	"fmt"
	"strings"
)

func maxVowels(s string, k int) int {
	ans := 0
	i := 0
	j := 0
	vowels := 0

	for j < len(s) {
		if strings.Contains("aeiou", string(s[j])) == true {
			vowels++
		}

		if j-i+1 > k {
			for j-i+1 > k {
				if strings.Contains("aeiou", string(s[i])) == true {
					vowels--
				}
				i++
			}
		}

		if j-i+1 == k {
			if vowels > ans {
				ans = vowels
			}
		}

		j++

	}
	return ans
}

func main() {
	result := maxVowels("abciiidef", 3)
	fmt.Println(result)
}
