package main

import (
	"fmt"
	"io"
	"os"
)

type point struct {
	x, y, vx, vy int
}

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func move(pts []point) (int, int, int, int, int) {
	maxX := minInt
	maxY := minInt
	minX := maxInt
	minY := maxInt

	for i := range pts {
		pts[i].x += pts[i].vx
		if pts[i].x > maxX {
			maxX = pts[i].x
		} else if pts[i].x < minX {
			minX = pts[i].x
		}
		pts[i].y += pts[i].vy
		if pts[i].y > maxY {
			maxY = pts[i].y
		} else if pts[i].y < minY {
			minY = pts[i].y
		}
	}

	return (maxX - minX) * (maxY - minY), maxX - minX, maxY - minY, minX, minY
}

func cp(old []point) []point {
	new := make([]point, len(old))
	for i, e := range old {
		new[i] = point{e.x, e.y, e.vx, e.vy}
	}
	return new
}

func print(pts []point, width int, height int, offX int, offY int) {
	f := make([][]byte, height+1)

	for i := 0; i < height+1; i++ {
		f[i] = make([]byte, width+1)
	}

	for _, p := range pts {
		f[p.y-offY][p.x-offX] = '#'
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

	old := cp(points)
	area, width, height, offX, offY := move(points)
	moves := 1

	for {
		old = cp(points)
		a, w, h, ox, oy := move(points)
		if a < area {
			area = a
			width = w
			height = h
			offX = ox
			offY = oy
		} else if a > area {
			print(old, width, height, offX, offY)
			fmt.Println("Seconds:", moves)
			break
		}
		moves++
	}
}
