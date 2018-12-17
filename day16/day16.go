package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type register [4]int

type opcode struct {
	label  string
	action func(r register, a, b, c int) register
}

type opcodes []opcode

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (op opcodes) run(r register, i register) register {
	return op[i[0]].action(r, i[1], i[2], i[3])
}

func (op opcodes) test(reg, ins, tar register) (matching int, matches []string) {
	for _, o := range op {
		result := o.action(reg, ins[1], ins[2], ins[3])
		if result == tar {
			matching++
			matches = append(matches, o.label)
		}
	}
	return
}

func intersection(sets ...[]string) (isx []string) {
	count := make(map[string]int)
	for _, s := range sets {
		for _, hit := range s {
			count[hit]++
		}
	}

	for code, hits := range count {
		if hits == len(sets) {
			isx = append(isx, code)
		}
	}
	return
}
func arrRemove(ar []string, s string) (res []string) {
	for _, v := range ar {
		if v != s {
			res = append(res, v)
		}
	}
	return
}

func main() {

	logic := opcodes{
		opcode{"addr", func(r register, a, b, c int) register {
			r[c] = r[a] + r[b]
			return r
		}},
		opcode{"addi", func(r register, a, b, c int) register {
			r[c] = r[a] + b
			return r
		}},
		opcode{"mulr", func(r register, a, b, c int) register {
			r[c] = r[a] * r[b]
			return r
		}},
		opcode{"muli", func(r register, a, b, c int) register {
			r[c] = r[a] * b
			return r
		}},
		opcode{"banr", func(r register, a, b, c int) register {
			r[c] = r[a] & r[b]
			return r
		}},
		opcode{"bani", func(r register, a, b, c int) register {
			r[c] = r[a] & b
			return r
		}},
		opcode{"borr", func(r register, a, b, c int) register {
			r[c] = r[a] | r[b]
			return r
		}},
		opcode{"bori", func(r register, a, b, c int) register {
			r[c] = r[a] | b
			return r
		}},
		opcode{"setr", func(r register, a, b, c int) register {
			r[c] = r[a]
			return r
		}},
		opcode{"seti", func(r register, a, b, c int) register {
			r[c] = a
			return r
		}},
		opcode{"gtir", func(r register, a, b, c int) register {
			r[c] = b2i(a > r[b])
			return r
		}},
		opcode{"gtri", func(r register, a, b, c int) register {
			r[c] = b2i(r[a] > b)
			return r
		}},
		opcode{"gtrr", func(r register, a, b, c int) register {
			r[c] = b2i(r[a] > r[b])
			return r
		}},
		opcode{"eqir", func(r register, a, b, c int) register {
			r[c] = b2i(a == r[b])
			return r
		}},
		opcode{"eqri", func(r register, a, b, c int) register {
			r[c] = b2i(r[a] == b)
			return r
		}},
		opcode{"eqrr", func(r register, a, b, c int) register {
			r[c] = b2i(r[a] == r[b])
			return r
		}}}

	data, _ := ioutil.ReadFile("input-op.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	three := 0
	possibilities := make([][][]string, 16)

	for i := 0; i < len(lines); i += 4 {
		var source, instruction, target register
		fmt.Sscanf(lines[i], "Before: [%d, %d, %d, %d]", &source[0], &source[1], &source[2], &source[3])
		fmt.Sscanf(lines[i+1], "%d %d %d %d", &instruction[0], &instruction[1], &instruction[2], &instruction[3])
		fmt.Sscanf(lines[i+2], "After: [%d, %d, %d, %d]", &target[0], &target[1], &target[2], &target[3])

		test, matches := logic.test(source, instruction, target)
		possibilities[instruction[0]] = append(possibilities[instruction[0]], matches)
		if test >= 3 {
			three++
		}
	}
	fmt.Println("Part One", three)

	reduced := make([][]string, 16)
	for i, p := range possibilities {
		reduced[i] = intersection(p...)
	}

	found := 0
	codes := make([]string, 16)

	for found < 16 {
		for i, r := range reduced {
			if len(r) == 1 {
				codes[i] = r[0]
				found++
				for j, k := range reduced {
					reduced[j] = arrRemove(k, r[0])
				}
			}
		}
	}

	orderedLogic := make(opcodes, 16)
	for i := range logic {
		for opcode, label := range codes {
			if logic[i].label == label {
				orderedLogic[opcode] = logic[i]
			}
		}
	}

	data, _ = ioutil.ReadFile("input-code.txt")
	lines = strings.Split(strings.TrimSpace(string(data)), "\n")

	var app register

	for _, l := range lines {
		var inx register
		fmt.Sscanf(l, "%d %d %d %d", &inx[0], &inx[1], &inx[2], &inx[3])
		app = orderedLogic.run(app, inx)
	}

	fmt.Println("Part Two:", app[0])
}
