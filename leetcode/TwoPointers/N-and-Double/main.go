package main

import "fmt"

func checkIfExist(arr []int) bool {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] == 2*arr[i] && i != j {
				return true
			}
		}
	}
	return false
}

func main() {
	// nums := []int{0, 0}
	// nums := []int{-2, 0, 10, -19, 4, 6, -8}
	nums := []int{3, 1, 7, 11}
	result := checkIfExist(nums)
	fmt.Println(result)
}
