package d09

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func PartOne(file *os.File) int {
	scanner := bufio.NewScanner(file)
	reNumber := regexp.MustCompile(`-?\d+`)
	score := 0
	for scanner.Scan() {
		rawNumbers := reNumber.FindAllStringSubmatch(scanner.Text(), -1)

		firstLayer := []int{}
		for _, rawNumber := range rawNumbers {
			number, _ := strconv.Atoi(rawNumber[0])
			firstLayer = append(firstLayer, number)
		}

		matrix := [][]int{firstLayer}
		diffLayer := firstLayer
		for true {
			setOfNumbers := make(map[int]bool)
			tmpLayer := []int{}
			for i := range diffLayer[:len(diffLayer)-1] {
				diffNumber := diffLayer[i+1] - diffLayer[i]
				setOfNumbers[diffNumber] = true
				tmpLayer = append(tmpLayer, diffNumber)
			}
			matrix = append(matrix, tmpLayer)
			diffLayer = tmpLayer

			if len(setOfNumbers) == 1 {
				break
			}
		}
		value := matrix[len(matrix)-1][0]
		for i := len(matrix) - 2; i >= 0; i-- {
			layer := matrix[i]
			value += layer[len(layer)-1]
		}
		score += value
	}

	return score
}

func PartTwo(file *os.File) int {
	scanner := bufio.NewScanner(file)
	reNumber := regexp.MustCompile(`-?\d+`)
	score := 0
	for scanner.Scan() {
		rawNumbers := reNumber.FindAllStringSubmatch(scanner.Text(), -1)

		firstLayer := []int{}
		for _, rawNumber := range rawNumbers {
			number, _ := strconv.Atoi(rawNumber[0])
			firstLayer = append(firstLayer, number)
		}

		matrix := [][]int{firstLayer}
		diffLayer := firstLayer
		for true {
			setOfNumbers := make(map[int]bool)
			tmpLayer := []int{}
			for i := range diffLayer[:len(diffLayer)-1] {
				diffNumber := diffLayer[i+1] - diffLayer[i]
				setOfNumbers[diffNumber] = true
				tmpLayer = append(tmpLayer, diffNumber)
			}
			matrix = append(matrix, tmpLayer)
			diffLayer = tmpLayer
			if len(setOfNumbers) == 1 {
				break
			}
		}

		value := matrix[len(matrix)-1][0]
		for i := len(matrix) - 2; i >= 0; i-- {
			layer := matrix[i]
			value = layer[0] - value
		}
		score += value

	}

	return score
}
