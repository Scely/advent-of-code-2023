package main

import (
	"advent-of-code-2023/d08"
	"fmt"
	"log"
	"os"
)

func readInputFile() *os.File {
	file, err := os.Open("d08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func main() {
	fmt.Println(d08.PartOne(readInputFile()))
	fmt.Println(d08.PartTwo(readInputFile()))
}
