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
	vx, vy        int
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

var turnLeft = map[byte]byte{
	'^': '<',
	'<': 'v',
	'v': '>',
	'>': '^'}
var turnRight = map[byte]byte{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^'}

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
				c := cart{paths[y][x], x, y, 0, 0, 0}
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

func move(car cart, pos byte) (r cart) {
	return
}

func moveCarts(paths []string, carts []cart) {
	sort.Sort(byPos(carts))

	var occupied map[string]bool
	for _, c := range carts {
		occupied[string(c.x)+","+string(c.y)] = true
	}

	for i, car := range carts {
		pos := paths[car.y][car.x]

		carts[i] = move(car, pos)
	}
}

func main() {
	data, _ := ioutil.ReadFile("input-test.txt")

	fmt.Println(string(data))

	paths := strings.Split(string(data), "\n")

	cars := removeCarts(paths)

	fmt.Println(cars)
	moveCarts(paths, cars)

	fmt.Println(cars)
}
