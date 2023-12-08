package d07

import (
	"bufio"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
)

// https://freshman.tech/snippets/go/concatenate-slices/
func concatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

func HandstoInt(hand string) int {
	// hand is always len == 5
	charToInt := map[rune]int{
		'2': 0,
		'3': 1,
		'4': 2,
		'5': 3,
		'6': 4,
		'7': 5,
		'8': 6,
		'9': 7,
		'T': 8,
		'J': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	score := charToInt[rune(hand[0])]*13*13*13*13 + charToInt[rune(hand[1])]*13*13*13 + charToInt[rune(hand[2])]*13*13 + charToInt[rune(hand[3])]*13 + charToInt[rune(hand[4])]
	return score
}

func HandstoInt2(hand string) int {
	// hand is always len == 5
	charToInt := map[rune]int{
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}
	score := charToInt[rune(hand[0])]*13*13*13*13 + charToInt[rune(hand[1])]*13*13*13 + charToInt[rune(hand[2])]*13*13 + charToInt[rune(hand[3])]*13 + charToInt[rune(hand[4])]
	return score
}

func findType(hand string) []int {
	// hand is always len == 5
	letterCounts := make(map[rune]int)
	for _, letter := range hand {
		letterCounts[letter]++
	}

	counts := []int{}
	for _, count := range letterCounts {
		counts = append(counts[:], count)
	}
	sort.Ints(counts[:])
	return counts
}

func findType2(hand string) []int {
	// hand is always len == 5
	letterCounts := make(map[rune]int)
	for _, letter := range hand {
		letterCounts[letter]++
	}
	jokers := letterCounts['J']
	delete(letterCounts, 'J')

	counts := []int{}
	for _, count := range letterCounts {
		counts = append(counts[:], count)
	}
	sort.Ints(counts[:])
	if len(counts) > 0 {
		counts[len(counts)-1] += jokers
	} else {
		counts = []int{5}
	}
	return counts
}

func PartOne(file *os.File) int {
	re := regexp.MustCompile(`(?P<hand>\S+)\s(?P<bet>\d+)`)
	scanner := bufio.NewScanner(file)

	fiveOfAKind := []int{}
	fourOfAKind := []int{}
	fullHouse := []int{}
	threeOfAKind := []int{}
	twoPairs := []int{}
	onePair := []int{}
	highCard := []int{}

	handIntToBet := make(map[int]int)

	for scanner.Scan() {
		result := make(map[string]string)
		groups := re.FindStringSubmatch(scanner.Text())

		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = groups[i]
			}
		}
		handInt := HandstoInt(result["hand"])
		handIntToBet[handInt], _ = strconv.Atoi(result["bet"])
		if reflect.DeepEqual(findType(result["hand"]), []int{5}) {
			fiveOfAKind = append(fiveOfAKind, handInt)
		} else if reflect.DeepEqual(findType(result["hand"]), []int{1, 4}) {
			fourOfAKind = append(fourOfAKind, handInt)
		} else if reflect.DeepEqual(findType(result["hand"]), []int{2, 3}) {
			fullHouse = append(fullHouse, handInt)
		} else if reflect.DeepEqual(findType(result["hand"]), []int{1, 1, 3}) {
			threeOfAKind = append(threeOfAKind, handInt)
		} else if reflect.DeepEqual(findType(result["hand"]), []int{1, 2, 2}) {
			twoPairs = append(twoPairs, handInt)
		} else if reflect.DeepEqual(findType(result["hand"]), []int{1, 1, 1, 2}) {
			onePair = append(onePair, handInt)
		} else {
			highCard = append(highCard, handInt)
		}
	}
	sort.Ints(fiveOfAKind[:])
	sort.Ints(fourOfAKind[:])
	sort.Ints(fullHouse[:])
	sort.Ints(threeOfAKind[:])
	sort.Ints(twoPairs[:])
	sort.Ints(onePair[:])
	sort.Ints(highCard[:])

	sortedHands := concatMultipleSlices([][]int{highCard, onePair, twoPairs, threeOfAKind, fullHouse, fourOfAKind, fiveOfAKind})
	sum := 0

	for i, hand := range sortedHands {
		sum += handIntToBet[hand] * (i + 1)
	}

	return sum
}

func PartTwo(file *os.File) int {
	re := regexp.MustCompile(`(?P<hand>\S+)\s(?P<bet>\d+)`)
	scanner := bufio.NewScanner(file)

	fiveOfAKind := []int{}
	fourOfAKind := []int{}
	fullHouse := []int{}
	threeOfAKind := []int{}
	twoPairs := []int{}
	onePair := []int{}
	highCard := []int{}

	handIntToBet := make(map[int]int)

	for scanner.Scan() {
		result := make(map[string]string)
		groups := re.FindStringSubmatch(scanner.Text())

		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = groups[i]
			}
		}
		handInt := HandstoInt2(result["hand"])
		handIntToBet[handInt], _ = strconv.Atoi(result["bet"])
		handType := findType2(result["hand"])
		if reflect.DeepEqual(handType, []int{5}) {
			fiveOfAKind = append(fiveOfAKind, handInt)
		} else if reflect.DeepEqual(handType, []int{1, 4}) {
			fourOfAKind = append(fourOfAKind, handInt)
		} else if reflect.DeepEqual(handType, []int{2, 3}) {
			fullHouse = append(fullHouse, handInt)
		} else if reflect.DeepEqual(handType, []int{1, 1, 3}) {
			threeOfAKind = append(threeOfAKind, handInt)
		} else if reflect.DeepEqual(handType, []int{1, 2, 2}) {
			twoPairs = append(twoPairs, handInt)
		} else if reflect.DeepEqual(handType, []int{1, 1, 1, 2}) {
			onePair = append(onePair, handInt)
		} else {
			highCard = append(highCard, handInt)
		}
	}
	sort.Ints(fiveOfAKind[:])
	sort.Ints(fourOfAKind[:])
	sort.Ints(fullHouse[:])
	sort.Ints(threeOfAKind[:])
	sort.Ints(twoPairs[:])
	sort.Ints(onePair[:])
	sort.Ints(highCard[:])

	sortedHands := concatMultipleSlices([][]int{highCard, onePair, twoPairs, threeOfAKind, fullHouse, fourOfAKind, fiveOfAKind})
	sum := 0

	for i, hand := range sortedHands {
		sum += handIntToBet[hand] * (i + 1)
	}

	return sum
}
