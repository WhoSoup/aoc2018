package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type point struct {
	x, y, size int
	infinite   bool
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (p point) distance(x int, y int) int {
	return abs(p.x-x) + abs(p.y-y)
}

type area struct {
	max, may int
	points   []point
}

func (a *area) add(x int, y int) {
	if x > a.max {
		a.max = x
	}
	if y > a.may {
		a.may = y
	}
	a.points = append(a.points, point{x, y, 0, false})
}

func (a *area) calc(x int, y int) bool {
	closest := a.points[0].distance(x, y)
	cpts := []point{a.points[0]}
	index := 0
	safety := closest

	for i := 1; i < len(a.points); i++ {
		d := a.points[i].distance(x, y)
		safety += d
		if d == closest {
			cpts = append(cpts, a.points[i])
		}
		if d < closest {
			closest = d
			cpts = []point{a.points[i]}
			index = i
		}
	}

	if len(cpts) == 1 {
		a.points[index].size++
		if x == 0 || y == 0 || x == a.max || y == a.may {
			a.points[index].infinite = true
		}
	}

	return safety < SAFETY
}

// SAFETY poop
const SAFETY = 10000

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var x, y int

	var universe area

	for {
		count, err := fmt.Fscanf(file, "%d, %d\n", &x, &y)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if count != 2 {
			continue
		}

		universe.add(x, y)
	}

	safe := 0
	for i := 0; i <= universe.may; i++ {
		for j := 0; j <= universe.max; j++ {
			if isSafe := universe.calc(j, i); isSafe {
				safe++
			}
		}
	}

	largest := universe.points[0].size
	for _, p := range universe.points {
		if !p.infinite && p.size > largest {
			largest = p.size
		}
	}

	fmt.Println("Largest:", largest)
	fmt.Println("Safesize:", safe)
}
