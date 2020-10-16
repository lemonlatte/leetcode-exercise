package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}

func threeSum(ints []int) [][]int {
	results := [][]int{}
	visited := map[[3]int]bool{}
	sort.Slice(ints, func(i, j int) bool { return ints[i] < ints[j] })
	for i := range ints {
		target := ints[i]
		if target > 0 {
			break
		}
		j := i + 1
		k := len(ints) - 1
		for j < k {
			v0 := ints[j]
			v1 := ints[k]
			if visited[[3]int{target, v0, v1}] {
				// There might be multiple identical number in the middle
				// This would cause a number visited.
				// Add 1 to check the next number
				j += 1
				k -= 1
				continue
			}
			sum := target + v0 + v1
			if sum < 0 {
				j += 1
			} else if sum > 0 {
				k -= 1
			} else {
				pair := [3]int{target, v0, v1}
				visited[pair] = true
				results = append(results, pair[:])
				// There might be numbers in between the result.
				// Add 1 to continue the loop
				j += 1
				// Do both j++ and k--. Since (j++, k) would be either identical or
				// greater than (j, k) , we do no need to do one more addition comparison
				k -= 1
			}
		}
	}
	return results
}
