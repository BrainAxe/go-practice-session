package main

import "fmt"

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	m := make(map[int]bool, len(nums1))

	for _, num := range nums1 {
		m[num] = true
	}

	for _, num := range nums2 {
		if m[num] == true && !contains(res, num) {
			res = append(res, num)
		}
	}

	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	result := intersection(nums1, nums2)
	fmt.Println(result)
}
