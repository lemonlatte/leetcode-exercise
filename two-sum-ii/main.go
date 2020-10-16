package main

import "fmt"

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3, 4, 32, 54}, 7))
}

func TwoSum(ints []int, target int) []int {
	i := 0
	j := len(ints) - 1

	for {
		sum := ints[i] + ints[j]
		if sum > target {
			j -= 1
		} else if sum < target {
			i += 1
		} else {
			return []int{i + 1, j + 1}
		}
	}
}
