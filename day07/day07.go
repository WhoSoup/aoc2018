package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "input.txt"
const letters = 26
const workers = 5
const timeoffset = 60

type node struct {
	dependencies int
	dependents   []int
}
type worker struct {
	letter int
	time   int
}

func main() {

	file, _ := os.Open(filename)
	defer file.Close()

	nodes := make([]node, letters)

	var a, b byte

	for {
		count, err := fmt.Fscanf(file, "Step %c must be finished before step %c can begin.\n", &a, &b)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if count != 2 {
			continue
		}

		nodes[b-'A'].dependencies++
		nodes[a-'A'].dependents = append(nodes[a-'A'].dependents, int(b-'A'))

		//fmt.Printf("%c, %c\n", a, b)
	}

Outer:
	for {
		for letter := range nodes {
			if nodes[letter].dependencies == 0 {
				fmt.Printf("%c", 'A'+letter)
				nodes[letter].dependencies--
				for _, i := range nodes[letter].dependents {
					nodes[i].dependencies--
				}
				continue Outer
			}
		}
		break Outer
	}

	fmt.Println()

	file.Close()
	file, _ = os.Open(filename)
	nodes = make([]node, letters)
	for {
		count, err := fmt.Fscanf(file, "Step %c must be finished before step %c can begin.\n", &a, &b)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if count != 2 {
			continue
		}

		nodes[b-'A'].dependencies++
		nodes[a-'A'].dependents = append(nodes[a-'A'].dependents, int(b-'A'))

		//fmt.Printf("%c, %c\n", a, b)
	}

	pool := make([]worker, workers)
	elapsed := 0

	for {
		// find free worker
		for i := range pool {
			if pool[i].time > 0 {
				pool[i].time--

				if pool[i].time == 0 {
					letter := pool[i].letter
					fmt.Printf("%c", 'A'+letter)
					for _, i := range nodes[letter].dependents {
						nodes[i].dependencies--
					}
				}
			}
		}

		//fmt.Println(pool)

		elapsed++

		done := true

	Assigned:
		for letter := range nodes {
			if nodes[letter].dependencies == 0 {
				for i := range pool {
					if pool[i].time == 0 {
						//fmt.Printf("Assigning %c to Pool#%d\n", letter+'A', i)
						pool[i].letter = letter
						pool[i].time = timeoffset + letter + 1
						nodes[letter].dependencies--
						continue Assigned
					}
				}
			} else if nodes[letter].dependencies > 0 {
				done = false
			}
		}

		if done {
			for _, w := range pool {
				if w.time > 0 {
					done = false
				}
			}

			if done {
				break
			}
		}
	}

	fmt.Println(" ", elapsed-1)

}
