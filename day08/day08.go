package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	meta     []int
	children []node
}

func getNode(data []int, offset *int) (n node) {
	n.children = make([]node, data[*offset])
	n.meta = make([]int, data[*offset+1])

	*offset += 2

	for i := 0; i < len(n.children); i++ {
		n.children[i] = getNode(data, offset)
	}

	for i := 0; i < len(n.meta); i++ {
		n.meta[i] = data[*offset]
		*offset++
	}

	return
}

func sumMeta(n node) (sum int) {
	for _, v := range n.meta {
		sum += v
	}

	for _, c := range n.children {
		sum += sumMeta(c)
	}

	return
}

func value(n node) (ret int) {
	if len(n.children) == 0 {
		for _, m := range n.meta {
			ret += m
		}
		return
	}

	for _, m := range n.meta {
		if m-1 < len(n.children) {
			ret += value(n.children[m-1])
		}
	}

	return
}

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(buf)), " ")
	data := make([]int, len(split))
	for i, s := range split {
		data[i], _ = strconv.Atoi(s)
	}

	start := 0
	root := getNode(data, &start)

	fmt.Println(sumMeta(root))
	fmt.Println(value(root))
}
