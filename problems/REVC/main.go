package main

import (
	"bufio"
	"fmt"
	"os"
)

func comp(s rune) rune {
	switch s {
	case 'A':
		return 'T'
	case 'T':
		return 'A'
	case 'C':
		return 'G'
	case 'G':
		return 'C'
	}
	panic("invalid symbol")
	return ' '
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		r := []rune(scanner.Text())
		c := make([]rune, 0, len(r))
		for i := len(r) - 1; i >= 0; i-- {
			c = append(c, comp(r[i]))
		}
		fmt.Println(string(c))
	}
}
