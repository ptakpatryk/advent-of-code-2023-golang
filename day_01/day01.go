package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/ptakpatryk/advent-of-code-2023-golang/lib"
)

var input, err = os.ReadFile("day_01/input.txt")

func main() {
	lib.CheckError(err)
	/* fmt.Println(partOne(string(input))) */
	fmt.Println(partTwo(string(input)))
}

func partOne(fileInput string) int {
	lines := strings.Split(fileInput, "\n")

	var numOuters [][]string

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		var lineNumbers []string
		for _, c := range l {
			letter := string(c)
			_, err := strconv.Atoi(letter)
			if err != nil {
				continue
			}

			lineNumbers = append(lineNumbers, letter)
		}

		if len(lineNumbers) > 0 {
			numOuters = append(numOuters, []string{lineNumbers[0], lineNumbers[len(lineNumbers)-1]})
		}
	}

	var sum int

	for _, numSet := range numOuters {
		pairSum, err := strconv.Atoi(numSet[0] + numSet[1])
		if err != nil {
			continue
		}
		sum += pairSum
	}

	return sum
}

var digitsMap = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func partTwo(fileInput string) int {
	lines := strings.Split(fileInput, "\n")

	var numOuters [][]string

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		lineNumbers := make(map[int]string)
		for i, c := range l {
			letter := string(c)
			_, err := strconv.Atoi(letter)
			if err != nil {
				continue
			}

			lineNumbers[i] = letter
		}

		for digit, digitVal := range digitsMap {
			lineCopy := l

			for strings.Contains(lineCopy, digit) {
				index := strings.Index(lineCopy, digit)
				lineNumbers[index] = digitVal
				lineCopy = strings.Replace(lineCopy, digit, createHashString(len(digit)), 1)
			}
		}

    keys := make([]int, 0, len(lineNumbers))
    for k := range lineNumbers {
      keys = append(keys, k)
    }
    slices.Sort(keys)

    if len(lineNumbers) > 0 {
			numOuters = append(numOuters, []string{lineNumbers[keys[0]], lineNumbers[keys[len(keys) - 1]]})
		}
	}

	var sum int

	for _, numSet := range numOuters {
		pairSum, err := strconv.Atoi(numSet[0] + numSet[1])
		if err != nil {
			continue
		}
		sum += pairSum
	}

	return sum
}

func createHashString(length int) string {
  strArr := make([]string, length, length)
  for i := range strArr {
    strArr[i] = "#"
  }

  str := strings.Join(strArr, "")

  return str
}
