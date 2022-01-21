package main

import "fmt"

func sortColors(nums []int) {
	zeroPos := 0
	start := 0
	end := len(nums) - 1

	for start <= end {
		if nums[start] == 0 {
			nums[start], nums[zeroPos] = nums[zeroPos], nums[start]
			zeroPos++
			start++
		} else if nums[start] == 2 {
			nums[start], nums[end] = nums[end], nums[start]
			end--
		} else {
			start++
		}
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
}
