package main

import "fmt"

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	contains := 0
	for start < end {
		temp := 0
		if height[start] < height[end] {
			temp = height[start] * (end - start)
			fmt.Println("start ", temp)
			start++
		} else {
			temp = height[end] * (end - start)
			fmt.Println("end ", temp)
			end--
		}
		if temp > contains {
			contains = temp
			fmt.Println("Contains ", contains)
		}
	}
	return contains
}

func main() {
	// nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	nums := []int{0, 2}
	result := maxArea(nums)
	fmt.Println(result)
}
