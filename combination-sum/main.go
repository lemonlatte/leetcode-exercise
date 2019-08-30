package main

import (
	"fmt"
)

func findNumber(numSet []int, target int) [][]int {
	results := [][]int{}
	for i, num := range numSet {
		if num == target {
			results = append(results, []int{num})
		} else if num < target {
			subSet := findNumber(numSet[i:], target-num)
			for _, set := range subSet {
				results = append(results, append([]int{num}, set...))
			}
		}

	}

	return results
}

func combinationSum(candidates []int, target int) [][]int {
	return findNumber(candidates, target)
}
