package main

import (
	"fmt"
	"math"
)

func minimumCardPickup(cards []int) int {
	ans := math.MaxInt32
	hashMap := make(map[int]int)

	for j, v := range cards {
		_, found := hashMap[v]
		if found {
			if ans > (j - hashMap[v] + 1) {
				ans = j - hashMap[v] + 1
			}
			hashMap[v] = j
		} else {
			hashMap[v] = j
		}

	}
	if ans == math.MaxInt32 {
		ans = -1
	}

	return ans
}

func main() {
	arr := []int{3, 4, 2, 3, 4, 7}
	result := minimumCardPickup(arr)
	fmt.Println(result)
}
