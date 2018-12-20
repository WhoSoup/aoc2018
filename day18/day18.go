package main

import (
	"bufio"
	"fmt"
	"os"
)

const size, open, tree, yard = 50, 0, 1, 2

func get(f *[size][size]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(*f) || y >= len(*f) {
		return 0
	}
	return (*f)[y][x]
}
func count(f *[size][size]int, x, y int) (trees, yards int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue
			}
			switch get(f, x-1+i, y-1+j) {
			case tree:
				trees++
			case yard:
				yards++
			}
		}
	}
	return
}

func run(f *[size][size]int) (res [size][size]int, cTree, cYard int) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			trees, yards := count(f, x, y)
			switch get(f, x, y) {
			case open:
				if trees >= 3 {
					res[y][x] = tree
					cTree++
				} else {
					res[y][x] = open
				}
			case tree:
				if yards >= 3 {
					res[y][x] = yard
					cYard++
				} else {
					res[y][x] = tree
					cTree++
				}
			case yard:
				if trees > 0 && yards > 0 {
					res[y][x] = yard
					cYard++
				} else {
					res[y][x] = open
				}
			}
		}
	}
	return
}

func print(field [size][size]int) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			switch field[y][x] {
			case open:
				fmt.Print(".")
			case tree:
				fmt.Print("|")
			case yard:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, _ := os.Open("input.txt")

	var field [size][size]int

	scan := bufio.NewScanner(file)
	y := 0
	for scan.Scan() {
		for x := 0; x < len(scan.Text()); x++ {
			switch scan.Text()[x] {
			case '.':
				field[y][x] = open
			case '|':
				field[y][x] = tree
			case '#':
				field[y][x] = yard
			}

		}
		y++
	}

	encountered := make(map[[size][size]int]int)

	var a, b int

	for i := 0; ; i++ {
		field, a, b = run(&field)
		if i == 9 {
			fmt.Println("Part One", a*b)
		}
		if encountered[field] > 0 {
			//fmt.Println("Cycle detected at", i, "to", encountered[field])
			togo := (1000000000 - i - 1) % (i - encountered[field])
			for j := 0; j < togo; j++ {
				field, a, b = run(&field)
			}

			fmt.Println("Part Two:", a*b)
			break
		}
		encountered[field] = i
	}
}
