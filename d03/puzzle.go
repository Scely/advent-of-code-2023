package d03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
)

type coordinates struct {
	x int
	y int
}

func makeGrid(file *os.File) []string {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := []string{}
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func findEnginePartCoordinates(grid []string) []coordinates {
	engineParts := []coordinates{}
	for a, line := range grid {
		for b, c := range line {
			chr := string(c)
			if chr != "." && !regexp.MustCompile(`\d`).MatchString(chr) {
				fmt.Print(chr)
				coords := coordinates{x: a, y: b}
				engineParts = append(engineParts, coords)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return engineParts
}

func getValidCoordinates(engineParts []coordinates) map[int][]int {
	m := make(map[int][]int)
	for _, enginePart := range engineParts {
		validCoordinates := []coordinates{}
		validCoordinates = append(validCoordinates, coordinates{x: enginePart.x - 1, y: enginePart.y})
		validCoordinates = append(validCoordinates, coordinates{x: enginePart.x + 1, y: enginePart.y})
		validCoordinates = append(validCoordinates, coordinates{x: enginePart.x, y: enginePart.y - 1})
		validCoordinates = append(validCoordinates, coordinates{x: enginePart.x, y: enginePart.y + 1})
		validCoordinates = append(validCoordinates, coordinates{x: enginePart.x - 1, y: enginePart.y - 1})
		validCoordinates = append(validCoordinates, coordinates{x: enginePart.x + 1, y: enginePart.y + 1})
		validCoordinates = append(validCoordinates, coordinates{x: enginePart.x - 1, y: enginePart.y + 1})
		validCoordinates = append(validCoordinates, coordinates{x: enginePart.x + 1, y: enginePart.y - 1})
		for _, validCoordinate := range validCoordinates {
			if _, ok := m[validCoordinate.x]; ok {
				if !slices.Contains(m[validCoordinate.x], validCoordinate.y) {
					m[validCoordinate.x] = append(m[validCoordinate.x], validCoordinate.y)
				}
			} else {
				m[validCoordinate.x] = []int{validCoordinate.y}
			}
		}
	}
	for _, v := range m {
		sort.Ints(v[:])
	}
	return m
}

func PartOne(file *os.File) int {
	sum := 0
	grid := makeGrid(file)
	validCoordinates := getValidCoordinates(findEnginePartCoordinates(grid))

	for y, line := range grid {
		re := regexp.MustCompile(`\d+`)

		foundNumbers := re.FindAllString(line, -1)
		foundNumbersIndex := re.FindAllStringIndex(line, -1)

		for i, foundNumber := range foundNumbers {
			numberIsValid := false
			for j := foundNumbersIndex[i][0]; j < foundNumbersIndex[i][1]; j++ {
				if slices.Contains(validCoordinates[y], j) {
					numberIsValid = true
				}
			}
			if numberIsValid {
				number, _ := strconv.Atoi(foundNumber)
				sum += number
			}
		}
	}
	return sum
}

func initEnginePartsWithNumbers(engineParts []coordinates) map[coordinates][]int {
	m := make(map[coordinates][]int)
	for _, enginePart := range engineParts {
		m[enginePart] = []int{}
	}
	return m
}

func findEnginePartFromCoordinates(coords coordinates, engineParts []coordinates) coordinates {
	for _, enginePart := range engineParts {
		if enginePart.x+1 == coords.x && enginePart.y == coords.y ||
			enginePart.x-1 == coords.x && enginePart.y == coords.y ||
			enginePart.x == coords.x && enginePart.y+1 == coords.y ||
			enginePart.x == coords.x && enginePart.y-1 == coords.y ||
			enginePart.x+1 == coords.x && enginePart.y+1 == coords.y ||
			enginePart.x-1 == coords.x && enginePart.y-1 == coords.y ||
			enginePart.x+1 == coords.x && enginePart.y-1 == coords.y ||
			enginePart.x-1 == coords.x && enginePart.y+1 == coords.y {
			return enginePart
		}
	}
	return coordinates{}
}

func PartTwo(file *os.File) int {
	sum := 0
	grid := makeGrid(file)
	engineParts := findEnginePartCoordinates(grid)
	validCoordinates := getValidCoordinates(engineParts)
	enginePartsWithNumbers := initEnginePartsWithNumbers(engineParts)

	for y, line := range grid {
		re := regexp.MustCompile(`\d+`)

		foundNumbers := re.FindAllString(line, -1)
		foundNumbersIndex := re.FindAllStringIndex(line, -1)

		for i, foundNumber := range foundNumbers {
			numberIsValid := false
			enginePartFromCoordinates := coordinates{}
			for j := foundNumbersIndex[i][0]; j < foundNumbersIndex[i][1]; j++ {
				if slices.Contains(validCoordinates[y], j) {
					numberIsValid = true
					enginePartFromCoordinates = findEnginePartFromCoordinates(coordinates{x: y, y: j}, engineParts)
					break
				}
			}
			if numberIsValid {
				number, _ := strconv.Atoi(foundNumber)
				enginePartsWithNumbers[enginePartFromCoordinates] = append(enginePartsWithNumbers[enginePartFromCoordinates], number)
			}
		}
	}
	for k, v := range enginePartsWithNumbers {
		fmt.Println(k, v)
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	return sum
}
