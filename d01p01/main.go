package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var sum int64

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Bytes()

		const nums = "0123456789"
		l := bytes.IndexAny(line, nums)
		r := l + bytes.LastIndexAny(line[l:], nums)

		sum += int64(10*(line[l]-'0') + (line[r] - '0'))
	}

	fmt.Println(sum)
}
