package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int, 4)
	for scanner.Scan() {
		for _, s := range scanner.Text() {
			counts[string(s)]++
		}
	}
	fmt.Printf("%d %d %d %d\n", counts["A"], counts["C"], counts["G"], counts["T"])
}
