package d06

import (
	"bufio"
	"math"
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

func findRangeOfTwoRoots(distance int, time int) int {
	d := float64(distance)
	t := float64(time)
	delta := t*t - 4*d
	root1 := math.Floor((t*t - math.Sqrt(delta)) / 2)
	root2 := math.Ceil((t*t + math.Sqrt(delta)) / 2)
	return int(root2 - root1 - 1)
}

func PartOne(file *os.File) int {
	reNumber := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	rawTimes := reNumber.FindAllStringSubmatch(scanner.Text(), -1)

	scanner.Scan()
	rawDistances := reNumber.FindAllStringSubmatch(scanner.Text(), -1)

	values := []int{}
	for i := 0; i < len(rawTimes); i++ {
		time, _ := strconv.Atoi(rawTimes[i][0])
		distance, _ := strconv.Atoi(rawDistances[i][0])
		values = append(values, findRangeOfTwoRoots(distance, time))
	}
	mult := 1
	for _, v := range values {
		mult *= v
	}
	return mult
}

func PartTwo(file *os.File) int {
	reNumber := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	rawTimes := reNumber.FindAllStringSubmatch(scanner.Text(), -1)

	scanner.Scan()
	rawDistances := reNumber.FindAllStringSubmatch(scanner.Text(), -1)

	var time, distance int

	strTime := ""
	for _, subArr := range rawTimes {
		strTime += strings.Join(subArr, ",") + " "
	}
	time, _ = strconv.Atoi(strings.ReplaceAll(strTime, " ", ""))

	strDistance := ""
	for _, subArr := range rawDistances {
		strDistance += strings.Join(subArr, ",") + " "
	}
	distance, _ = strconv.Atoi(strings.ReplaceAll(strDistance, " ", ""))

	return findRangeOfTwoRoots(distance, time)
}
