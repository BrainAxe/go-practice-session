package main

import "fmt"

func removeElement(nums []int, val int) int {
	start, end := 0, len(nums)-1

	for start <= end && end < len(nums) {
		if nums[start] == val {
			for end > start && end < len(nums) && nums[end] == val {
				end--
			}
			nums[start], nums[end] = nums[end], nums[start]
			end--
		}
		start++
	}

	fmt.Println(nums)

	return end + 1
}

func main() {
	nums := []int{3, 3}
	// nums := []int{1}
	// nums := []int{3, 2, 2, 3}
	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 3
	result := removeElement(nums, target)
	fmt.Println(result)
}
