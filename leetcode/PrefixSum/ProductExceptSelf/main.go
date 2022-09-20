//238. Product of Array Except Self

package main

import "fmt"

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	k := 1
	l := 1
	for i := 0; i < len(nums); i++ {
		k = k * nums[i]
		ans[i] = k
	}

	for j := len(nums) - 1; j > 0; j-- {
		ans[j] = ans[j-1] * l
		l = l * nums[j]
	}
	ans[0] = l
	return ans
}
func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result)
}
