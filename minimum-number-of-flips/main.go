package main

import "fmt"

func main() {
	// m := Matrix{[][]int{{0, 0}, {0, 1}}, 2, 2}
	// m.Flip(1, 0)
	// fmt.Println(m.Elements)
	// fmt.Println(m.AllZero())
	fmt.Println(MyFlip([][]int{{0, 0}, {0, 1}}))
	fmt.Println(MyFlip([][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 0}}))
	fmt.Println(MyFlip([][]int{{1, 0, 0}, {1, 0, 0}}))
	fmt.Println(MyFlip([][]int{{0}}))
}

type Matrix struct {
	Elements [3][3]int
	MaxX     int
	MaxY     int
}

func (m *Matrix) Flip(x, y int) {

	if x-1 >= 0 {
		m.Elements[x-1][y] = (m.Elements[x-1][y] + 1) & 1
	}
	if y-1 >= 0 {
		m.Elements[x][y-1] = (m.Elements[x][y-1] + 1) & 1
	}
	if x+1 < m.MaxX {
		m.Elements[x+1][y] = (m.Elements[x+1][y] + 1) & 1
	}
	if y+1 < m.MaxY {
		m.Elements[x][y+1] = (m.Elements[x][y+1] + 1) & 1
	}

	m.Elements[x][y] = (m.Elements[x][y] + 1) & 1
}

func (m *Matrix) AllZero() bool {
	for i := 0; i < m.MaxX; i++ {
		for j := 0; j < m.MaxY; j++ {
			if m.Elements[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func NewMatrix(matrix [][]int) Matrix {
	maxX := len(matrix)
	maxY := len(matrix[0])

	elements := [3][3]int{}
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			elements[i][j] = matrix[i][j]
		}
	}

	return Matrix{
		Elements: elements,
		MaxX:     maxX,
		MaxY:     maxY,
	}
}

func MyFlip(mat [][]int) int {
	maxX := len(mat)
	maxY := len(mat[0])

	visited := map[Matrix]bool{}

	m := NewMatrix(mat)

	if m.AllZero() {
		return 0
	}

	queue := []Matrix{m}
	counter := 0

	for len(queue) > 0 {
		currentQueueSize := len(queue)
		counter += 1
		for l := 0; l < currentQueueSize; l++ {
			for i := 0; i < maxX; i++ {
				for j := 0; j < maxY; j++ {
					m := queue[l] // copy by value
					m.Flip(i, j)
					if visited[m] {
						continue
					}
					visited[m] = true

					if m.AllZero() {
						return counter
					} else {
						queue = append(queue, m)
					}
				}
			}
		}

		queue = queue[currentQueueSize:]
	}

	return -1
}
