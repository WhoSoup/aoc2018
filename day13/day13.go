package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type cart struct {
	car           byte
	x, y          int
	intersections int
	crashed       bool
}

func (c cart) coord() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

type byPos []cart

func (c byPos) Len() int      { return len(c) }
func (c byPos) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c byPos) Less(i, j int) bool {
	if c[i].x < c[j].x {
		return true
	} else if c[i].x == c[j].x {
		return c[i].y < c[j].y
	}
	return false
}

func removeCarts(paths []string) (carts []cart) {
	for y := 0; y < len(paths); y++ {
		for x := 0; x < len(paths[y]); x++ {
			switch paths[y][x] {
			case '<':
				fallthrough
			case '>':
				fallthrough
			case 'v':
				fallthrough
			case '^':
				c := cart{paths[y][x], x, y, 0, false}
				carts = append(carts, c)
			}
		}

		paths[y] = strings.Replace(paths[y], ">", "-", -1)
		paths[y] = strings.Replace(paths[y], "<", "-", -1)
		paths[y] = strings.Replace(paths[y], "^", "|", -1)
		paths[y] = strings.Replace(paths[y], "v", "|", -1)
	}
	return
}

func move(car cart, pos byte) cart {
	switch car.car {
	case '^':
		if pos == '+' {
			pos = []byte{'\\', '|', '/'}[car.intersections%3]
			car.intersections++
		}
		switch pos {
		case '/':
			car.car = '>'
			car.x++
		case '\\':
			car.car = '<'
			car.x--
		case '|':
			car.y--
		}

	case '>':
		if pos == '+' {
			pos = []byte{'/', '-', '\\'}[car.intersections%3]
			car.intersections++
		}
		switch pos {
		case '-':
			car.x++
		case '/':
			car.car = '^'
			car.y--
		case '\\':
			car.car = 'v'
			car.y++
		}
	case 'v':
		if pos == '+' {
			pos = []byte{'\\', '|', '/'}[car.intersections%3]
			car.intersections++
		}
		switch pos {
		case '|':
			car.y++
		case '/':
			car.car = '<'
			car.x--
		case '\\':
			car.car = '>'
			car.x++
		}
	case '<':
		if pos == '+' {
			pos = []byte{'/', '-', '\\'}[car.intersections%3]
			car.intersections++
		}
		switch pos {
		case '-':
			car.x--
		case '/':
			car.car = 'v'
			car.y++
		case '\\':
			car.car = '^'
			car.y--
		}
	}
	return car
}

func moveCarts(paths []string, carts []cart) (remaining []cart, crashed bool) {
	sort.Sort(byPos(carts))

	occupied := make(map[string]bool, len(carts))
	for _, c := range carts {
		occupied[c.coord()] = true
	}

	for i, car := range carts {
		if car.crashed {
			continue
		}

		pos := paths[car.y][car.x]

		delete(occupied, car.coord())
		carts[i] = move(car, pos)

		if _, ok := occupied[carts[i].coord()]; ok {
			fmt.Println("Crash at", carts[i].coord())

			for j := range carts {
				if carts[j].coord() == carts[i].coord() { // includes i = j
					carts[j].crashed = true
					delete(occupied, carts[j].coord())
				}
			}
		} else {
			occupied[carts[i].coord()] = true
		}
	}

	for _, c := range carts {
		if !c.crashed {
			remaining = append(remaining, c)
		} else {
			crashed = true
		}
	}

	return
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	paths := strings.Split(string(data), "\n")
	cars := removeCarts(paths)

	for {
		cars, _ = moveCarts(paths, cars)
		if len(cars) == 1 {
			break
		}
	}

	fmt.Println("Last Car at:", cars[0].coord())
}
