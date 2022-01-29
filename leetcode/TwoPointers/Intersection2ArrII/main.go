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
	m := make(map[int]int)

	for _, num := range nums1 {
		_, ok := m[num]
		if !ok {
			m[num] = 1
		} else {
			m[num] += 1
		}
	}

	for _, num := range nums2 {
		count, ok := m[num]
		if ok && count > 0 {
			res = append(res, num)
			m[num] -= 1
		}
	}

	return res
}

func main() {
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}
	// nums1 := []int{1, 2, 2, 1}
	// nums2 := []int{2, 2}
	nums1 := []int{1, 2}
	nums2 := []int{1, 1}
	result := intersection(nums1, nums2)
	fmt.Println(result)
}
