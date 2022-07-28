package main

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := 100000
	i := k - 1
	for i < len(nums) {
		res := nums[i] - nums[i-k+1]

		if res < ans {
			ans = res
		}
		i++
	}
	return ans
}

func main() {
	nums := []int{9, 4, 1, 7}
	k := 2

	result := minimumDifference(nums, k)
	fmt.Println(result)
}
