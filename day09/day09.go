package main

import (
	"fmt"
)

const players = 10
const lastMarble = 1618

func insert(b []int, current int, val int) []int {
	b = append(b, 0)
	copy(b[current+1:], b[current:])
	b[current] = val

	return b
}

func remove(b []int, current int) []int {
	return append(b[:current], b[current+1:]...)
}

func cw(b []int, current int, amount int) int {
	return (current + amount) % len(b)
}

func ccw(b []int, current int, amount int) (r int) {
	r = current - amount
	if r < 0 {
		return r + len(b)
	}
	return
}

func main() {

	// 455 players; last marble is worth 71223 points
	board := make([]int, 1, lastMarble)
	player, current := 0, 0
	var score [players]int

	highscore := 0

	for marble := 1; marble <= lastMarble; marble++ {
		if marble%23 == 0 {
			score[player] += marble
			current = ccw(board, current, 7)
			score[player] += board[current]
			board = remove(board, current)

			if score[player] > highscore {
				highscore = score[player]
			}
		} else {
			current = cw(board, current, 2)
			board = insert(board, current, marble)
		}

		player = (player + 1) % players
	}

	fmt.Println(highscore)

}
