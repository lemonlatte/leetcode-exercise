package main

import "fmt"

func main() {
	fmt.Println(subArray([]int{10, 5, 2, 6}, 100))
}

func subArray(ints []int, target int) int {
	counts := 0
	intsCounts := len(ints)
	for i := 0; i < intsCounts; i++ {
		prod := ints[i]
		if prod < target {
			counts += 1
		} else {
			continue
		}
		j := i + 1
		for j < intsCounts && prod < target {
			prod *= ints[j]
			if prod < target {
				j++
			} else {
				break
			}
		}
		counts += (j - i - 1)
	}

	return counts
}
