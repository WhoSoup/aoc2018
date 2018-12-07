package main

import (
	"fmt"
	"io"
	"os"
)

const filename = "input.txt"
const letters = 26
const workers = 5
const timeoffset = 60

type node struct {
	parents  int
	children []int
}
type worker struct {
	letter int
	time   int
}

func build() []node {
	file, _ := os.Open(filename)
	defer file.Close()

	nodes := make([]node, letters)

	var a, b byte
	for {
		_, err := fmt.Fscanf(file, "Step %c must be finished before step %c can begin.\n", &a, &b)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}

		nodes[b-'A'].parents++
		nodes[a-'A'].children = append(nodes[a-'A'].children, int(b-'A'))
	}

	return nodes
}

func part1() {
	nodes := build()

Outer:
	for {
		for letter := range nodes {
			if nodes[letter].parents == 0 {
				fmt.Printf("%c", 'A'+letter)
				nodes[letter].parents--
				for _, i := range nodes[letter].children {
					nodes[i].parents--
				}
				continue Outer
			}
		}
		break Outer
	}

	fmt.Println()
}

func part2() {
	nodes := build()

	pool := make([]worker, workers)
	elapsed := 0

Work:
	for {
		elapsed++

		// process workers
		for i := range pool {
			if pool[i].time > 0 { // worker doing work
				pool[i].time--

				if pool[i].time == 0 { // worker is done
					letter := pool[i].letter
					fmt.Printf("%c", 'A'+letter)
					for _, i := range nodes[letter].children {
						nodes[i].parents--
					}
				}
			}
		}

		// assignments
		for letter := range nodes {
			if nodes[letter].parents == 0 {
				for i := range pool {
					if pool[i].time == 0 { // worker available
						pool[i] = worker{letter, timeoffset + letter + 1}
						nodes[letter].parents--
						break
					}
				}
			}
		}

		for _, w := range pool {
			if w.time > 0 {
				continue Work
			}
		}

		break
	}

	fmt.Println(" ", elapsed-1)
}

func main() {
	part1()
	part2()
}
