//238. Product of Array Except Self

package main

import "fmt"

func largestAltitude(gain []int) int {
	result := 0
	curr_gain := 0

	for i := 0; i < len(gain); i++ {
		curr_gain += gain[i]
		if curr_gain > result {
			result = curr_gain
		}
	}

	return result
}
func main() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	result := largestAltitude(gain)
	fmt.Println(result)
}
