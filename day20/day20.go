package main

import (
	"fmt"
	"io/ioutil"
)

type coord struct {
	x, y int
}

func (c coord) s() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}
func main() {
	meow, _ := ioutil.ReadFile("input.txt")

	var stack []coord
	minDoors := make(map[string]int)
	var pos, npos coord
	minDoors[pos.s()] = 0

	for i := 1; i < len(meow)-1; i++ {
		npos = pos

		switch meow[i] {
		case '(':
			stack = append(stack, pos)
			continue
		case ')':
			pos = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			continue
		case '|':
			pos = stack[len(stack)-1]
			continue
		case 'N':
			npos.y++
		case 'S':
			npos.y--
		case 'E':
			npos.x++
		case 'W':
			npos.x--
		}

		if doors, ex := minDoors[npos.s()]; !ex || doors > minDoors[pos.s()]+1 {
			minDoors[npos.s()] = minDoors[pos.s()] + 1
		}

		pos = npos
	}

	longest, thou := 0, 0
	for _, c := range minDoors {
		if c > longest {
			longest = c
		}
		if c >= 1000 {
			thou++
		}
	}
	fmt.Println(longest)
	fmt.Println(thou)
}
