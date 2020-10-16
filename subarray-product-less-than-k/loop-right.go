package main

import "fmt"

func main() {
	fmt.Println(subArray([]int{10, 5, 2, 6}, 100))                                 // 8
	fmt.Println(subArray([]int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, 19)) // 18
	fmt.Println(subArray([]int{1, 2, 3}, 0))                                       // 0
}

func subArray(ints []int, target int) int {
	intsCounts := len(ints)
	if intsCounts <= 1 {
		return 1
	}

	counts := 0

	i := 0
	// j := 0
	prod := 1

	for j := 0; j < intsCounts; j++ {
		prod *= ints[j]

		for prod >= target && i < intsCounts {
			prod /= ints[i]
			i++
		}
		if j >= i {
			counts += j - i + 1
		}
	}

	return counts
}
