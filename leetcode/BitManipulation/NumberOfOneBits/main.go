//238. Product of Array Except Self

package main

import "fmt"

func hammingWeight(num uint32) int {
	cnt := 0

	for num > 0 {
		num &= (num - 1)
		cnt++
	}
	return cnt
}

func main() {
	var num uint32 = 00000000000000000000000000001011
	result := hammingWeight(num)
	fmt.Println(result)
}
