package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ruleSet map[string]string

func (rules ruleSet) evolve(base string) (evolution string) {
	evolution = ".."
	base += "....."
	for i := 2; i < len(base)-2; i++ {
		rule, _ := rules[base[i-2:i+3]]
		evolution += rule
	}

	return strings.TrimRight(evolution, ".")
}

func count(s string) (count int) {
	for i, l := range s {
		if l == '#' {
			count += (i - 5)
		}
	}
	return
}

func compare(a string, b string) bool {
	return strings.Trim(a, ".") == strings.Trim(b, ".")
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	initial := lines[0][15:]

	initial = "....." + initial

	rules := make(ruleSet, len(lines)-2)
	for _, rule := range lines[2:] {
		r := strings.Split(rule, " => ")
		rules[r[0]] = r[1]
	}

	big := 50000000000

	for i := 1; ; i++ {
		s := rules.evolve(initial)

		if i == 20 {
			fmt.Println("Part One:", count(s))
		}

		if compare(s, initial) { // period found
			start := count(initial)
			growth := count(s) - start

			//fmt.Println("Period found at", i-1, "to", i, "with growth of", growth)
			fmt.Println("Part Two: ", start+growth*(big-i+1))
			break
		}
		initial = s
	}
}
