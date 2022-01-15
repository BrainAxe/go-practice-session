package main

import (
	"fmt"
	// "sort"
)

func abs(val int) int {
	fmt.Println("Entered", val)
	if val < 0 {
		fmt.Println("IF", -val)
		return -val
	} else {
		fmt.Println(" ELSE ", val)
		return val
	}
}

func sortedSquares(nums []int) []int {
	// for i, num := range nums {
	// 	nums[i] = num * num
	// }
	// sort.Ints(nums)
	// return nums
	high := len(nums) - 1
	low := 0
	i := high
	var arr []int = make([]int, len(nums))

	for low <= high {
		if abs(nums[low]) > abs(nums[high]) {
			fmt.Println("entered")
			arr[i] = nums[low] * nums[low]
			low = low + 1
			fmt.Println(arr)
		} else {
			arr[i] = nums[high] * nums[high]
			high = high - 1
		}
		i = i - 1

	}
	return arr
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	result := sortedSquares(nums)
	fmt.Println(result)
}
