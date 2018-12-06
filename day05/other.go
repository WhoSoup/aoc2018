package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		log.Fatalf("could not open input: %v", err)
	}

	reacted := part51(f)
	log.Printf("Length of reaction is: %d\n", len(reacted))
	part52(reacted)
}

func part51(r io.Reader) []rune {
	br := bufio.NewReader(r)

	var result []rune
	for {
		if c, _, err := br.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
		} else {
			if len(result) == 0 {
				result = append(result, c)
				continue
			}

			last := result[len(result)-1]
			switch {
			case unicode.IsUpper(c) && unicode.IsLower(last) && unicode.ToLower(c) == last:
				fallthrough
			case unicode.IsLower(c) && unicode.IsUpper(last) && unicode.ToUpper(c) == last:
				result = result[:len(result)-1]
				break
			default:
				result = append(result, c)
				break
			}
		}
	}

	return result
}

func part52(reacted []rune) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	reactedString := string(reacted)
	bestLength := len(reacted)
	for _, l := range alphabet {
		replaced := strings.Replace(strings.Replace(reactedString, string(l), "", -1), strings.ToUpper(string(l)), "", -1)
		result := part51(strings.NewReader(replaced))
		if bestLength > len(result) {
			bestLength = len(result)
		}
	}

	log.Printf("Best length is: %d\n", bestLength)
}
