package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type guard struct {
	id    int
	total int
	sleep []int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	data := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		data = append(data, scanner.Text())
	}

	// lexicographical sorting works since it's y-m-d h:m:s
	sort.Strings(data)

	guards := make(map[int]*guard)

	var current *guard

	start := 0

	for i := range data {
		entry := strings.Split(data[i], " ")

		switch entry[2] { // first word after date
		case "Guard":
			id, _ := strconv.Atoi(entry[3][1:]) // second word after date, minus '#'

			_, ok := guards[id]
			if !ok { // make guard if it doesn't exist
				guards[id] = &guard{id, 0, make([]int, 60)}
			}
			current = guards[id]

		case "falls":
			start, _ = strconv.Atoi(entry[1][3:5]) // only the minute is important
		case "wakes":
			end, _ := strconv.Atoi(entry[1][3:5])

			for j := start; j < end; j++ {
				current.sleep[j]++
			}
			current.total += end - start
		}
	}

	// part 1
	var max *guard
	max = &guard{0, 0, []int{}}

	// find max guard
	for _, g := range guards {
		if g.total > max.total {
			max = g
		}
	}

	// get max's highest minute
	mIndex, mValue := 0, 0
	for i, m := range max.sleep {
		if m > mValue {
			mIndex = i
			mValue = m
		}
	}

	fmt.Println("P1 Worker:", max.id, "@", mIndex, "=", max.id*mIndex)

	// part 2
	mIndex, mValue = 0, 0
	for i := 0; i < 60; i++ {
		most, mid := 0, 0
		for _, g := range guards {
			if g.sleep[i] > most {
				most = g.sleep[i]
				mid = g.id
			}
		}

		if most > mValue {
			mIndex = i
			mValue = most
			max = guards[mid]
		}
	}

	fmt.Println("P2 Worker:", max.id, "@", mIndex, "=", max.id*mIndex)
}
