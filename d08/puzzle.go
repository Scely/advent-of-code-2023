package d08

import (
	"bufio"
	"os"
	"regexp"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func PartOne(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()

	mapOfCells := make(map[string][2]string)

	for scanner.Scan() {
		result := make(map[string]string)
		re := regexp.MustCompile(`^(?P<id>\S{3}) = \((?P<left>\S{3}), (?P<right>\S{3})\)`)

		groups := re.FindStringSubmatch(scanner.Text())
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = groups[i]
			}
		}
		mapOfCells[result["id"]] = [2]string{result["left"], result["right"]}
	}

	score := 0
	currentCell := "AAA"
	endCell := "ZZZ"
	for true {
		char := directions[score%len(directions)]
		if char == 'L' {
			currentCell = mapOfCells[currentCell][0]
		} else {
			currentCell = mapOfCells[currentCell][1]
		}
		score++
		if currentCell == endCell {
			break
		}
	}

	return score
}

func PartTwo(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()

	mapOfCells := make(map[string][2]string)

	for scanner.Scan() {
		result := make(map[string]string)
		re := regexp.MustCompile(`^(?P<id>\S{3}) = \((?P<left>\S{3}), (?P<right>\S{3})\)`)

		groups := re.FindStringSubmatch(scanner.Text())
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = groups[i]
			}
		}
		mapOfCells[result["id"]] = [2]string{result["left"], result["right"]}
	}

	currentCells := []string{}
	for k := range mapOfCells {
		if k[len(k)-1] == 'A' {
			currentCells = append(currentCells, k)
		}
	}

	totalScore := 1
	scores := []int{}

	for _, cell := range currentCells {
		currentCell := cell
		score := 0
		for true {
			char := directions[score%len(directions)]
			if char == 'L' {
				currentCell = mapOfCells[currentCell][0]
			} else {
				currentCell = mapOfCells[currentCell][1]
			}
			score++
			if currentCell[len(currentCell)-1] == 'Z' {
				break
			}
		}
		totalScore *= score
		scores = append(scores, score)
	}

	return LCM(scores[0], scores[1], scores[2:]...)
}
