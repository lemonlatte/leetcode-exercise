package main

import (
	"fmt"
)

func main() {
	fmt.Println(squareOfArray([]int{-4, -1, 0, 3, 10}))
	fmt.Println(squareOfArray([]int{-7, -3, 2, 3, 11}))
}

func squareOfArray(ints []int) []int {
	result := make([]int, len(ints))
	i := 0
	j := len(ints) - 1

	counter := j
	for {
		lowerSquare := ints[i] * ints[i]
		higherSquare := ints[j] * ints[j]

		if i == j {
			result[counter] = lowerSquare
			break
		}

		if lowerSquare >= higherSquare {
			i += 1
			result[counter] = lowerSquare
		} else {
			j -= 1
			result[counter] = higherSquare
		}
		counter -= 1
	}

	return result
}
