package main

import (
	"fmt"
)

func longestOnes(nums []int, k int) int {
	i, j := 0, 0

	for j < len(nums) {

		if nums[j] == 0 {
			k--
		}

		if k < 0 {
			if nums[i] == 0 {
				k++
			}
			i++
		}

		j++
	}
	return j - i
}

func main() {
	// nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	// k := 2
	nums := []int{0, 0, 0, 0}
	k := 0
	result := longestOnes(nums, k)
	fmt.Println(result)
}
