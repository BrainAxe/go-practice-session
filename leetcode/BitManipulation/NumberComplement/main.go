package main

import (
	"fmt"
	"math/bits"
)

func findComplement(num int) int {
	shift := bits.Len(uint(num))
	return num ^ (1<<shift - 1)
}

func main() {
	num := 5
	result := findComplement(num)
	fmt.Println(result)
}
