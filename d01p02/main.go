package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sum int64

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Bytes()

		l, ok := numFwd(line)
		if !ok {
			println("invalid: " + string(line))
			continue
		}

		r, ok := numRev(line)
		if !ok {
			println("invalid: " + string(line))
			continue
		}

		// fmt.Println(string(line), l, r)
		sum += int64(l*10 + r)
	}

	fmt.Println(sum)
}

func numFwd(s []byte) (int, bool) {
	for i := range s {
		rem := len(s) - i
		switch s[i] {
		case '0':
			return 0, true
		case '1':
			return 1, true
		case '2':
			return 2, true
		case '3':
			return 3, true
		case '4':
			return 4, true
		case '5':
			return 5, true
		case '6':
			return 6, true
		case '7':
			return 7, true
		case '8':
			return 8, true
		case '9':
			return 9, true

		case 'o': // one
			if rem < 3 {
				continue
			}
			if s[i+1] == 'n' && s[i+2] == 'e' {
				return 1, true
			}
		case 't': // two, three
			if rem < 3 {
				continue
			}
			switch s[i+1] {
			case 'w': // two
				if s[i+2] == 'o' {
					return 2, true
				}
			case 'h': // three
				if rem < 5 {
					continue
				}
				if s[i+2] == 'r' && s[i+3] == 'e' && s[i+4] == 'e' {
					return 3, true
				}
			}
		case 'f': // four, five
			if rem < 4 {
				continue
			}
			switch s[i+1] {
			case 'o': // four
				if s[i+2] == 'u' && s[i+3] == 'r' {
					return 4, true
				}
			case 'i': // five
				if s[i+2] == 'v' && s[i+3] == 'e' {
					return 5, true
				}
			}
		case 's': // six, seven
			if rem < 3 {
				continue
			}
			switch s[i+1] {
			case 'i': // six
				if s[i+2] == 'x' {
					return 6, true
				}
			case 'e': // seven
				if rem < 5 {
					continue
				}
				if s[i+2] == 'v' && s[i+3] == 'e' && s[i+4] == 'n' {
					return 7, true
				}
			}
		case 'e': // eight
			if rem < 5 {
				continue
			}
			if s[i+1] == 'i' && s[i+2] == 'g' && s[i+3] == 'h' && s[i+4] == 't' {
				return 8, true
			}
		case 'n': // nine
			if rem < 4 {
				continue
			}
			if s[i+1] == 'i' && s[i+2] == 'n' && s[i+3] == 'e' {
				return 9, true
			}
		case 'z': // zero
			if rem < 4 {
				continue
			}
			if s[i+1] == 'e' && s[i+2] == 'r' && s[i+3] == 'o' {
				return 0, true
			}
		}
	}
	return 0, false
}

func numRev(s []byte) (int, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		rem := i
		switch s[i] {
		case '0':
			return 0, true
		case '1':
			return 1, true
		case '2':
			return 2, true
		case '3':
			return 3, true
		case '4':
			return 4, true
		case '5':
			return 5, true
		case '6':
			return 6, true
		case '7':
			return 7, true
		case '8':
			return 8, true
		case '9':
			return 9, true

		case 'e': // one, three, five, nine
			if rem < 3 {
				continue
			}
			switch s[i-1] {
			case 'n': // one, nine
				if rem < 3 {
					continue
				}
				switch s[i-2] {
				case 'o': // one
					return 1, true
				case 'i': // nine
					if rem < 4 {
						continue
					}
					if s[i-3] == 'n' {
						return 9, true
					}
				}
			case 'e': // three
				if rem < 5 {
					continue
				}
				if s[i-2] == 'r' && s[i-3] == 'h' && s[i-4] == 't' {
					return 3, true
				}
			case 'v': // five
				if rem < 4 {
					continue
				}
				if s[i-2] == 'i' && s[i-3] == 'f' {
					return 5, true
				}
			}
		case 'o': // two, zero
			if rem < 3 {
				continue
			}
			switch s[i-1] {
			case 'w': // two
				if s[i-2] == 't' {
					return 2, true
				}
			case 'r': // zero
				if rem < 4 {
					continue
				}
				if s[i-2] == 'e' && s[i-3] == 'z' {
					return 0, true
				}
			}
		case 'r': // four
			if rem < 4 {
				continue
			}
			if s[i-1] == 'u' && s[i-2] == 'o' && s[i-3] == 'f' {
				return 4, true
			}
		case 'x': // six
			if rem < 3 {
				continue
			}
			if s[i-1] == 'i' && s[i-2] == 's' {
				return 6, true
			}
		case 'n': // seven
			if rem < 5 {
				continue
			}
			if s[i-1] == 'e' && s[i-2] == 'v' && s[i-3] == 'e' && s[i-4] == 's' {
				return 7, true
			}
		case 't': // eight
			if rem < 5 {
				continue
			}
			if s[i-1] == 'h' && s[i-2] == 'g' && s[i-3] == 'i' && s[i-4] == 'e' {
				return 8, true
			}
		}
	}
	return 0, false
}
