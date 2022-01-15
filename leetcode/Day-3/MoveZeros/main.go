package main

import "fmt"

func moveZeroes(nums []int) []int {
	left := 0
	right := len(nums)
	pos := 0

	for left < right {
		if nums[left] != 0 {
			temp := nums[pos]
			nums[pos] = nums[left]
			nums[left] = temp
			pos++
			left++
		} else {
			left++
		}
	}

	// for left < right {
	// 	temp := nums[left]
	// 	nums[left] = nums[right]
	// 	nums[right] = temp
	// 	left++
	// 	right--
	// }

	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	result := moveZeroes(nums)
	fmt.Println(result)
}
