package main

import "fmt"

type cookbook struct {
	recipes []int
	size    int
	a, b    int
}

func (cb *cookbook) cook() {
	new := cb.recipes[cb.a] + cb.recipes[cb.b]
	if new >= 10 {
		cb.recipes[cb.size] = new / 10
		if cb.size+1 < len(cb.recipes) {
			cb.recipes[cb.size+1] = new % 10
		}
		cb.size += 2
	} else {
		cb.recipes[cb.size] = new
		cb.size++
	}
	cb.a = (cb.a + 1 + cb.recipes[cb.a]) % cb.size
	cb.b = (cb.b + 1 + cb.recipes[cb.b]) % cb.size
}

func main() {
	input := 236021

	cb := cookbook{make([]int, input+10), 2, 0, 1}
	cb.recipes[0] = 3
	cb.recipes[1] = 7

	for {
		cb.cook()
		if cb.size >= len(cb.recipes) {
			break
		}
	}
	fmt.Print("Part One: ")
	for _, d := range cb.recipes[len(cb.recipes)-10:] {
		fmt.Print(d)
	}
	fmt.Println()
}
