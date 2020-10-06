package main

import (
	"fmt"
	"math"
)

func numSquares(target int) int {
	squareNumbers := []int{}
	for i := 1; float64(i) <= math.Sqrt(float64(target)); i++ {
		square := i * i
		if square == target {
			return 1
		}
		squareNumbers = append(squareNumbers, square)
	}

	step := 1
	visitedNum := map[int]struct{}{}
	queue := make([]int, len(squareNumbers))
	copy(queue, squareNumbers)
	for {
		step += 1
		queueSize := len(queue)
		for _, i := range queue {
			for _, n := range squareNumbers {
				sum := i + n
				if sum == target {
					return step
				} else {
					if _, ok := visitedNum[sum]; !ok {
						visitedNum[sum] = struct{}{}
						queue = append(queue, sum)
					}
				}
			}
		}
		queue = queue[queueSize:]
	}
}

func main() {
	fmt.Println(numSquares(16))
}
