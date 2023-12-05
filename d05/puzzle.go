package d05

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// `^(?P<dst>\d+) (?P<src>\d+) (?P<rg>\d+)`

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
	// parsing
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	reSeeds := regexp.MustCompile(`seeds: (?P<seeds>.*)`)
	seeds := stringNumbertoIntArray(reSeeds.FindStringSubmatch(scanner.Text())[1])
	reRules := regexp.MustCompile(`(?P<dst>\d+) (?P<src>\d+) (?P<rg>\d+)`)

	gameSteps := make(map[int][]map[string]int)
	i := 0
	tmp := []map[string]int{}
	for scanner.Scan() {
		rules := reRules.FindStringSubmatch(scanner.Text())
		result := make(map[string]int)
		if len(rules) == 0 {
			if len(tmp) != 0 {
				gameSteps[i] = tmp
				tmp = []map[string]int{}
				i++
			}
			continue
		}
		for i, name := range reRules.SubexpNames() {
			if i != 0 && name != "" {
				numInt, err := strconv.Atoi(rules[i])
				if err != nil {
					continue
				}
				result[name] = numInt
			}
		}
		tmp = append(tmp, result)
	}
	gameSteps[i] = tmp

	// game
	results := []int{}
	for _, seed := range seeds {
		res := seed
		for i := 0; i < 7; i++ {
			for _, rule := range gameSteps[i] {
				if rule["src"] <= res && res < rule["src"]+rule["rg"] {
					res += rule["dst"] - rule["src"]
					break
				}
			}
		}
		results = append(results, res)
	}
	return slices.Min(results)
}

func PartTwo(file *os.File) int {
	// parsing
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	reSeeds := regexp.MustCompile(`seeds: (?P<seeds>.*)`)
	seeds := stringNumbertoIntArray(reSeeds.FindStringSubmatch(scanner.Text())[1])
	reRules := regexp.MustCompile(`(?P<dst>\d+) (?P<src>\d+) (?P<rg>\d+)`)

	gameSteps := make(map[int][]map[string]int)
	i := 0
	tmp := []map[string]int{}
	for scanner.Scan() {
		rules := reRules.FindStringSubmatch(scanner.Text())
		result := make(map[string]int)
		if len(rules) == 0 {
			if len(tmp) != 0 {
				gameSteps[i] = tmp
				tmp = []map[string]int{}
				i++
			}
			continue
		}
		for i, name := range reRules.SubexpNames() {
			if i != 0 && name != "" {
				numInt, err := strconv.Atoi(rules[i])
				if err != nil {
					continue
				}
				result[name] = numInt
			}
		}
		tmp = append(tmp, result)
	}
	gameSteps[i] = tmp

	brute := 0
	for true {
		res := brute
		for i := 6; i >= 0; i-- {
			for _, rule := range gameSteps[i] {
				res2 := res - rule["dst"] + rule["src"]
				if rule["src"] <= res2 && res2 < rule["src"]+rule["rg"] {
					res = res2
					break
				}
			}
		}
		for i := 0; i < len(seeds); i += 2 {
			if seeds[i] <= res && res < seeds[i]+seeds[i+1] {
				fmt.Println("found", res, seeds[i], seeds[i+1], brute)
				return brute
			}
		}
		brute++
	}
	return 0
}
