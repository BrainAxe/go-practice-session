package main

import (
	"fmt"
)

func numberOfSubarrays(nums []int, k int) int {

	distinct := func(k int) int {

		i, j, odd_cnt, ans := 0, 0, 0, 0

		for j < len(nums) {
			if nums[j]%2 != 0 {
				odd_cnt++
			}

			if odd_cnt > k {
				for odd_cnt > k {
					if nums[i]%2 != 0 {
						odd_cnt--
					}
					i++
				}
			}

			ans += (j - i)
			j++
		}

		return ans
	}

	return distinct(k) - distinct(k-1)

}

func main() {
	arr := []int{1, 1, 2, 1, 1}
	k := 3

	result := numberOfSubarrays(arr, k)
	fmt.Println(result)
}
