package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"unicode/utf8"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(splitWordsKeepNL)

	var sum int64
	var game int
	for {
		game++
		score, err := readGameScore(s)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		sum += score
	}
	fmt.Println(sum)
}

func readGameScore(s *bufio.Scanner) (score int64, err error) {
	// Skip "Game", "n:"
	s.Scan()
	s.Scan()

	// Read winning numbers
	var winners []int64
	for {
		v, err := next(s)
		if err != nil {
			return 0, err
		}
		if v[0] == '|' {
			break
		}
		n, err := parseInt64(v)
		if err != nil {
			return 0, err
		}
		winners = append(winners, n)
	}
	slices.Sort(winners)

	// Read scratch numbers.
	for {
		v, err := next(s)
		if err != nil {
			return 0, err
		}
		if v[0] == '\n' {
			return score, nil
		}

		n, err := parseInt64(v)
		if err != nil {
			return 0, err
		}
		if _, wins := slices.BinarySearch(winners, n); wins {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
}

func badTokError(exp, got string) error {
	return fmt.Errorf("expected %q, saw %q", exp, got)
}

func parseInt64(v []byte) (int64, error) {
	n, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("expected number, got %q", v)
	}
	return n, nil
}

func next(s *bufio.Scanner) ([]byte, error) {
	if !s.Scan() {
		if err := s.Err(); err != nil {
			return nil, err
		}
		return nil, io.EOF
	}
	return s.Bytes(), nil
}

func discard(s *bufio.Scanner, text string) error {
	got, err := next(s)
	if err != nil {
		return err
	}
	if string(got) != text {
		return badTokError(text, string(got))
	}
	return nil
}

func splitWordsKeepNL(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpaceNotNL(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == '\n' {
			if i == 0 {
				// Return the '\n' on its own.
				return width, data[:i+width], nil
			} else {
				// Return everything up to the '\n', leaving '\n' for next.
				return i, data[start:i], nil
			}
		}
		if isSpaceNotNL(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
func isSpaceNotNL(r rune) bool {
	if r <= '\u00FF' {
		switch r {
		case '\r', '\n':
			return false
		case ' ', '\t', '\v', '\f':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}
