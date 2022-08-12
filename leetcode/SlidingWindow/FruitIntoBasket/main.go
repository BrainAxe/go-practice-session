package main

import (
	"fmt"
)

func totalFruit(fruits []int) int {

	if len(fruits) == 0 {
		return 0
	}
	if len(fruits) == 1 {
		return 1
	}

	i, j, ans := 0, 0, 0
	hashMap := make(map[int]int)

	for j < len(fruits) {
		_, found := hashMap[fruits[j]]
		if found {
			hashMap[fruits[j]]++
		} else {
			hashMap[fruits[j]] = 1
		}

		if len(hashMap) <= 2 {
			if ans < j-i+1 {
				ans = j - i + 1
			}
		} else if len(hashMap) > 2 {
			for len(hashMap) > 2 {
				hashMap[fruits[i]]--
				if hashMap[fruits[i]] == 0 {
					delete(hashMap, fruits[i])
				}
				i++
			}
		}
		j++
	}

	return ans

}

func main() {
	// arr := []int{1, 2, 1}
	// arr := []int{0, 1, 2, 2}
	arr := []int{1, 2, 3, 2, 2}
	result := totalFruit(arr)
	fmt.Println(result)
}
