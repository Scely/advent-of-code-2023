package main

import (
	"advent-of-code-2023/d05"
	"fmt"
	"log"
	"os"
)

func readInputFile() *os.File {
	file, err := os.Open("d05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func main() {
	fmt.Println(d05.PartOne(readInputFile()))
	fmt.Println(d05.PartTwo(readInputFile()))
}
