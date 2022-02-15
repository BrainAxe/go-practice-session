package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	pointer := 0
	sum := 0
	max_sum := 0

	for pointer < len(nums) {
		if pointer < k {
			sum += nums[pointer]
			max_sum = sum
			pointer++
		} else {
			sum += nums[pointer] - nums[pointer-k]

			if sum > max_sum {
				max_sum = sum
			}
			pointer++
		}
	}

	return float64(max_sum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	result := findMaxAverage(nums, k)
	fmt.Println(result)
}
