package main

import "fmt"

type Point struct {
	A int
	B int
	C int
	D int
}

func (p Point) ID() int {
	return p.A<<24 | p.B<<16 | p.C<<8 | p.D
}

func (p1 Point) Equal(p2 Point) bool {
	return p1.A == p2.A && p1.B == p2.B && p1.C == p2.C && p1.D == p2.D
}

func (p Point) Adjacent() []Point {
	return []Point{
		{(p.A + 1) % 10, p.B, p.C, p.D},
		{p.A, (p.B + 1) % 10, p.C, p.D},
		{p.A, p.B, (p.C + 1) % 10, p.D},
		{p.A, p.B, p.C, (p.D + 1) % 10},
		{(p.A + 9) % 10, p.B, p.C, p.D},
		{p.A, (p.B + 9) % 10, p.C, p.D},
		{p.A, p.B, (p.C + 9) % 10, p.D},
		{p.A, p.B, p.C, (p.D + 9) % 10},
	}
}

func NewPoint(pointStr string) Point {
	b := []byte(pointStr)
	if len(b) < 4 {
		panic("invaid point")
	}

	return Point{int(b[0]) - 48, int(b[1]) - 48, int(b[2]) - 48, int(b[3]) - 48}
}

func openLock(deadends []string, target string) int {
	targetPoint := NewPoint(target)

	if targetPoint.Equal(Point{0, 0, 0, 0}) {
		return 0
	}

	fmt.Println(targetPoint)
	minStep := 9 * 9 * 9 * 9

	visitedPoint := map[int]struct{}{}
	for _, p := range deadends {
		if p == "0000" {
			return -1
		}
		pp := NewPoint(p)
		visitedPoint[pp.ID()] = struct{}{}
	}
	// add locked point as initial visited point

	currentStep := 0
	queue := []Point{{0, 0, 0, 0}}

	for len(queue) > 0 {
		stageLen := len(queue)
		currentStep += 1
		for i := 0; i < stageLen; i++ {
			for _, p := range queue[i].Adjacent() {

				if _, ok := visitedPoint[p.ID()]; ok {
					continue
				}

				visitedPoint[p.ID()] = struct{}{}

				if p.Equal(targetPoint) {
					if currentStep < minStep {
						minStep = currentStep
					}
				} else {
					queue = append(queue, p)
				}
			}
		}

		queue = queue[stageLen:]
	}
	if minStep == 9*9*9*9 {
		return -1
	}
	return minStep
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
