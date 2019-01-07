package main

import (
	"bufio"
	"fmt"
	"os"
)

type star struct {
	w, x, y, z, constellation int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}

func (s star) distance(o star) int {
	return abs(s.w-o.w) + abs(s.x-o.x) + abs(s.y-o.y) + abs(s.z-o.z)
}

func rename(stars []star, a, b int) []star {
	for i := range stars {
		if stars[i].constellation == a {
			stars[i].constellation = b
		}
	}
	return stars
}

func count(stars []star) (sum int) {
	found := make(map[int]bool)
	for _, s := range stars {
		if !found[s.constellation] {
			found[s.constellation] = true
			sum++
		}
	}
	return
}

func main() {
	file, _ := os.Open("input.txt")
	s := bufio.NewScanner(file)

	constellations := 0
	var stars []star

	for s.Scan() {
		a := star{}
		fmt.Sscanf(s.Text(), "%d,%d,%d,%d", &a.w, &a.x, &a.y, &a.z)
		for _, s := range stars {
			if s.distance(a) <= 3 {
				if a.constellation > 0 {
					stars = rename(stars, s.constellation, a.constellation)
				} else {
					a.constellation = s.constellation
				}
			}
		}

		if a.constellation == 0 {
			constellations++
			a.constellation = constellations
		}
		stars = append(stars, a)
	}

	fmt.Println(stars)
	fmt.Println(count(stars))
}
