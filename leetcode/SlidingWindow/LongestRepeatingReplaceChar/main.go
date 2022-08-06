package main

import (
	"fmt"
)

func characterReplacement(s string, k int) int {
	l := 0
	maxFreq := 0
	ans := 0
	hashMap := make(map[string]int)

	for j, v := range s {
		_, found := hashMap[string(v)]
		if found {
			hashMap[string(v)]++
		} else {
			hashMap[string(v)] = 1
		}

		// fmt.Println(string(s[j]))

		if hashMap[string(s[j])] > maxFreq {
			maxFreq = hashMap[string(s[j])]
		}

		if (j-l+1)-maxFreq > k {
			hashMap[string(s[l])]--
			l++
		}

		if j-l+1 > ans {
			ans = j - l + 1
		}
	}
	fmt.Println(hashMap)
	return ans
}

func main() {
	s := "AABABBA"
	k := 1
	result := characterReplacement(s, k)
	fmt.Println(result)
}
