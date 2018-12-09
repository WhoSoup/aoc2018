package main

import (
	"fmt"
)

type llist struct {
	val        int
	prev, next *llist
}

func (l *llist) insertToTheRight(val int) *llist {
	node := &llist{val, l, l.next}

	l.next.prev = node
	l.next = node

	return node
}

func (l *llist) removeThis() (*llist, int) {
	next := l.next
	prev := l.prev

	prev.next = next
	next.prev = prev

	return next, l.val
}

func game(players int, marbles int) (highscore int) {
	current := &llist{0, nil, nil}
	current.prev = current
	current.next = current

	score := make([]int, players)

	for marble := 1; marble <= marbles; marble++ {
		player := (marble - 1) % players
		if marble%23 == 0 {
			score[player] += marble

			// enjoy
			ptr, val := current.prev.prev.prev.prev.prev.prev.prev.removeThis()
			current = ptr
			score[player] += val
			if score[player] > highscore {
				highscore = score[player]
			}
		} else {
			current = current.next.insertToTheRight(marble)
		}
	}

	return
}

func main() {
	// 455 players; last marble is worth 71223 points
	fmt.Println("Part One:", game(455, 71223))
	fmt.Println("Part Two:", game(455, 71223*100))
}
