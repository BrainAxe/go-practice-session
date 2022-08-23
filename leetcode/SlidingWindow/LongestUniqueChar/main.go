package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	i, j, ans := 0, 0, 0
	hashMap := make(map[string]int)

	for j < len(s) {
		_, found := hashMap[string(s[j])]
		if found {
			hashMap[string(s[j])]++
		} else {
			hashMap[string(s[j])] = 1
		}

		if len(hashMap) == j-i+1 {
			if j-i+1 > ans {
				ans = j - i + 1
			}
		} else if len(hashMap) < j-i+1 {
			for len(hashMap) < j-i+1 {
				hashMap[string(s[i])]--
				if hashMap[string(s[i])] == 0 {
					delete(hashMap, string(s[i]))
				}
				i++
			}
		}
		j++
	}

	return ans

}

func main() {
	result := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(result)
}
