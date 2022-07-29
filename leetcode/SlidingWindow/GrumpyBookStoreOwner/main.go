package main

import (
	"fmt"
	"math"
)

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	satisfied := 0
	not_satisfied := 0
	max_satisfied := 0

	for i := 0; i < len(grumpy); i++ {
		if grumpy[i] == 0 {
			satisfied += customers[i]
			customers[i] = 0
		}
		not_satisfied += customers[i]

		if i >= minutes {
			if grumpy[i-minutes] == 1 {
				not_satisfied -= customers[i-minutes]
			}
		}

		max_satisfied = int(math.Max(float64(max_satisfied), float64(not_satisfied)))

	}
	fmt.Println(max_satisfied, satisfied, not_satisfied)
	return (satisfied + max_satisfied)
}

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3

	result := maxSatisfied(customers, grumpy, minutes)
	fmt.Println(result)
}
