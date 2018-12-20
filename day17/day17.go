package main

import (
	"bufio"
	"fmt"
	"os"
)

const SAND, CLAY, STILL, FLOWING = 0, 1, 2, 3

type field map[string]int

func (f field) c(a, b int) int {
	return f[fmt.Sprintf("%d,%d", a, b)]
}

func (f field) s(a, b, c int) {
	f[fmt.Sprintf("%d,%d", a, b)] = c
}

func (f field) reachable(a, b int) bool {
	return f.c(a, b) == SAND || f.c(a, b) == FLOWING
}

func isThereAWall(field field, x, y, direction int) (bool, int, int) {
	if !field.reachable(x, y) {
		return true, x, y
	}
	field.s(x, y, FLOWING)
	if field.reachable(x, y+1) {
		return false, x, y + 1
	}
	return isThereAWall(field, x+direction, y, direction)
}

func down(field field, x, y, max int) {
	if y > max {
		return
	}

	field.s(x, y, FLOWING)

	if field.c(x, y+1) == FLOWING { // hitting another branch
		return
	}

	if field.c(x, y+1) > SAND {
		leftWall, leftX, leftY := isThereAWall(field, x-1, y, -1)
		rightWall, rightX, rightY := isThereAWall(field, x+1, y, 1)
		if leftWall && rightWall {
			for i := leftX + 1; i < rightX; i++ {
				field.s(i, y, STILL)
			}
			down(field, x, y-1, max)
		}
		if !leftWall {
			down(field, leftX, leftY, max)
		}
		if !rightWall {
			down(field, rightX, rightY, max)
		}
	} else {
		down(field, x, y+1, max)
	}
}

func main() {
	file, _ := os.Open("input.txt")

	field := make(field)

	scan := bufio.NewScanner(file)
	var a, b byte
	var i, j, k int

	minX, maxX, minY, maxY := 5000, 0, 5000, 0
	for scan.Scan() {
		fmt.Sscanf(scan.Text(), "%c=%d, %c=%d..%d", &a, &i, &b, &j, &k)
		for y := j; y <= k; y++ {
			if a == 'x' {
				if i < minX {
					minX = i
				}
				if i > maxX {
					maxX = i
				}
				if j < minY {
					minY = j
				}
				if k > maxY {
					maxY = j
				}
				field.s(i, y, 1)
			} else {
				if i < minY {
					minY = i
				}
				if i > maxY {
					maxY = i
				}
				if j < minX {
					minX = j
				}
				if k > maxX {
					maxX = j
				}
				field.s(y, i, 1)
			}
		}
	}

	down(field, 500, 1, maxY)
	//fmt.Println("Reached:", reached)
	reached, still := 0, 0

	for _, i := range field {
		if i == FLOWING || i == STILL {
			reached++
		}
		if i == STILL {
			still++
		}
	}

	fmt.Println("Part One:", reached)
	fmt.Println("Part Two:", still)

	/*	// write game to file
		var buf strings.Builder
		for y := minY; y <= maxY; y++ {
			for x := minX - 1; x <= maxX+1; x++ {
				if x == 500 && y == 0 {
					buf.WriteString("@")
				} else if field.c(x, y) == STILL {
					buf.WriteString("~")
				} else if field.c(x, y) == CLAY {
					buf.WriteString("#")
				} else if field.c(x, y) == FLOWING {
					buf.WriteString("|")
				} else {
					buf.WriteString(".")
				}
			}
			buf.WriteString("\n")
		}
		ioutil.WriteFile("testout.txt", []byte(buf.String()), 0x777)
	*/

}
