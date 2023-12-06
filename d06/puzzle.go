package d06

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

		root1, root2 := 0, 0
		for x := 0; x < time/2; x++ {
			if -x*x+time*x > distance {
				root1 = x
				break
			}
		}
		for x := time; x > time/2; x-- {
			if -x*x+time*x > distance {
				root2 = x
				break
			}
		}
		values = append(values, root2-root1+1)
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

	values := []int{}
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

	root1, root2 := 0, 0
	for x := 0; x < time/2; x++ {
		if -x*x+time*x > distance {
			root1 = x
			break
		}
	}
	for x := time; x > time/2; x-- {
		if -x*x+time*x > distance {
			root2 = x
			break
		}
	}
	values = append(values, root2-root1+1)
	mult := 1
	for _, v := range values {
		mult *= v
	}
	return mult
}
