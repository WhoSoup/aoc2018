package main

import (
	"fmt"
	"math"
)

type cookbook struct {
	recipes []int
	size    int
	a, b    int
}

func (cb *cookbook) cook() {
	new := cb.recipes[cb.a] + cb.recipes[cb.b]

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
	cb.a = (cb.a + 1 + cb.recipes[cb.a]) % cb.size
	cb.b = (cb.b + 1 + cb.recipes[cb.b]) % cb.size
}

func (cb *cookbook) nums(at, count int) (a int) {
	for i := 0; i < count && at-i >= 0; i++ {
		a += cb.recipes[at-i] * int(math.Pow10(i))
	}
	return
}

func (cb *cookbook) check(target int) bool {
	if target == cb.nums(cb.size-1, 6) {
		fmt.Println("Target found at:", cb.size-6)
		return true
	} else if target == cb.nums(cb.size-2, 6) {
		fmt.Println("Target found at:", cb.size-1-6)
		return true
	}
	return false
}

func main() {
	input := 236021

	cb := cookbook{make([]int, input+10), 2, 0, 1}
	cb.recipes[0] = 3
	cb.recipes[1] = 7

	once := true
	for {
		cb.cook()
		if once && cb.size >= input+10 {
			fmt.Println("Part One:", cb.nums(input-1+10, 10))
			once = false
		}

		ok := cb.check(input)

		if ok {
			break
		}
	}
}
