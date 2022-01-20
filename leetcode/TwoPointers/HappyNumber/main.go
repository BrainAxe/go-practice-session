package main

import "fmt"

func isHappy(n int) bool {
	slow := n
	fast := squireSum(n)

	for slow != fast {
		slow = squireSum(slow)
		fast = squireSum(fast)
		fast = squireSum(fast)
	}

	if slow == 1 {
		return true
	} else {
		return false
	}

}

func squireSum(n int) int {
	sum := 0
	for n > 0 {
		m := n % 10
		sum = sum + m*m
		n = n / 10
	}
	return sum
}

func main() {
	result := isHappy(2)
	fmt.Println(result)
}
