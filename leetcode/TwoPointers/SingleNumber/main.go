package main

import "fmt"

func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] = nums[0] ^ nums[i]
	}
	return nums[0]
}

func main() {
	nums := []int{2, 1, 2}
	result := singleNumber(nums)
	fmt.Println(result)
}
