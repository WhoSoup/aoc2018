package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type register [6]int

type instruction struct {
	command string
	a, b, c int
}

type instructions []instruction

type opcodes map[string](func(r register, a, b, c int) register)

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (op opcodes) run(r register, i instruction) register {
	return op[i.command](r, i.a, i.b, i.c)
}

func main() {

	logic := opcodes{
		"addr": func(r register, a, b, c int) register {
			r[c] = r[a] + r[b]
			return r
		},
		"addi": func(r register, a, b, c int) register {
			r[c] = r[a] + b
			return r
		},
		"mulr": func(r register, a, b, c int) register {
			r[c] = r[a] * r[b]
			return r
		},
		"muli": func(r register, a, b, c int) register {
			r[c] = r[a] * b
			return r
		},
		"banr": func(r register, a, b, c int) register {
			r[c] = r[a] & r[b]
			return r
		},
		"bani": func(r register, a, b, c int) register {
			r[c] = r[a] & b
			return r
		},
		"borr": func(r register, a, b, c int) register {
			r[c] = r[a] | r[b]
			return r
		},
		"bori": func(r register, a, b, c int) register {
			r[c] = r[a] | b
			return r
		},
		"setr": func(r register, a, b, c int) register {
			r[c] = r[a]
			return r
		},
		"seti": func(r register, a, b, c int) register {
			r[c] = a
			return r
		},
		"gtir": func(r register, a, b, c int) register {
			r[c] = b2i(a > r[b])
			return r
		},
		"gtri": func(r register, a, b, c int) register {
			r[c] = b2i(r[a] > b)
			return r
		},
		"gtrr": func(r register, a, b, c int) register {
			r[c] = b2i(r[a] > r[b])
			return r
		},
		"eqir": func(r register, a, b, c int) register {
			r[c] = b2i(a == r[b])
			return r
		},
		"eqri": func(r register, a, b, c int) register {
			r[c] = b2i(r[a] == b)
			return r
		},
		"eqrr": func(r register, a, b, c int) register {
			r[c] = b2i(r[a] == r[b])
			return r
		}}

	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	ip := 0
	fmt.Sscanf(lines[0], "#ip%d", &ip)

	var commands []instruction
	for i := 1; i < len(lines); i++ {
		var ins instruction
		fmt.Sscanf(lines[i], "%s %d %d %d", &ins.command, &ins.a, &ins.b, &ins.c)
		commands = append(commands, ins)
	}

	var app register
	pc := 0

	for pc < len(commands) {
		app[ip] = pc
		app = logic.run(app, commands[pc])
		pc = app[ip] + 1
	}

	fmt.Println("Part One:", app[0])

	app = register{1, 0, 0, 0, 0, 0}
	pc = 0

	for pc < len(commands) {
		app[ip] = pc
		app = logic.run(app, commands[pc])
		pc = app[ip] + 1
	}

	fmt.Println("Part Two:", app[0])
}
