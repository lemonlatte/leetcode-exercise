package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) Adjacent() []Point {
	points := []Point{
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
	}

	return points
}

func countIsland() int {

	world := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0},
	}

	maxX := len(world)
	if maxX == 0 {
		return 0
	}
	maxY := len(world[0])

	visitedArray := map[int]struct{}{}

	islandCount := 0

	for i := 0; i < len(world); i++ {
		for j := 0; j < len(world[i]); j++ {
			visitID := i<<8 | j
			if _, ok := visitedArray[visitID]; ok {
				continue
			}

			if world[i][j] == 0 {
				visitedArray[visitID] = struct{}{}
				continue
			}

			queue := []Point{{X: i, Y: j}}
			for len(queue) > 0 {
				stageLen := len(queue)
				for k := 0; k < stageLen; k++ {
					p := queue[k]

					for _, p := range p.Adjacent() {
						x := p.X
						y := p.Y
						if x < 0 || y < 0 || x >= maxX || y >= maxY {
							continue
						}

						visitID := x<<8 | y
						// need to break breadth here
						if _, ok := visitedArray[visitID]; ok {
							continue
						}
						visitedArray[visitID] = struct{}{}

						if world[x][y] == 1 {
							queue = append(queue, p)
						}
					}
				}
				queue = queue[stageLen:]
			}
			islandCount += 1
		}
	}
	return islandCount
}

func main() {
	fmt.Println(countIsland())
}
