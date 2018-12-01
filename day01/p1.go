package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// part 1
	data := []int{}

	var i int
	offset := 0

	for {
		count, err := fmt.Fscanf(file, "%d\n", &i)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// blank line, invalid line
		if count != 1 {
			continue
		}

		offset += i

		data = append(data, i)
	}

	fmt.Println("Part One: ", offset)

	// part 2
	m := make(map[int]bool)
	offset = 0
Outer:
	for {
		for _, j := range data {
			offset += j
			if _, ok := m[offset]; ok {
				fmt.Println("Day Two:", offset)
				break Outer
			}
			m[offset] = true
		}
	}
}
