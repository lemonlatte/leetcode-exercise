package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) Adjacent() []Point {
	x := p.X
	y := p.Y
	return []Point{
		{x + 1, y},
		{x, y + 1},
		{x - 1, y},
		{x, y - 1},
	}
}

func main() {

	var maze = [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}

	maxX := len(maze)
	// Assume an array has fixed size.
	// If not, try to get the max size by looping `maze`
	maxY := len(maze[0])
	fmt.Println(maxX, maxY)
	var mazeResult [][]int

	for i := 0; i < maxX; i++ {
		mazeResult = append(mazeResult, make([]int, maxY))
		for j := 0; j < maxY; j++ {
			step := 0
			found := false
			if maze[i][j] == 0 || maze[i][j] == -1 {
				mazeResult[i][j] = maze[i][j]
				continue
			}

			queue := []Point{{
				X: i,
				Y: j,
			}}

			visitedPoints := map[int]struct{}{}
		SEARCH_GATE:
			for len(queue) > 0 {
				qSize := len(queue)
				step += 1
				for n := 0; n < qSize; n++ {
					for _, p := range queue[n].Adjacent() {
						if p.X < 0 || p.Y < 0 || p.X >= maxX || p.Y >= maxY {
							continue
						}

						visitPoint := p.X<<8 | p.Y
						if _, ok := visitedPoints[visitPoint]; ok {
							fmt.Println(visitPoint)
							continue
						}

						visitedPoints[visitPoint] = struct{}{}

						switch maze[p.X][p.Y] {
						case 0:
							found = true
							break SEARCH_GATE
						case -1:
							continue
						default: // INF 2147483647
							queue = append(queue, p)
						}
					}
				}
				queue = queue[qSize:]
			}
			// check if no way (INF)
			if found {
				mazeResult[i][j] = step
			} else {
				mazeResult[i][j] = -1
			}
		}
	}

	fmt.Println(mazeResult)
}
