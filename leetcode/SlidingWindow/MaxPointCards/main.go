package main

import (
	"fmt"
	"math"
)

func maxScore(cardPoints []int, k int) int {
	i, j, total_sum, curr_sum := 0, 0, 0, 0
	mini := math.MaxInt32
	rem := len(cardPoints) - k

	for j < len(cardPoints) {
		total_sum += cardPoints[j]
		curr_sum += cardPoints[j]

		if j-i+1 == rem {
			if mini > curr_sum {
				mini = curr_sum
			}
			curr_sum -= cardPoints[i]
			i++
		}
		j++

	}

	if mini == math.MaxInt32 {
		return total_sum
	} else {
		return total_sum - mini
	}

}

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6, 1}
	// arr := []int{1, 1000, 1}
	arr := []int{11, 49, 100, 20, 86, 29, 72}
	result := maxScore(arr, 4)
	fmt.Println(result)
}
