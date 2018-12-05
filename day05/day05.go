package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func trigger(a byte, b byte) bool {
	return a-b == 32 || a-b == 224 // 256 - 32
}

func reduce(bts []byte) []byte {
	var new bytes.Buffer
	for i := 0; i < len(bts); i++ {
		if i == len(bts)-1 || !trigger(bts[i], bts[i+1]) {
			new.WriteByte(bts[i])
		} else {
			i++
		}
	}
	return new.Bytes()
}

func solve(buf []byte) int {
	for {
		red := reduce(buf)
		if len(red) == len(buf) {
			return len(red)
		}
		buf = red
	}
}

func main() {
	buf, _ := ioutil.ReadFile("input.txt")

	min := solve(buf)
	fmt.Println("Part One:", min)

	for c := byte('A'); c <= 'Z'; c++ {
		b := bytes.Replace(buf, []byte{c}, []byte{}, -1)
		b = bytes.Replace(b, []byte{c + 32}, []byte{}, -1)
		if low := solve(b); low < min {
			min = low
		}
	}

	fmt.Println("Part Two:", min)
}
