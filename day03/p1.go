package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// DIM fabric max size
const DIM = 1000

type claim struct {
	id    int
	count int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var id, x, y, w, h int

	// just write it all to a big array
	cloth := make([][]claim, DIM)
	for i := range cloth {
		cloth[i] = make([]claim, DIM)
	}

	// keep track of the ids that overlapped
	tainted := make(map[int]bool, 64)

	for {
		count, err := fmt.Fscanf(file, "#%d @ %d,%d: %dx%d\n", &id, &x, &y, &w, &h)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if count != 5 {
			continue
		}

		// write rectangle to array
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				// if it overlaps, taint both rectangles
				if cloth[i][j].id > 0 {
					tainted[cloth[i][j].id] = true
					tainted[id] = true
				}

				// we only need the most recent id
				// since all previous ids will have been tainted
				cloth[i][j].count++
				cloth[i][j].id = id
			}
		}
	}

	overlap := 0
	for i := 0; i < DIM; i++ {
		for j := 0; j < DIM; j++ {
			// count how many square inches overlap
			if cloth[i][j].count > 1 {
				overlap++
			}

			// check if the id is tainted
			if _, taint := tainted[cloth[i][j].id]; cloth[i][j].count == 1 && !taint {
				fmt.Println("ID that doesn't overlap:", cloth[i][j].id)
				tainted[cloth[i][j].id] = true // only print it once
			}
		}
	}

	fmt.Println("Square Inch Overlap:", overlap)
}
