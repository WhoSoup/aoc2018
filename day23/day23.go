package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type bot struct {
	x, y, z, r int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (b bot) distance(o bot) int {
	return abs(b.x-o.x) + abs(b.y-o.y) + abs(b.z-o.z)
}

func (b bot) overlap(o bot) bool {
	return b.distance(o) <= b.r+o.r
}

func max(a, b []bot) []bot {
	if len(a) > len(b) {
		return a
	}
	return b
}

func approach(a, b int) int {
	if a == b {
		return a
	}
	if b > a {
		return a + 2
	}
	return a - 2
}

func (b bot) moveCloser(o bot) bot {
	if rand.Intn(2) == 0 {
		b.x = approach(b.x, o.x)
	}
	if rand.Intn(2) == 0 {
		b.y = approach(b.y, o.y)
	}
	if rand.Intn(2) == 0 {
		b.z = approach(b.z, o.z)
	}
	return b
}

var maxoverlap int

func test(bots []bot) {
	if len(bots) == 0 {
		fmt.Println("fin")
	} else {
		fmt.Print(bots[0], "||")
		test(bots[1:])
	}
}

func ass(bots, overlapping []bot) []bot {
	if len(bots) == 0 || len(overlapping)+len(bots) <= maxoverlap {
		//fmt.Println(bots, len(overlapping))
		return overlapping
	}

	//fmt.Println(len(bots), len(overlapping))
	if len(overlapping) > maxoverlap {
		maxoverlap = len(overlapping)
	}

	overlaps := true
	for _, o := range overlapping {
		if !bots[0].overlap(o) {
			overlaps = false
			break
		}
	}

	if overlaps {
		return max(ass(bots[1:], append(overlapping, bots[0])), ass(bots[1:], overlapping))
	}
	return ass(bots[1:], overlapping)
}

func find(overlapping []bot, b bot) bot {
	var outliers []bot

	for _, c := range overlapping {
		if b.distance(c) > c.r {
			outliers = append(outliers, c)
		}
	}

	if len(outliers) > 0 {
		//fmt.Println(b, len(outliers))

		b = b.moveCloser(outliers[rand.Intn(len(outliers))])
		return find(overlapping, b)
	}

	return b
}

func main() {
	file, _ := os.Open("input.txt")
	r := bufio.NewScanner(file)

	var bots []bot

	var maxBot bot

	for r.Scan() {
		var b bot
		fmt.Sscanf(r.Text(), "pos=<%d,%d,%d>, r=%d", &b.x, &b.y, &b.z, &b.r)
		bots = append(bots, b)
		if b.r > maxBot.r {
			maxBot = b
		}
	}

	//fmt.Println(bots)
	//fmt.Println(maxBot)

	//test(bots)

	inRange := 0
	for _, b := range bots {
		if maxBot.distance(b) <= maxBot.r {
			inRange++
		}
	}

	fmt.Println(inRange)

	overlapping := ass(bots, []bot{})

	finder := bot{0, 0, 0, 0}
	finder = find(overlapping, finder)
	fmt.Println(finder.distance(bot{0, 0, 0, 0}))
}
