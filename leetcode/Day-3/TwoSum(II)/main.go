package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
	var arr []int = make([]int, 2)

	for start < end {
		if numbers[start]+numbers[end] < target {
			start++
		} else if numbers[start]+numbers[end] > target {
			end--
		} else {
			arr[0] = start + 1
			arr[1] = end + 1
			break
		}
	}
	return arr
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
