package main

import "fmt"

func search(nums []int, target int) int {

	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := ((high + low) / 2)

		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	nums := []int{2, 5, 8, 12, 65, 76, 143, 222, 999}
	target := 8
	result := search(nums, target)
	fmt.Println(result)
}
