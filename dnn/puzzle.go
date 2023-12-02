package dnn

import (
	"bufio"
	"os"
)

func PartOne(file *os.File) int {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}

	return 0
}

func PartTwo(file *os.File) int {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}

	return 0
}
