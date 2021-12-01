package main

import (
	"fmt"
	"rrbonham96/advent2021/util"
)

func main() {
	partA()
	partB()
}

func partA() {
	path := util.GetInputPath()
	input := util.ReadNumbersToList(path)

	greaterCount := 0
	l, r := 0, 1
	for r < len(input) {
		if input[r] > input[l] {
			greaterCount++
		}
		l++
		r++
	}

	fmt.Printf("Solution: %d\n", greaterCount)
}

func partB() {
	path := util.GetInputPath()
	input := util.ReadNumbersToList(path)

	greaterCount := 0
	l, r := 0, 1
	for r+2 < len(input) {
		if (input[l] + input[l+1] + input[l+2]) < (input[r] + input[r+1] + input[r+2]) {
			greaterCount++
		}
		l++
		r++
	}

	fmt.Printf("Solution: %d", greaterCount)
}
