package main

import "fmt"

const depth = 3198
const targetX = 12
const targetY = 757

var precalc map[string]int

func index(x, y int) int {
	if x == 0 || y == 0 {
		return x*16807 + y*48271
	}

	if x == targetX && y == targetY {
		return 0
	}

	if z, ok := precalc[fmt.Sprintf("%d,%d", x, y)]; ok {
		return z
	}
	z := level(x-1, y) * level(x, y-1)
	precalc[fmt.Sprintf("%d,%d", x, y)] = z
	return z
}

func level(x, y int) int {
	return (index(x, y) + depth) % 20183
}

func main() {
	precalc = make(map[string]int)

	risklevel := 0
	for y := 0; y <= targetY; y++ {
		for x := 0; x <= targetX; x++ {
			/*			switch level(x, y) % 3 {
						case 0:
							fmt.Print(".")
						case 1:
							fmt.Print("=")
						case 2:
							fmt.Print("|")
						}*/
			risklevel += level(x, y) % 3
		}
		//fmt.Println()
	}

	fmt.Println(risklevel)
}
