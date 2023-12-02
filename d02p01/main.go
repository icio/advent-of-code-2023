package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	p := parser{r: bufio.NewReader(os.Stdin)}

	var sum int
	for {
		game, maxR, maxG, maxB, err := readGameMaxes(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(os.Stderr, game, maxR, maxG, maxB)
		if maxR > 12 || maxG > 13 || maxB > 14 {
			continue
		}
		sum += game
	}
	fmt.Println(sum)
}

func readGameMaxes(t *parser) (game, maxR, maxG, maxB int, err error) {
	game, err = t.read(gamestart)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	if _, err := t.read(drawstart); err != nil {
		return 0, 0, 0, 0, err
	}
	for {
		tok, val, err := t.next()
		if err != nil {
			return 0, 0, 0, 0, err
		}
		switch tok {
		case drawstart, drawend:
			// continue.
		case red:
			if val > maxR {
				maxR = val
			}
		case green:
			if val > maxG {
				maxG = val
			}
		case blue:
			if val > maxB {
				maxB = val
			}
		case gameend:
			return game, maxR, maxG, maxB, nil
		}
	}
}

type token string

const (
	invalid   token = ""
	gamestart       = "gamestart"
	drawstart       = "drawstart"
	blue            = "blue"
	red             = "red"
	green           = "green"
	drawend         = "drawend"
	gameend         = "gameend"
)

type parser struct {
	r *bufio.Reader
	q token

	hasNum bool
	num    int
}

func (p *parser) read(exp token) (int, error) {
	tok, val, err := p.next()
	if err != nil {
		return 0, err
	}
	if tok != exp {
		return 0, fmt.Errorf("expected %s token, got %s", exp, tok)
	}
	return val, nil
}

func (p *parser) next() (token, int, error) {
	for {
		if p.q != invalid {
			var next token
			next, p.q = p.q, invalid
			return next, 0, nil
		}

		r, _, err := p.r.ReadRune()
		if err != nil {
			return invalid, 0, err
		}

		if '0' <= r && r <= '9' {
			if p.hasNum {
				p.num = p.num*10 + int(r-'0')
			} else {
				p.num = int(r - '0')
			}
			p.hasNum = true
			continue
		}

		switch r {
		case ' ':
			continue

		case 'G':
			if err := p.discard("ame "); err != nil {
				return invalid, 0, err
			}
			continue
		case ':':
			p.hasNum = false
			p.q = drawstart
			return gamestart, p.num, nil
		case ',':
			p.q = drawstart
			return drawend, 0, nil
		case '\n':
			p.q = gameend
			return drawend, 0, nil

		case 'b':
			if err := p.discard("lue"); err != nil {
				return invalid, 0, err
			}
			p.hasNum = false
			return blue, p.num, nil
		case 'g':
			if err := p.discard("reen"); err != nil {
				return invalid, 0, err
			}
			p.hasNum = false
			return green, p.num, nil
		case 'r':
			if err := p.discard("ed"); err != nil {
				return invalid, 0, err
			}
			p.hasNum = false
			return red, p.num, nil
		}
	}
}

func (p *parser) discard(s string) error {
	v, err := p.r.Peek(len(s))
	if err != nil {
		return err
	}
	if !bytes.Equal(v, []byte(s)) {
		return fmt.Errorf("expected %q saw %q", s, v)
	}
	p.r.Discard(len(s))
	return nil
}
