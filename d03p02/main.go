package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	next := bytes.Clone(s.Bytes())
	curr := make([]byte, len(next))
	prev := make([]byte, len(next))

	var sum int64
	for {
		if next == nil {
			break
		}

		prev, curr, next = curr, next, prev
		if s.Scan() {
			line := s.Bytes()
			if len(line) != len(curr) {
				panic("changed line length")
			}
			copy(next, line)
		} else {
			if err := s.Err(); err != nil {
				log.Fatal(err)
			}
			next = nil
		}

	gears:
		for i, c := range curr {
			if c == 0 || c == '.' || isnum(c) {
				continue
			}

			var count int
			product := int64(1)
			for _, n := range [...]int64{
				replaceNum(curr, i-1), // west
				replaceNum(curr, i+1), // east
				replaceNum(prev, i),   // north
				replaceNum(prev, i-1), // north west
				replaceNum(prev, i+1), // north east
				replaceNum(next, i),   // south
				replaceNum(next, i-1), // south west
				replaceNum(next, i+1), // south east
			} {
				if n == 0 {
					continue
				}
				count++
				if count > 2 {
					continue gears
				}
				product *= n
			}
			if count == 2 {
				sum += product
			}
		}
	}

	fmt.Println(sum)
}

func isnum(c byte) bool {
	return '0' <= c && c <= '9'
}

// replaceNum looks for an integer at offset i in line and returns the full
// integer using the sequence of digits that continues to the left and right.
// All digits are replaced with '.'. Returns zero if line[i] would panic.
func replaceNum(line []byte, i int) int64 {
	if i < 0 || i >= len(line) || !isnum(line[i]) {
		return 0
	}

	l, r := i, i
	for l > 0 && isnum(line[l-1]) {
		l--
	}
	for r < len(line)-1 && isnum(line[r+1]) {
		r++
	}

	var n int64
	for j := l; j <= r; j++ {
		n = n*10 + int64(line[j]-'0')
		line[j] = '.'
	}

	return n
}
