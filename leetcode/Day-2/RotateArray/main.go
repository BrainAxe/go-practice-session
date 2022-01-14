package main

import "fmt"

func rotate(nums []int, k int) {

	// high := len(nums) - 1

	// for k > 0 {
	// 	fmt.Println(k)
	// 	previous := nums[high]
	// 	for i := 0; i <= high; i++ {
	// 		temp := nums[i]
	// 		nums[i] = previous
	// 		previous = temp
	// 	}
	// 	fmt.Println(nums)
	// 	k--
	// }
	k %= len(nums) // k = k % len(nums)
	reverseArr := reverse(nums, 0, len(nums)-1)
	reverse1kItems := reverse(reverseArr, 0, k-1)
	result := reverse(reverse1kItems, k, len(nums)-1)
	fmt.Println(result)
}

func reverse(nums []int, start int, end int) []int {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
	return nums

}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	k := 3
	rotate(nums, k)
}
