package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_04/input.txt")

func main() {
	lib.CheckError(err)
	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

func partOne(fileInput string) int {
	var sum int
	for _, line := range strings.Split(fileInput, "\n") {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, ":")[1]
		splitNumbers := strings.Split(split, "|")
		winStr, resultStr := splitNumbers[0], splitNumbers[1]
		winningsNumbers := strings.Fields(winStr)
		resultNumbers := strings.Fields(resultStr)

		var lineResult int

		for _, num := range resultNumbers {
			if slices.Contains(winningsNumbers, num) {
				if lineResult == 0 {
					lineResult = 1
				} else {
					lineResult *= 2
				}
			}
		}

		sum += lineResult
	}

	return sum
}

func partTwo(fileInput string) int {
	cards := make(map[int]int)
	lines := strings.Split(fileInput, "\n")

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		cardNum := i + 1
		cards[cardNum]++

		split := strings.Split(line, ":")[1]
		splitNumbers := strings.Split(split, "|")
		winStr, resultStr := splitNumbers[0], splitNumbers[1]
		winningsNumbers := strings.Fields(winStr)
		resultNumbers := strings.Fields(resultStr)

		var matches int

		for _, num := range resultNumbers {
			if slices.Contains(winningsNumbers, num) {
				matches++
			}
		}

		for i := 1; i <= matches; i++ {
			cards[cardNum+i] += cards[cardNum]
		}
	}

	var sum int
	for _, v := range cards {
		sum += v
	}

	return sum
}
