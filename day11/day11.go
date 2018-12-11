package main

import (
	"fmt"
)

const serial = 7511

func g(x int, y int, size int) (sum int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += p(x+i, y+j)
		}
	}
	return
}

func p(x int, y int) (power int) {
	power = (x + 10) * y
	power += serial
	power *= (x + 10)
	power /= 100
	power %= 10
	power -= 5

	return
}

func largest(power [][]int, dim int) (largest int, x int, y int) {
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			sum := 0

			for k := 0; k < dim; k++ {
				for l := 0; l < dim; l++ {
					sum += power[i+k-1][j+l-1]
				}
			}

			if sum > largest {
				largest = sum
				x = i
				y = j
			}
		}
	}
	return
}

func main() {
	power := make([][]int, 600)
	for i := range power {
		power[i] = make([]int, 600)
	}

	for i := 0; i < 600; i++ {
		for j := 0; j < 600; j++ {
			power[i][j] = p(i+1, j+1)
		}
	}
	_, x, y := largest(power, 3)
	fmt.Printf("Part One: %d,%d\n", x, y)

	max, msize := 0, 0
	for size := 1; size <= 300; size++ {
		m, i, j := largest(power, size)
		if m > max {
			max = m
			msize = size
			x = i
			y = j
		}

		if size%10 == 0 {
			fmt.Println(size, "done")
		}
	}

	fmt.Printf("Part Two: %d,%d,%d\n", x, y, msize)
}
