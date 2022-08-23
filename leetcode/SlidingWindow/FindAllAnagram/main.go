package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {

	i, j, cnt := 0, 0, 0
	ans := []int{}
	hashMap := make(map[string]int)

	for r := 0; r < len(p); r++ {
		_, found := hashMap[string(p[r])]
		if found {
			hashMap[string(p[r])]++
		} else {
			hashMap[string(p[r])] = 1
			cnt++
		}
	}

	for j < len(s) {
		_, found := hashMap[string(s[j])]
		if found {
			hashMap[string(s[j])]--
			if hashMap[string(s[j])] == 0 {
				cnt--
			}
		}

		if j-i+1 == len(p) {
			if cnt == 0 {
				ans = append(ans, i)
			}

			_, found := hashMap[string(s[i])]
			if found {
				hashMap[string(s[i])]++
				if hashMap[string(s[i])] == 1 {
					cnt++
				}
			}
			i++

		}
		j++
	}

	return ans

}

func main() {
	result := findAnagrams("cbaebabacd", "abc")
	fmt.Println(result)
}
