package main

import (
	"fmt"
)

func main() {
	var months int
	var pairs int

	fmt.Scanln(&months, &pairs)

	curGen := 1
	prevGen := 1
	for m := 2; m < months; m++ {
		nextGen := prevGen*pairs + curGen
		prevGen = curGen
		curGen = nextGen
	}
	fmt.Println(curGen)
}
