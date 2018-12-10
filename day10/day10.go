package main

import (
	"fmt"
	"io"
	"os"
)

type point struct {
	x, y, vx, vy int
}

type state struct {
	minX, maxX, minY, maxY int
}

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func (s state) area() int {
	return s.width() * s.height()
}
func (s state) width() int {
	return s.maxX - s.minX
}
func (s state) height() int {
	return s.maxY - s.minY
}
func (s state) offsetX() int {
	return s.minX
}
func (s state) offsetY() int {
	return s.minY
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func move(pts []point) state {
	s := state{maxInt, minInt, maxInt, minInt}

	for i := range pts {
		pts[i].x += pts[i].vx
		s.maxX = max(s.maxX, pts[i].x)
		s.minX = min(s.minX, pts[i].x)

		pts[i].y += pts[i].vy
		s.maxY = max(s.maxY, pts[i].y)
		s.minY = min(s.minY, pts[i].y)
	}

	return s
}

func print(pts []point, s state) {
	f := make([][]byte, s.height()+1)

	for i := 0; i < s.height()+1; i++ {
		f[i] = make([]byte, s.width()+1)
	}

	for _, p := range pts {
		// undo the last step too
		f[p.y-s.offsetY()-p.vy][p.x-s.offsetX()-p.vx] = '#'
	}

	for _, line := range f {
		for _, c := range line {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func main() {
	file, _ := os.Open("input.txt")

	points := []point{}

	for {
		var p point
		_, err := fmt.Fscanf(file, "position=<%d,%d> velocity=<%d,%d>\n", &p.x, &p.y, &p.vx, &p.vy)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		points = append(points, p)
	}

	state := move(points)
	moves := 1

	for {
		newstate := move(points)
		if newstate.area() > state.area() {
			print(points, state)
			fmt.Println("Seconds:", moves)
			break
		}
		state = newstate
		moves++
	}
}
