package main

import "fmt"

func singleNumber(nums []int) []int {
	var xorBit int
	results := []int{0, 0}
	for _, i := range nums {
		xorBit ^= i
	}

	firstDiffBit := xorBit & ^(xorBit - 1)

	for _, i := range nums {
		if (i & firstDiffBit) > 0 {
			results[0] ^= i
		} else {
			results[1] ^= i
		}
	}
	return results
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 4, 2, 5}))
	fmt.Println(singleNumber([]int{1, 2, 1, 4, 2, 1, 2, 1, 2, 7}))
}
