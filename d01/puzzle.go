package d01

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInputFile() *os.File {
	file, err := os.Open("d01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

var digitnameToDigit = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func PartOne() int {
	file := readInputFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		re := regexp.MustCompile("[0-9]")
		matches := re.FindAllString(scanner.Text(), -1)
		extractedDigits := strings.Join(matches, "")

		newS := extractedDigits
		if len(extractedDigits) > 2 {
			newS = string(extractedDigits[0]) + string(extractedDigits[len(extractedDigits)-1])
		} else if len(extractedDigits) == 1 {
			newS = strings.Repeat(extractedDigits, 2)
		}

		value, _ := strconv.Atoi(newS)
		sum += value

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func PartTwo() int {
	file := readInputFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d).*(one|two|three|four|five|six|seven|eight|nine|\d)|\d`)
		groups := re.FindStringSubmatch(scanner.Text())
		var extractedDigits string
		if groups[1] == "" {
			extractedDigits = groups[0] + groups[0]
		} else {
			leftDigit := groups[1]
			rightDigit := groups[len(groups)-1]
			if len(groups[1]) > 1 {
				leftDigit = digitnameToDigit[groups[1]]
			}
			if len(groups[len(groups)-1]) > 1 {
				rightDigit = digitnameToDigit[groups[len(groups)-1]]
			}

			extractedDigits = leftDigit + rightDigit
		}

		value, _ := strconv.Atoi(extractedDigits)
		sum += value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}
