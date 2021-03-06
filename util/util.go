package util

import (
	"bufio"
	"os"
	"strconv"
)

func GetInputPath() string {
	if len(os.Args) != 2 {
		panic("Please provide the input path")
	}
	return os.Args[1]
}

func ReadLinesToList(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			panic(err)
		}
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadNumbersToList(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, num)
	}
	return lines
}
