package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	start := 0
	end := k
	strNum := strconv.Itoa(num) // Int to String
	count := 0

	for end <= len(strNum) {
		sliced_num, _ := strconv.Atoi(strNum[start:end]) // String to Int
		if sliced_num == 0 {
			start++
			end++
			continue
		}

		if num%sliced_num == 0 {
			count++
		}

		start++
		end++
	}

	return count
}

func main() {
	num := 430043
	k := 2

	result := divisorSubstrings(num, k)
	fmt.Println(result)
}
