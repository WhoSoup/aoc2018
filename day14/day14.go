package main

import (
	"fmt"
)

type cookbook struct {
	recipes    []int
	size       int
	elf, zwölf int
}

func (cb *cookbook) cook() {
	new := cb.recipes[cb.elf] + cb.recipes[cb.zwölf]

	if len(cb.recipes) <= cb.size+2 {
		bigger := make([]int, len(cb.recipes)*2)
		copy(bigger, cb.recipes)
		cb.recipes = bigger
	}

	if new >= 10 {
		cb.recipes[cb.size] = new / 10
		cb.recipes[cb.size+1] = new % 10
		cb.size += 2
	} else {
		cb.recipes[cb.size] = new
		cb.size++
	}
	cb.elf = (cb.elf + 1 + cb.recipes[cb.elf]) % cb.size
	cb.zwölf = (cb.zwölf + 1 + cb.recipes[cb.zwölf]) % cb.size
}

func main() {
	input := 236021
	target := []int{2, 3, 6, 0, 2, 1}

	cb := cookbook{make([]int, input), 2, 0, 1}
	cb.recipes[0] = 3
	cb.recipes[1] = 7

	i := 0

Kitchen:
	for {
		cb.cook()

		for ; i <= cb.size; i++ { // i runs once or twice, depending on how many recipes were added
			if i == input+10 {
				fmt.Print("Part One:")
				for j := input; j < i; j++ {
					fmt.Print(cb.recipes[j])
				}
				fmt.Println()
			}

			if i < len(target) {
				continue
			}

			m := true
			for x := 0; x < len(target); x++ {
				if target[x] != cb.recipes[i-len(target)+x] {
					m = false
					break
				}
			}

			if m {
				fmt.Println("Part Two:", i-len(target))
				break Kitchen
			}
		}
	}
}
