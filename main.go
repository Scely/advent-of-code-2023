package main

import (
	"advent-of-code-2023/d04"
	"fmt"
	"log"
	"os"
)

func readInputFile() *os.File {
	file, err := os.Open("d04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func main() {
	fmt.Println(d04.PartOne(readInputFile()))
	fmt.Println(d04.PartTwo(readInputFile()))
}
