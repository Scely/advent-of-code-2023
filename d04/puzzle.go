package d04

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func stringNumbertoIntArray(s string) []int {
	var numsInt []int
	for _, nb := range strings.Split(s, " ") {
		numInt, err := strconv.Atoi(nb)
		if err != nil {
			continue
		}
		numsInt = append(numsInt, numInt)
	}
	return numsInt
}

func PartOne(file *os.File) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		re := regexp.MustCompile(`Card +\d+: (?P<winningNbs>.+) \| (?P<Nbs>.+)`)
		groups := re.FindStringSubmatch(scanner.Text())
		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = groups[i]
			}
		}

		Nbs := stringNumbertoIntArray(result["Nbs"])
		winningNbs := stringNumbertoIntArray(result["winningNbs"])

		i := 0
		for _, nb := range Nbs {
			for _, winningNb := range winningNbs {
				if nb == winningNb {
					if i == 0 {
						i = 1
					} else {
						i = i * 2
					}
				}
			}
		}
		sum += i
	}

	return sum
}

func PartTwo(file *os.File) int {
	scanner := bufio.NewScanner(file)

	cards := map[int]int{}

	for scanner.Scan() {
		re := regexp.MustCompile(`Card +(?P<id>\d+): (?P<winningNbs>.+) \| (?P<Nbs>.+)`)
		groups := re.FindStringSubmatch(scanner.Text())
		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = groups[i]
			}
		}

		Nbs := stringNumbertoIntArray(result["Nbs"])
		winningNbs := stringNumbertoIntArray(result["winningNbs"])
		cardId, _ := strconv.Atoi(result["id"])

		cards[cardId]++
		clonesNb := cards[cardId]

		i := 0
		for _, nb := range Nbs {
			for _, winningNb := range winningNbs {
				if nb == winningNb {
					i++
					cards[cardId+i] += clonesNb
				}
			}
		}
	}
	sum := 0
	for _, value := range cards {
		sum += value
	}

	return sum
}
