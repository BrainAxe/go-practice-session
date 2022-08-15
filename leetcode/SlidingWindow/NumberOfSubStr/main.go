package main

import (
	"fmt"
)

func numberOfSubstrings(s string) int {

	i, j, ans := 0, 0, 0
	hashMap := make(map[string]int)

	for j < len(s) {
		// _, found := hashMap[string(s[j])]
		// if found {
		// 	hashMap[string(s[j])]++
		// } else {
		// 	hashMap[string(s[j])] = 1
		// }

		hashMap[string(s[j])]++

		for hashMap["a"] >= 1 && hashMap["b"] >= 1 && hashMap["c"] >= 1 {
			ans += len(s) - j
			hashMap[string(s[i])]--
			i++
		}

		j++
	}

	// fmt.Println(hashMap)

	return ans

}

func main() {
	result := numberOfSubstrings("abcabc")
	fmt.Println(result)
}
