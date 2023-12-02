package d02

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func PartOne(file *os.File) int {
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reGameId := regexp.MustCompile(`Game (\d+)`)
		reCube := regexp.MustCompile(`(\d+)\s+(\w+)`)

		groupsGameId := reGameId.FindStringSubmatch(scanner.Text())
		groupsCube := reCube.FindAllStringSubmatch(scanner.Text(), -1)

		id, _ := strconv.Atoi(groupsGameId[1])
		valid := true
		for _, group := range groupsCube {
			color := group[2]
			value, _ := strconv.Atoi(group[1])
			if (color == "red" && value > 12) || (color == "green" && value > 13) || (color == "blue" && value > 14) {
				valid = false
				break
			}
		}
		if valid {
			sum += id
		}
	}
	return sum
}

func PartTwo(file *os.File) int {
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		reCube := regexp.MustCompile(`(\d+)\s+(\w+)`)

		groupsCube := reCube.FindAllStringSubmatch(scanner.Text(), -1)

		for _, group := range groupsCube {
			color := group[2]
			value, _ := strconv.Atoi(group[1])
			if cubes[color] < value {
				cubes[color] = value
			}
		}
		sum += cubes["red"] * cubes["green"] * cubes["blue"]
	}

	return sum
}
