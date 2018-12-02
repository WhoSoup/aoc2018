package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// all strings are same lengths
func diff(a string, b string) (int, string) {
	diff := 0
	same := []byte{}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		} else {
			same = append(same, a[i])
		}
	}
	return diff, string(same)
}

func analyze(s string) (bool, bool) {
	double, triple := false, false
	counter := make(map[rune]int)

	for _, i := range s {
		counter[i]++
	}

	for _, v := range counter {
		if v == 2 {
			double = true
		} else if v == 3 {
			triple = true
		}
	}

	return double, triple
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// part 1
	var s string
	single, double := 0, 0

	db := []string{}

	for {
		count, err := fmt.Fscanf(file, "%s\n", &s)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// blank line, invalid line
		if count != 1 {
			continue
		}

		a, b := analyze(s)

		if a {
			single++
		}
		if b {
			double++
		}

		db = append(db, s)
	}

	fmt.Println("Checksum:", (single * double))

Outer:
	for i := 0; i < len(db); i++ {
		for j := i + 1; j < len(db); j++ {
			diff, same := diff(db[i], db[j])
			if diff == 1 {
				fmt.Println("Match:", db[i], db[j], same)
				break Outer
			}
		}
	}
}
