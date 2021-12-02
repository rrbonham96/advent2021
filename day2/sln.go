package main

import (
	"fmt"
	"rrbonham96/advent2021/util"
	"strconv"
	"strings"
)

func main() {
	path := util.GetInputPath()
	input := util.ReadLinesToList(path)

	partA(input)
	partB(input)
}

type direction int

const (
	UP direction = iota
	DOWN
	FORWARD
)

type movement struct {
	Direction direction
	Magnitude int
}

func parseMovement(line string) movement {
	fields := strings.Fields(line)
	m, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	var d direction
	switch dir := fields[0]; dir {
	case "up":
		d = UP
	case "down":
		d = DOWN
	case "forward":
		d = FORWARD
	default:
		panic(fmt.Sprintf("Unknown direction %s", dir))
	}

	return movement{
		Direction: d,
		Magnitude: m,
	}
}

func partA(input []string) {
	movements := make([]movement, len(input))
	for _, line := range input {
		movements = append(movements, parseMovement(line))
	}

	x, y := 0, 0
	for _, movement := range movements {
		switch movement.Direction {
		case UP:
			y -= movement.Magnitude
		case DOWN:
			y += movement.Magnitude
		case FORWARD:
			x += movement.Magnitude
		}
	}

	fmt.Printf("Solution: %d\n", x*y)
}

func partB(input []string) {
	movements := make([]movement, len(input))
	for _, line := range input {
		movements = append(movements, parseMovement(line))
	}

	x, y, aim := 0, 0, 0
	for _, movement := range movements {
		switch movement.Direction {
		case UP:
			aim -= movement.Magnitude
		case DOWN:
			aim += movement.Magnitude
		case FORWARD:
			x += movement.Magnitude
			y += aim * movement.Magnitude
		}
	}

	fmt.Printf("Solution: %d\n", x*y)

}
