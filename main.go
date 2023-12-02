package main

import (
	"advent-of-code-2023/d01"
	"fmt"
	"log"
	"os"
)

func readInputFile() *os.File {
	file, err := os.Open("d01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func main() {
	fmt.Println(d01.PartOne(readInputFile()))
	fmt.Println(d01.PartTwo(readInputFile()))
}
