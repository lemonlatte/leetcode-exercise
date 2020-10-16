package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumSmaller([]int{-2, 0, 1, 3}, 2))                                                                                        // 2
	fmt.Println(threeSumSmaller([]int{-2, 0, -1, 3}, 2))                                                                                       // 3
	fmt.Println(threeSumSmaller([]int{0, -2, -2, -4, 4, 3, 1, -2, -5, 1, 0, -5, -4, 4, 0, -4}, -4))                                            // 211
	fmt.Println(threeSumSmaller([]int{-3, 5, 6, -3, -8, -1, -3, 7, -5, -3, 4, 4, -7, -7, 4, -8, 5, -10, -2, 7, 2, 8, -4, -3, 3, 5, 2, -6}, 8)) // 2726
}

func threeSumSmaller(ints []int, target int) int {
	results := 0
	intsCount := len(ints)

	sort.Slice(ints, func(i, j int) bool { return ints[i] < ints[j] })
	for i := range ints {
		v0 := ints[i]
		if v0 >= target && v0 >= 0 { // If target is negative, we will miss pair with `v0 >= target`.
			break
		}

		for j := i + 1; j < intsCount-1; j++ {
			v1 := ints[j]
			for k := intsCount - 1; k > j; k-- {
				if v0+v1+ints[k] < target {
					results += k - j
					break
				}
			}
		}
	}

	return results
}
